package rules

// State of the result: Success, Warning, Failure
type State int

// Possible State values
const (
	Success State = iota
	Warning
	Failed
)

// Result of running a rule against a resource
type Result struct {
	description string
	state       State
	awsType     string
}

// NewResult creates and initialises a new Result
func NewResult(state State, description string, awsType string) Result {
	var result Result
	result.setState(state)
	result.setDescription(description)
	result.setAwsType(awsType)
	return result
}

// Description returns the description of the result
func (r Result) Description() string {
	return r.description
}

// State returns the state of the result
func (r Result) State() State {
	return r.state
}

// AwsType returns the AWS CloudFormation Type of the resource that caused the result
func (r Result) AwsType() string {
	return r.awsType
}

func (r *Result) setDescription(description string) {
	r.description = description
}

func (r *Result) setState(state State) {
	r.state = state
}

func (r *Result) setAwsType(awsType string) {
	r.awsType = awsType
}
