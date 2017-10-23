package rules

import (
	"github.com/awslabs/goformation/cloudformation"
)

// CheckCloudFormationTemplate checks resources against rules
func CheckCloudFormationTemplate(template cloudformation.Template) ([]Result, error) {
	var results []Result
	sgs := template.GetAllAWSEC2SecurityGroupResources()
	for name, sg := range sgs {
		r := CheckEc2SecurityGroup(name, sg)
		for _, result := range r {
			results = append(results, result)
		}
	}
	return results, nil
}

// LintCloudFormationTemplate lints resources against rules
func LintCloudFormationTemplate(template cloudformation.Template) ([]Result, error) {
	var results []Result
	sgs := template.GetAllAWSEC2SecurityGroupResources()
	for name, sg := range sgs {
		r := LintEc2SecurityGroup(name, sg)
		for _, result := range r {
			results = append(results, result)
		}
	}
	return results, nil
}
