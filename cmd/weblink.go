/*
Package cmd provides the command-line interface for interacting with Capacities.

Copyright © 2024 Guilherme Thomazi Bonicontro <thomazi@linux.com>
*/
package cmd

import (
	"github.com/guitmz/capable/capacities"
	"github.com/guitmz/capable/utils"
	"github.com/spf13/cobra"
)

// weblinkCmd represents the weblink command.
var weblinkCmd = &cobra.Command{
	Use:   "weblink",
	Short: "Save a weblink to a space",
	Long:  `The Weblink will be analyzed and saved as an object to your space.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Retrieve Capacities authentication details.
		capacitiesSpaceId, capacitiesApiToken := utils.GetCapacitiesAuth()

		// Retrieve input flags.
		url, err := cmd.Flags().GetString("url")
		utils.CheckError(err)

		title, err := cmd.Flags().GetString("title")
		utils.CheckError(err)

		description, err := cmd.Flags().GetString("description")
		utils.CheckError(err)

		userInput, err := cmd.Flags().GetString("input")
		utils.CheckError(err)

		tags, err := cmd.Flags().GetStringArray("tags")
		utils.CheckError(err)

		debug, err := cmd.Flags().GetBool("debug")
		utils.CheckError(err)

		// Save the weblink to the specified space.
		capacities.SaveWeblink(capacitiesSpaceId, capacitiesApiToken, url, title, description, userInput, tags, debug)
	},
}

// Initialize the command flags.
func init() {
	rootCmd.AddCommand(weblinkCmd)
	weblinkCmd.Flags().StringP("url", "u", "", "URL of the weblink")
	weblinkCmd.Flags().StringP("title", "t", "", "Title of the weblink. If not specified, it will be fetched from the URL.")
	weblinkCmd.Flags().StringP("description", "d", "", "Description of the weblink. If not specified, it will be fetched from the URL.")
	weblinkCmd.Flags().StringP("input", "i", "", "Text formatted as markdown to be added to the notes section.")
	weblinkCmd.Flags().StringArrayP("tags", "", []string{}, "Tags to add to the weblink. Tags must match existing tag names in Capacities.")
	weblinkCmd.Flags().BoolP("debug", "", false, "Display the return JSON object from Capacities.io API.")
}
