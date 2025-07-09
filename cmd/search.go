/*
Package cmd provides functionality to interact with daily notes in Capacities.

Copyright © 2025 Guilherme Thomazi Bonicontro <thomazi@linux.com>
*/
package cmd

import (
	"github.com/guitmz/capable/capacities"
	"github.com/guitmz/capable/utils"
	"github.com/spf13/cobra"
)

// searchCmd represents the daily command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search content based on a specified search term and space",
	Long:  `Returns content based on a search term in a specific space.`,
	Run: func(cmd *cobra.Command, args []string) {
		capacitiesSpaceId, capacitiesApiToken := utils.GetCapacitiesAuth()

		mode, err := cmd.Flags().GetString("mode")
		utils.CheckError(err)

		filter, err := cmd.Flags().GetStringArray("filter")
		utils.CheckError(err)

		userInput, err := cmd.Flags().GetString("search")
		utils.CheckError(err)

		if userInput == "stdin" {
			userInput, err = utils.ReadFromStdin()
			utils.CheckError(err)
		}
		// Save the daily note to the specified space.
		capacities.SearchContent(capacitiesSpaceId, capacitiesApiToken, userInput, mode, filter, true)
	},
}

// Initialize the command flags.
func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.Flags().StringP("mode", "m", "title", "The mode of search. Options are: 'title' or 'fullText'.")
	searchCmd.Flags().StringArrayP("filter", "f", []string{}, "Optionally filter by a list of tags. If not specified, all structures will be searched.")
	searchCmd.Flags().StringP("search", "s", "stdin", "Search term to query Capacities content. Reads from STDIN first")
}
