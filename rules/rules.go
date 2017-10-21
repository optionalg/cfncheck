package rules

// State of the result: Success, Warning, Failure
type State int

// Possible State values
const (
	Success State = iota
	Warning
	Failure
)

// Result of running a rule against a resource
type Result struct {
	description string
	state       State
}

// NewResult creates and initialises a new Result
func NewResult(state State, description string) Result {
	var result Result
	result.SetState(state)
	result.SetDescription(description)
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

// SetDescription sets the description value for the result
func (r *Result) SetDescription(description string) {
	r.description = description
}

// SetState sets the state value for the result
func (r *Result) SetState(state State) {
	r.state = state
}
