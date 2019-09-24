package version

import "fmt"

var (
	// GitCommit is the SHA of the current commit
	GitCommit = "UNKNOWN"
	// GitStatus is the current state of the repository
	GitStatus = "UNKNOWN"
	// Release is the SemVer release number
	Release = "UNKNOWN"
	// Repo is the URL of the upstream repository
	Repo = "UNKNOWN"
)

// Print writes the binary version to stdout
func Print() string {
	return fmt.Sprintf("%s\n commit: git-%s-%s\n repo: %s", Release, GitCommit, GitStatus, Repo)
}
