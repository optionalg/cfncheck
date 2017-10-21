package main

import (
	"flag"
	"log"
	"os"

	"github.com/awslabs/goformation"
)

var (
	quiet        = flag.Bool("quiet", false, "suppress test outputs")
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
	if *quiet != true {
		log.Printf("version: %s", template.AWSTemplateFormatVersion)
		log.Printf("description: %s", template.Description)
	}
}
