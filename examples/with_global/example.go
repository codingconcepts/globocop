package example

var (
	// unexportedVariable allows us to check that globocop picks up unexported variables.
	unexportedVariable = "unexported"

	// ExportedVariable allows us to check that globocop picks up exported variables.
	ExportedVariable = "exported"
)

const (
	// unexportedConstant allows us to check that globocop picks up unexported constant.
	unexportedConstant = "unexported"

	// ExportedConstant allows us to check that globocop picks up exported constant.
	ExportedConstant = "exported"
)

// unexportedFunction allows us to check that globocop doesn't pick up unexported functions.
func unexportedFunction() {
	// inside allows us to check that globocop doens't pick up non-global variables.
	var insideVariable = 1
	_ = insideVariable

	// inside allows us to check that globocop doens't pick up non-global constants.
	const insideConstant = 2
	_ = insideConstant
}

// ExportedFunction allows us to check that globocop doesn't pick up exported functions.
func ExportedFunction() {
	// inside allows us to check that globocop doens't pick up non-global variables.
	var insideVariable = 1
	_ = insideVariable

	// inside allows us to check that globocop doens't pick up non-global constants.
	const insideConstant = 2
	_ = insideConstant
}
