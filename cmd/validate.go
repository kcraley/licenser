package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	validateUse   = "validate"
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
	fmt.Println("validate called")
}
