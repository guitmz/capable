/*
Package cmd provides functionality to interact with daily notes in Capacities.

Copyright © 2024 Guilherme Thomazi Bonicontro <thomazi@linux.com>
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
	Short: "Search for content",
	Long:  `Returns content based on a search term in a set of spaces.`,
	Run: func(cmd *cobra.Command, args []string) {
		_, capacitiesApiToken := utils.GetCapacitiesAuth()

		debug, err := cmd.Flags().GetBool("debug")
		utils.CheckError(err)

		mode, err := cmd.Flags().GetString("mode")
		utils.CheckError(err)

		query, err := cmd.Flags().GetString("query")
		utils.CheckError(err)

		spaces, err := cmd.Flags().GetStringArray("spaces")
		utils.CheckError(err)

		filters, err := cmd.Flags().GetStringArray("filters")
		utils.CheckError(err)

		// Search for content in specified spaces.
		capacities.SearchContent(capacitiesApiToken, mode, query, spaces, filters, debug)
	},
}

// Initialize the command flags.
func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.Flags().StringP("mode", "m", "title", "Search either by title or fullText")
	searchCmd.Flags().StringP("query", "q", "", "The search term")
	searchCmd.Flags().StringArrayP("spaces", "s", []string{}, "IDs of the spaces to search in. Example: \"c722edd0-7cee-xxxx-xxxx-27915dcdbf3b, yyyyyyyy-7cee-xxxx-xxxx-27915dcdbf3b\"")
	searchCmd.Flags().StringArrayP("filters", "f", []string{}, "Filter for a set of structureIds. If not specified, all content will be searched. Example: \"Tag1,Tag2\"")
	searchCmd.Flags().BoolP("debug", "", false, "Display the return JSON object from Capacities.io API.")
}
