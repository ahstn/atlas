package builder

//go:generate mockery -name Builder -case underscore

// Builder defines the actions that an app packager typically preforms.
// i.e "mvn clean install", "go install"
//
// Each builder should have an external "NewClient()" method that allows users
// to initialise a Builder easily with sensible defaults.
type Builder interface {
	// Run is the method that executes the builder, assuming it has been
	// initialised correctly
	Run(bool) error

	// Args allows users to quickly have access to the arguments
	// the builder is planning to execute
	Args() string

	// ModifyArgs allows users to edit the commands arguments
	ModifyArgs([]string)
}
