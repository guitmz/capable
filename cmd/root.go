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
	Long: `Capable is CLI tool is designed to interact with the Capacities.io API, 
enabling users to send daily entries and weblinks to the Capacities platform 
directly from the command line.`,
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
