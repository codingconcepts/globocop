package example

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
