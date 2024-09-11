package configure

import (
	"fmt"
	"os"
	"strings"

	"mytool/general"
)

type OutcomeFunc func(inputconfig *Scope)

// find flag
func FlagFound(flags []string, targetFlag string) bool {
	for _, flag := range flags {
		if flag == targetFlag {
			return true
		}
	}
	return false
}

// distribute flag
func ParsingFlag(inputconfig *Scope, flags []string) {
	target := FilterFlags(flags, inputconfig)
	flagSequence := []string{"-i", "-s", "-cookie", "-u", "-xss", "-sql"}

	// check if flag is in defined scope
	for _, flag := range flagSequence {
		if FlagFound(inputconfig.General.FlagGarage, flag) {
			inputconfigCopy := *inputconfig
			inputconfigCopy.HTTP.URL = target
			GeneratingResponse(flag, &inputconfigCopy)
		}
	}
}

func GeneratingResponse(flag string, inputconfig *Scope) {
	conditionMap := map[string]OutcomeFunc{
		"-u":      func(httpconfig *Scope) { HTTPresponse(httpconfig) },
		"-cookie": func(httpconfig *Scope) { general.IterateCookie(httpconfig) },
		//"-s":   func(subemu *Scope) { SubdomainEnumeration(subemu) },
		//"-xss": func(xss *Scope) { CrossSiteScript(xss) },
		//"-sql": func(sql *Scope) { SQLInjection(sql) },
		//"-OA":  func(tool *Scope) { ENAC(tool) },
		//"-OP":  func(tool *Scope) { Light(tool) },
	}
	if outcomeFunc, ok := conditionMap[flag]; ok {
		BabyCastingMagic()
		outcomeFunc(inputconfig)
	} else {
		fmt.Println("Invalid flag:", flag)
	}
}

func FilterFlags(flags []string, inputconfig *Scope) string {
	var url string
	isURL := false

	// sort flags
	for _, flag := range flags {

		switch {
		// find the URL value from the flags
		case isURL:
			url = strings.TrimPrefix(flag, "https://")
			url = strings.TrimPrefix(url, "http://")
			inputconfig.HTTP.URL = url
			return url

		case strings.HasPrefix(flag, "http://") || strings.HasPrefix(flag, "https://"):
			if url != "" {
				fmt.Println("Multiple URLs are not allowed. Please provide only one URL.")
				os.Exit(1)
			}
			url = flag

		// append flags into flag garage
		default:
			inputconfig.General.FlagGarage = append(inputconfig.General.FlagGarage, flag)
		}

		if flag == "-OA" || flag == "-OP" {
			isURL = true
		}

	}

	return url
}
