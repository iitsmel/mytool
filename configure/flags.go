package configure

import (
	"fmt"
	"os"
	"strings"
)

type OutcomeFunc func(penconfig *Scope)

func ParseFlags(penconfig *Scope, flags []string) {
	target := FilteringURL(flags, penconfig)
	flagSequence := []string{"-r", "-h", "-u", "-s", "-xss", "-sql", "-OA", "-OP"}
	BabyCastingMagic()

	for _, flag := range flagSequence {
		if FlagFound(penconfig.General.FlagGarage, flag) {
			penconfigCopy := *penconfig
			penconfigCopy.Attack.Target = target
			GeneratingResponse(flag, &penconfigCopy)
			fmt.Println()
		}
	}
}

func GeneratingResponse(flag string, penconfig *Scope) {
	CheckTargetStatus(penconfig)

	conditionMap := map[string]OutcomeFunc{
		"-u":   func(httpconfig *Scope) { ShowResponse(httpconfig) },
		"-s":   func(subemu *Scope) { SubdomainEnumeration(subemu) },
		"-xss": func(xss_config *Scope) { CrossSiteScript(xss_config) },
		"-sql": func(sql_config *Scope) { RunInjection(sql_config) },
		"-OA":  func(tool *Scope) { DoENAC(tool) },
		"-OP":  func(tool *Scope) { DoLight(tool) },
	}
	if outcomeFunc, ok := conditionMap[flag]; ok {
		outcomeFunc(penconfig)
	} else {
		fmt.Println("Invalid flag:", flag)
	}
}

func FilteringURL(flags []string, penconfig *Scope) string {
	var url string
	isPlan := false

	// Find the URL value from the flags
	for _, flag := range flags {
		switch {
		case isPlan:
			url = strings.TrimPrefix(flag, "https://")
			url = strings.TrimPrefix(url, "http://")
			penconfig.Plan.Target = url
			return url

		case strings.HasPrefix(flag, "http://") || strings.HasPrefix(flag, "https://"):
			if url != "" {
				fmt.Println("Multiple URLs are not allowed. Please provide one URL per time.")
				os.Exit(1)
			}
			url = flag

		default:
			penconfig.General.FlagGarage = append(penconfig.General.FlagGarage, flag)
		}

		if flag == "-OA" || flag == "-OP" {
			isPlan = true
		}
	}

	return url
}

func FlagFound(flags []string, targetFlag string) bool {
	for _, flag := range flags {
		if flag == targetFlag {
			return true
		}
	}
	return false
}
