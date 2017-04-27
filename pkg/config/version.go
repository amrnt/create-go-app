package config

var (
	// Version of the build
	Version = "0.0.0-alpha"

	// CommitHash contains the current Git revision.
	// Use make to build to make sure this gets set.
	CommitHash = "dev-snapshot"

	// BuildDate contains the date of the current build.
	BuildDate string
)
