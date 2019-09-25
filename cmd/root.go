package cmd

import (
	"fmt"
	"os"

	"github.com/kcraley/licenser/pkg/version"

	"github.com/spf13/cobra"
)

const (
	rootUse   = "licenser"
	rootShort = "Licenser validates, adds or modifies the license headers of source code"
	rootLong  = `
Licenser:

A command line interface tools which can validate,
add or modify the headers of source code.`
)

var (
	globalOpts  = NewOptions()
	rootVersion = version.Print()
	rootCmd     = &cobra.Command{
		Use:     rootUse,
		Short:   rootShort,
		Long:    rootLong,
		Version: rootVersion,
	}
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&globalOpts.Author, "author", "a", globalOpts.Author, "the author to be used for the license header")
	rootCmd.PersistentFlags().StringArrayVarP(&globalOpts.Exclude, "exclude", "e", globalOpts.Exclude, "an arrary of paths to exclude")
	rootCmd.PersistentFlags().StringVarP(&globalOpts.License, "license", "l", globalOpts.License, "the license to use for all source code")
	rootCmd.PersistentFlags().BoolVarP(&globalOpts.Verbose, "verbose", "v", globalOpts.Verbose, "displays verbose output when executing")
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
