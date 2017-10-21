package rules

import "github.com/awslabs/goformation/cloudformation"

// CheckEc2SecurityGroup checks the security group
func CheckEc2SecurityGroup(sg cloudformation.AWSEC2SecurityGroup) []Result {
	var results []Result

	// Check for required properties
	if sg.GroupDescription == "" {
		results = append(results, NewResult(2, "security groups must have a Group Description"))
	}

	return results
}
