package constants

const (
	// DevTimeFormat is the printed time format for messages intended for logging from the `dev` application
	DevTimeFormat = "20060102_150405"

	// DevHumanTimeFormat is the printed time format for messages intended for human reading from the `dev` application
	DevHumanTimeFormat = "2006-01-02 15:04:05"

	// DateOnlyTimeFormat is the time format when printing just the date
	DateOnlyTimeFormat = "2006-01-02"

	// GithubAPITimeFormat defines the timestamp format of timestamps returned by the Github API
	GithubAPITimeFormat = "2006-01-02T15:04:05.99Z"

	// GitlabAPITimeFormat defines the timestamp format of timestamps returned by the Gitlab API
	GitlabAPITimeFormat = "2006-01-02T15:04:05.999Z"

	// PivotalTrackerAPITimeFormat defines the timestampformat of timestamps returned by the Pivotal Tracker API
	PivotalTrackerAPITimeFormat = "2006-01-02T15:04:05Z"
)
