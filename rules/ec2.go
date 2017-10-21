package rules

import "github.com/awslabs/goformation/cloudformation"

// CheckEc2SecurityGroup checks the security group
func CheckEc2SecurityGroup(sg cloudformation.AWSEC2SecurityGroup) []Result {
	var results []Result

	// Check for required properties
	if sg.GroupDescription == "" {
		r := NewResult(Failed, "security groups must have a Group Description", sg.AWSCloudFormationType())
		results = append(results, r)
	}

	// Check for presence of egress rules
	if len(sg.SecurityGroupEgress) < 1 {
		r := NewResult(Warning, "best practice is to set explicit egress rules for security groups", sg.AWSCloudFormationType())
		results = append(results, r)
	}

	return results
}
