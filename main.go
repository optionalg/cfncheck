package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	. "github.com/keithsharp/cfncheck/rules"

	"github.com/awslabs/goformation"
)

var (
	quiet        = flag.Bool("quiet", false, "suppress test outputs")
	warn         = flag.Bool("warn", false, "treat warnings as failures")
	templateFile = flag.String("template", "", "the CloudFormation template to analyse")
)

func main() {
	flag.Parse()
	if *templateFile == "" {
		flag.PrintDefaults()
		os.Exit(3)
	}

	template, err := goformation.Open(*templateFile)
	if err != nil {
		log.Fatalf("could not parse template file %s, error %v", *templateFile, err)
	}

	results := make(map[string][]Result)

	// Check EC2 Security groups
	sgs := template.GetAllAWSEC2SecurityGroupResources()
	for name, sg := range sgs {
		r := CheckEc2SecurityGroup(sg)
		if len(r) > 0 {
			results[name] = r
		}
	}

	// Print out results
	if *quiet != true {
		for name, details := range results {
			fmt.Printf("Resource: %s\n", name)
			for _, r := range details {
				fmt.Printf("\t")
				if r.State() == 0 {
					fmt.Printf("PASSED")
				} else if r.State() == 1 {
					fmt.Printf("WARNING")
				} else if r.State() == 2 {
					fmt.Printf("FAILED")
				}
				fmt.Printf(": %s.\n", r.Description())
			}
		}
		os.Exit(1)
	}

}
