package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	addUse   = "add"
	addShort = "adds the license header to source code files"
	addLong  = `
Add:

Add will walk the directory structure and add the license
headers to source code.  If used with the --exclude/-e flag
those paths will be excluded and the header will not be added`
)

var addCmd = &cobra.Command{
	Use:   addUse,
	Short: addShort,
	Long:  addLong,
	Run:   addCmdFunc,
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func addCmdFunc(cmd *cobra.Command, args []string) {
	fmt.Println("add called")
}
