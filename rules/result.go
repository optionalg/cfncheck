package rules

// State of the result: Pass, Warn, Fail
type State int

// Possible State values
const (
	Pass State = iota
	Warn
	Fail
)

// Result of running a rule against a resource
type Result struct {
	resource    string
	cfnType     string
	description string
	result      State
}

// NewResult creates and initialises a new Result
func NewResult(resource string, cfnType string, result State, description string) Result {
	var res Result
	res.setResult(result)
	res.setDescription(description)
	res.setAwsType(cfnType)
	return res
}

// Description returns the description of the result
func (r Result) Description() string {
	return r.description
}

// Result returns the result of the result
func (r Result) Result() State {
	return r.result
}

// AwsType returns the AWS CloudFormation Type of the resource that caused the result
func (r Result) AwsType() string {
	return r.cfnType
}

// Resource returns the name of the CloudFormation resource
func (r Result) Resource() string {
	return r.resource
}

func (r *Result) setDescription(description string) {
	r.description = description
}

func (r *Result) setResult(result State) {
	r.result = result
}

func (r *Result) setAwsType(cfnType string) {
	r.cfnType = cfnType
}

func (r *Result) setResource(resource string) {
	r.resource = resource
}
