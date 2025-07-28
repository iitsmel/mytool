package configure

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

func Help(flags *flag.FlagSet) {
	fmt.Print("command format: mytool [arguments] {target}\n\n")
	fmt.Print("flags:\n")

	var flagName []string

	flags.VisitAll(func(f *flag.Flag) {
		flagName = append(flagName, f.Name)
	})

	sort.Strings(flagName)

	for _, name := range flagName {
		f := flags.Lookup(name)
		usage := strings.ReplaceAll(f.Usage, "string", "")
		fmt.Printf("  -%s\t%s\n", f.Name, usage)
	}

	print("\n")
	os.Exit(0)
}

func CleanResultFiles() {
	os.RemoveAll("results/")
}

func LoadingIcon(stop chan bool) {
	spinner := []string{"^owo^", "^=w=^"} // Customize the spinner with desired Unicode characters

	for {
		for _, symbol := range spinner {
			fmt.Printf("\r%s ", symbol)
			time.Sleep(250 * time.Millisecond)
		}
		select {
		case <-stop:
			// Stop the goroutine when a signal is received on the stop channel
			return
		default:
		}
	}
}

func BabyCastingMagic() {
	art := `
           ∧＿∧
          (・◦・)   ✧･ﾟ:*     MAGIC PEU PEU
          /づ~ ♡･ﾟ:*:･★‧₊˚
		  `
	fmt.Println(art)
}

func BabyAsking() {
	art := `
                    /＞　 フ
                   | 　_　_|     of course my child
                  ／ ミ＿xノ 
                 /　　　　|
                /　　　　 |
               /　　　　  |
              /　 ヽ　　  ﾉ
             / │　　|　|　|       ╱|、
            / │　　 |　|　|      (˚ˎ。7   may I perform some magic
    _______|  |     |　|　|      |、 ˜〵
   (___________ヽ___ヽ_)__)      じしˍ,)ノ
		

		  `
	fmt.Println(art)
}
