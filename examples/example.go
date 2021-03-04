package example

var (
	// unexportedVariable allows us to check that globocop picks up
	// unexported variables.
	unexportedVariable = "unexported"

	// ExportedVariable allows us to check that globocop picks up
	// exported variables.
	ExportedVariable = "exported"
)

const (
	// unexportedConstant allows us to check that globocop doesn't pick up
	// unexported constant.
	unexportedConstant = "unexported"

	// ExportedConstant allows us to check that globocop doesn't pick up
	// exported constant.
	ExportedConstant = "exported"
)

// unexportedType allows us to check that globocop doesn't pick up
// unexported types.
type unexportedType struct{}

// ExportedType allows us to check that globocop doesn't pick up
// exported types.
type ExportedType struct{}

// unexportedFunction allows us to check that globocop doesn't pick up
// identifiers in unexported functions.
func unexportedFunction() {
	// inside allows us to check that globocop doens't pick up
	// non-global variables.
	var insideVariable = 1
	_ = insideVariable

	// inside allows us to check that globocop doens't pick up
	// non-global constants.
	const insideConstant = 2
	_ = insideConstant
}

// ExportedFunction allows us to check that globocop doesn't pick up
// identifiers in exported functions.
func ExportedFunction() {
	// inside allows us to check that globocop doens't pick up
	// non-global variables.
	var insideVariable = 1
	_ = insideVariable

	// inside allows us to check that globocop doens't pick up
	//non-global constants.
	const insideConstant = 2
	_ = insideConstant
}
