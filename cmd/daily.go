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

// dailyCmd represents the daily command
var dailyCmd = &cobra.Command{
	Use:   "daily",
	Short: "Save text to today's daily note",
	Long: `Saves a text to today's daily note in a space. The text can be formatted in markdown.

Do not use this to import large amounts of content as this could break the daily note.
Use the import system of Capacities instead.`,
	Run: func(cmd *cobra.Command, args []string) {
		capacitiesSpaceId, capacitiesApiToken := utils.GetCapacitiesAuth()

		timestamp, err := cmd.Flags().GetBool("timestamp")
		utils.CheckError(err)

		origin, err := cmd.Flags().GetString("origin")
		utils.CheckError(err)

		userInput, err := cmd.Flags().GetString("input")
		utils.CheckError(err)

		if userInput == "stdin" {
			userInput, err = utils.ReadFromStdin()
			utils.CheckError(err)
		}
		// Save the daily note to the specified space.
		capacities.SaveToDailyNote(capacitiesSpaceId, capacitiesApiToken, userInput, origin, timestamp, false)
	},
}

// Initialize the command flags.
func init() {
	rootCmd.AddCommand(dailyCmd)
	dailyCmd.Flags().BoolP("timestamp", "t", true, "If false, no time stamp will be added to the note")
	dailyCmd.Flags().StringP("input", "i", "stdin", "Markdown text that should be added to today's daily note. Reads from STDIN first")
	dailyCmd.Flags().StringP("origin", "o", "commandPalette", "The origin of the save action. This determines the icon added to your note")
}
