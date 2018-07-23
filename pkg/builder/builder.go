package builder

// Builder defines the actions that an app packager typically preforms.
// i.e "mvn clean install", "go install"
//
// Intended to be used as a builder pattern: Builder struct holds a exec.Cmd{}
//   initialiseCommand() will create the exec.Cmd if nil.
// Then functions like Clean() and Build() can be chained on the implmentation.
// The Run() func then executes the built exec.Cmd{}, tracking progress and
//   reporting an error back.
type Builder interface {
	//command exec.Cmd
	initialiseCommand()

	Clean()
	Build()
	Package()
	Run() error
}
