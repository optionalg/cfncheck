package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/keithsharp/cfncheck/rules"

	"github.com/awslabs/goformation"
)

var (
	check        = flag.Bool("check", false, "run checks as well as lint")
	quiet        = flag.Bool("quiet", false, "suppress test outputs")
	warn         = flag.Bool("warn", false, "treat warnings as failures")
	templateFile = flag.String("template", "", "the CloudFormation template to analyse")
)

func main() {
	flag.Parse()
	if *templateFile == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	template, err := goformation.Open(*templateFile)
	if err != nil {
		log.Fatalf("could not parse template file %s, error %v", *templateFile, err)
	}

	var lint, chk []rules.Result

	lint, err = rules.LintCloudFormationTemplate(*template)
	if err != nil {
		log.Fatalf("could not lint template file %s, error %v", *templateFile, err)
	}

	if *check == true {
		chk, err = rules.CheckCloudFormationTemplate(*template)
		if err != nil {
			log.Fatalf("could not check template file %s, error %v", *templateFile, err)
		}
	}

	// Print out results
	if len(lint) > 0 {
		fmt.Printf("LINT: %v\n", lint)
	}
	if len(chk) > 0 {
		fmt.Printf("CHECK: %v\n", chk)
	}
	if len(lint) > 0 || len(chk) > 0 {
		os.Exit(2)
	}
}
