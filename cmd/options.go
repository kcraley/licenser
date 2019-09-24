package cmd

const (
	appName            = "licenser"
	defaultAuthor      = "UNKNOWN"
	defaultDryRun      = false
	defaultLicense     = "UNKNOWN"
	defaultShowVersion = false
	defaultVerbose     = false
)

var (
	defaultExclude = []string{}
)

// Options defines the command line interface flags
type Options struct {
	Author      string
	DryRun      bool
	Exclude     []string
	License     string
	ShowVersion bool
	Verbose     bool
}

// NewOptions creates and returns new options
func NewOptions() *Options {
	opts := Options{
		Author:      defaultAuthor,
		DryRun:      defaultDryRun,
		Exclude:     defaultExclude,
		License:     defaultLicense,
		ShowVersion: defaultShowVersion,
		Verbose:     defaultVerbose,
	}
	return &opts
}
