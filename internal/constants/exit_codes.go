package constants

const (
	// ExitOK indicates there are no errors
	ExitOK = 0

	// ExitErrorSystem indicates the error is because of the system
	ExitErrorSystem = 1

	// ExitErrorUser indicates the error is because of the user
	ExitErrorUser = 2

	// ExitErrorInput indicates the error is because of the user's input
	ExitErrorInput = 4

	// ExitErrorConfiguration indicates the error is because of the
	// configuration
	ExitErrorConfiguration = 8

	// ExitErrorApplication indicates the error is a logical bug with the
	// code
	ExitErrorApplication = 16

	// ExitErrorValidation indicates the error is related to validation
	ExitErrorValidation = 32
)
