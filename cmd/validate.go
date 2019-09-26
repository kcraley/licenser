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
	Run:   validateCmdFunc,
}

func init() {
	rootCmd.AddCommand(validateCmd)
}

func validateCmdFunc(cmd *cobra.Command, args []string) {
	var exitCode int = exitSuccess
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
	fmt.Println("Scanning files...")
	for _, file := range modified {
		fmt.Printf("%q\n", file)
	}

	results := license.Validate(modified, globalOpts.Author, globalOpts.License)
	for _, result := range results {
		if result.Error != nil {
			fmt.Printf("An error occurred validating files: %s, %s\n",
				result.FileName, result.Error)
		}

		if result.Validated == false {
			exitCode = exitNeedsHeaders
			fmt.Printf("Missing license headers in file: %q\n", result.FileName)
		} else {
			fmt.Printf("Headers exist in file: %q\n", result.FileName)
		}
	}
	os.Exit(exitCode)
}
