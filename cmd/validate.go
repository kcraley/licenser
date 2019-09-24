package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "validates headers exist in source files given a path",
	Long: `Validate:

Will validate that the source code within a given path
have the necessary licenser header.  If the validation
fails, the files missing the license header will be
output and a non-zero exit code will be returned`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("validate called")
	},
}

func init() {
	rootCmd.AddCommand(validateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// validateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// validateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
