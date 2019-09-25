package cmd

import (
	"fmt"
	"os"

	"github.com/kcraley/licenser/pkg/license"
	"github.com/kcraley/licenser/pkg/path"
	"github.com/spf13/cobra"
)

const (
	validateUse   = "validate [path]"
	validateShort = "validates a header exists for source code files"
	validateLong  = `
Validate:

Validate walks the directory structure and verifies the license
headers exist in the source code.  If the headers do not exist,
the command will exit with a non-zero return code.  If used with
the --exclude/-e flag, those paths will be excludes from the
validation.`
)

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use:   validateUse,
	Short: validateShort,
	Long:  validateLong,
	Args:  cobra.MinimumNArgs(1),
	Run:   validateCmdFunc,
}

func init() {
	rootCmd.AddCommand(validateCmd)
}

func validateCmdFunc(cmd *cobra.Command, args []string) {
	var paths []string
	if len(args) > 0 {
		paths = args
	} else {
		paths = defaultPath
	}

	for _, v := range defaultExclude {
		globalOpts.Exclude = append(globalOpts.Exclude, v)
	}

	modified, err := path.Walk(paths, globalOpts.Exclude, args)
	if err != nil {
		fmt.Printf("An error occurred: %v", err)
	}
	fmt.Printf("%q", modified)

	valid, err := license.ValidateLicense(modified, globalOpts.License)
	if err != nil {
		fmt.Printf("An error occurred validating files: %s", err)
		os.Exit(1)
	}

	if !valid {
		fmt.Printf("The source code does not have the correct license headers")
		os.Exit(1)
	}
}
