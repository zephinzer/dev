package constants

const (
	// CheckSuccessFormat is the standardised format used to print messages
	// where the check was successful
	CheckSuccessFormat = "\033[32m✅\033[0m \033[32m%s\033[0m "

	// CheckFailureFormat is the standardised format used to print messsages
	// where the check failed
	CheckFailureFormat = "\033[31m❌ \033[31m%s\033[0m "

	// CheckSkippedFormat is the standardised format used to print messages
	// where the check was skipped
	CheckSkippedFormat = "\033[33m👀\033[0m \033[33m%s\033[0m "
)
