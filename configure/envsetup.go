package configure

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"sync"
)

const configFile = "packages_config.json"

// PackageInfo represents package information stored in the config file
type PackageInfo struct {
	Name      string `json:"name"`
	Installed bool   `json:"installed"`
}

func EnvironmentSetup() {
	packages := []string{"sqlmap", "nmap"}

	switch runtime.GOOS {
	case "darwin":
		checkAndInstallHomebrewPackages(packages)
	case "linux":
		checkAndInstallLinuxPackages(packages)
	default:
		fmt.Println("Unsupported OS")
		os.Exit(1)
	}
}

func loadConfig() map[string]*PackageInfo {
	config := make(map[string]*PackageInfo)

	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		return config
	}

	file, err := os.Open(configFile)
	if err != nil {
		fmt.Println("Error opening config file:", err)
		os.Exit(1)
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&config); err != nil {
		fmt.Println("Error decoding config file:", err)
		os.Exit(1)
	}

	return config
}

func saveConfig(config map[string]*PackageInfo) {
	file, err := os.Create(configFile)
	if err != nil {
		fmt.Println("Error creating config file:", err)
		os.Exit(1)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(config); err != nil {
		fmt.Println("Error encoding config file:", err)
		os.Exit(1)
	}
}

func checkAndInstallHomebrewPackages(packages []string) {
	config := loadConfig()
	var wg sync.WaitGroup
	packagesToInstall := []string{}
	mu := sync.Mutex{}

	for _, pkg := range packages {
		wg.Add(1)
		go func(pkg string) {
			defer wg.Done()
			if info, exists := config[pkg]; !exists || !info.Installed {
				if !isPackageInstalledHomebrew(pkg) {
					mu.Lock()
					packagesToInstall = append(packagesToInstall, pkg)
					config[pkg] = &PackageInfo{Name: pkg, Installed: false}
					mu.Unlock()
				}
			} else {
				config[pkg].Installed = true
			}
		}(pkg)
	}
	wg.Wait()

	if len(packagesToInstall) > 0 {
		for _, pkg := range packagesToInstall {
			fmt.Printf("Installing %s...\n", pkg)
			runCommand("brew", "install", pkg)
			config[pkg].Installed = true
		}
		saveConfig(config)
	}
}

func isPackageInstalledHomebrew(pkg string) bool {
	cmd := exec.Command("brew", "list", "--formula")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error checking installed packages:", err)
		os.Exit(1)
	}
	return strings.Contains(string(output), pkg)
}

func checkAndInstallLinuxPackages(packages []string) {
	config := loadConfig()
	var wg sync.WaitGroup
	packagesToInstall := []string{}
	mu := sync.Mutex{}

	for _, pkg := range packages {
		wg.Add(1)
		go func(pkg string) {
			defer wg.Done()
			if info, exists := config[pkg]; !exists || !info.Installed {
				if !isPackageInstalledLinux(pkg) {
					mu.Lock()
					packagesToInstall = append(packagesToInstall, pkg)
					config[pkg] = &PackageInfo{Name: pkg, Installed: false}
					mu.Unlock()
				}
			} else {
				config[pkg].Installed = true
			}
		}(pkg)
	}
	wg.Wait()

	if len(packagesToInstall) > 0 {
		for _, pkg := range packagesToInstall {
			fmt.Printf("Installing %s...\n", pkg)
			runCommand("sudo", "apt-get", "install", "-y", pkg)
			config[pkg].Installed = true
		}
		saveConfig(config)
	}
}

func isPackageInstalledLinux(pkg string) bool {
	cmd := exec.Command("dpkg", "--list")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error checking installed packages:", err)
		os.Exit(1)
	}
	return strings.Contains(string(output), pkg)
}

func runCommand(name string, args ...string) {
	cmd := exec.Command(name, args...)
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error running command %s %v: %s\n", name, args, err)
		os.Exit(1)
	}
}
