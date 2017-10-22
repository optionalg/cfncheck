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

	// Check for 0.0.0.0/0 in IPv4 CIDR blocks
	if len(sg.SecurityGroupIngress) > 0 {
		for _, rule := range sg.SecurityGroupIngress {
			if rule.CidrIp == "0.0.0.0/0" {
				r := NewResult(Failed, "security group rules should not use open 0.0.0.0/0 IPv4 blocks", rule.AWSCloudFormationType())
				results = append(results, r)
			}
		}
	}
	if len(sg.SecurityGroupEgress) > 0 {
		for _, rule := range sg.SecurityGroupEgress {
			if rule.CidrIp == "0.0.0.0/0" {
				r := NewResult(Failed, "security group rules should not use open 0.0.0.0/0 IPv4 blocks", rule.AWSCloudFormationType())
				results = append(results, r)
			}
		}
	}

	// Check for ::/0 in IPv6 CIDR blocks
	if len(sg.SecurityGroupIngress) > 0 {
		for _, rule := range sg.SecurityGroupIngress {
			if rule.CidrIp == "::/0" || rule.CidrIp == "0:0:0:0:0:0:0:0/0" {
				r := NewResult(Failed, "security group rules should not use open ::/0 IPv6 blocks", rule.AWSCloudFormationType())
				results = append(results, r)
			}
		}
	}
	if len(sg.SecurityGroupEgress) > 0 {
		for _, rule := range sg.SecurityGroupEgress {
			if rule.CidrIp == "::/0" || rule.CidrIp == "0:0:0:0:0:0:0:0/0" {
				r := NewResult(Failed, "security group rules should not use open ::/0 in IPv6 blocks", rule.AWSCloudFormationType())
				results = append(results, r)
			}
		}
	}

	return results
}
