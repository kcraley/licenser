package cmd

const (
	// Default values for Options
	appName            = "licenser"
	defaultAuthor      = "UNKNOWN"
	defaultDryRun      = false
	defaultLicense     = "UNKNOWN"
	defaultShowVersion = false
	defaultVerbose     = false
)

const (
	// Exit return codes
	exitSuccess = iota
	exitNeedsHeaders
	exitAccessFailure
)

var (
	defaultPath    = []string{"."}
	defaultExclude = []string{
		".git",
		".gitignore",
		".vscode",
		"LICENSE",
		"NOTICE",
		"README.md",
	}
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
	opts := &Options{
		Author:      defaultAuthor,
		DryRun:      defaultDryRun,
		Exclude:     defaultExclude,
		License:     defaultLicense,
		ShowVersion: defaultShowVersion,
		Verbose:     defaultVerbose,
	}
	return opts
}
