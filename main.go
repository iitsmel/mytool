package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"

	"mytool/configure"
)

func Settings(flags *configure.Scope, args []string) configure.Scope {

	flagSet := flag.NewFlagSet("pentest", flag.ContinueOnError)

	// default feature
	flagSet.BoolVar(&flags.General.Delete, "r", false, "Delete all files in Result folder.")
	flagSet.StringVar(&flags.Attack.Target, "u", flags.Attack.Target, "Show HTTP response status code.")
	flagSet.StringVar(&flags.Attack.Target, "s", flags.Attack.Target, "Subdomains Enumeration")
	flagSet.BoolVar(&flags.General.Help, "h", false, "Show flags useage.")

	// archived
	//flagSet.BoolVar(&flags.General.Ignore, "i", false, "Proceed when HTTP response is not 200.")
	//flagSet.IntVar(&flags.HTTP.Cookie, "cookie", flags.HTTP.Cookie, "Length of cookie.")

	// pentest attack
	flagSet.BoolVar(&flags.Attack.XSS, "xss", false, "Use XSS attack.")
	flagSet.BoolVar(&flags.Attack.SQL, "sql", false, "Use SQL Injection attack.")

	// auto-testing plans
	flagSet.StringVar(&flags.Plan.Target, "OA", flags.Plan.Target, "full on aggressive mode")
	flagSet.StringVar(&flags.Plan.Target, "OP", flags.Plan.Target, "reduced passive mode")

	// Set the output destination to a custom buffer for capturing the error message
	errorBuffer := bytes.Buffer{}
	flagSet.SetOutput(&errorBuffer)

	// Set up environment
	configure.EnvironmentSetup()

	// Check for errors
	flagerr := flagSet.Parse(args)

	if flagerr != nil {
		configure.BabyAsking()
		fmt.Println(flagerr)
		configure.Help(flagSet)
	}

	// Print help message and exit if
	// 1. there is no flag
	// 2. the -h flag is provided
	if len(args) == 0 || flags.General.Help {
		configure.BabyAsking()
		configure.Help(flagSet)
	}

	if flags.General.Delete {
		configure.CleanResultFiles()
	}

	return *flags
}

func main() {
	ptrt := configure.Scope{}
	args := os.Args[1:]
	Settings(&ptrt, args)
	configure.ParseFlags(&ptrt, args)
}
