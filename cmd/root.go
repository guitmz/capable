/*
Package cmd provides functionality for interacting with the Capacities.io API through a CLI application.

Author: Guilherme Thomazi Bonicontro <thomazi@linux.com>
*/

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "capable",
	Short: "Interact with Capacities.io API",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

// Execute runs the root command and exits the application if an error occurs
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

// SetVersionInfo sets the version information for the application
func SetVersionInfo(version, commit, date string) {
	rootCmd.Version = fmt.Sprintf("%s (built on %s from Git SHA %s)", version, date, commit)
}
