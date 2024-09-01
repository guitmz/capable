/*
Package capacities provides functions for interacting with a remote service to save web links and daily notes.

Copyright © 2024 Guilherme Thomazi Bonicontro <thomazi@linux.com>
*/
package capacities

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/briandowns/spinner"
	"github.com/guitmz/capable/utils"
)

// doRequest sends an HTTP POST request with the provided payload to a specified endpoint.
// It handles the response, logging any errors, and optionally prints the result in a pretty JSON format if debug is true.
func doRequest(payload any, endpoint, token string, debug bool) {
	// Initialize spinner for visual indication of progress
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Suffix = " Sending payload...\n"
	s.FinalMSG = "Complete!\n"
	s.Start()

	// Marshal payload data to JSON
	payloadJson, err := json.Marshal(payload)
	utils.CheckError(err)

	// Create HTTP request
	req, err := http.NewRequest(http.MethodPost, baseURL+endpoint, bytes.NewBuffer(payloadJson))
	utils.CheckError(err)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	// Send HTTP request
	resp, err := http.DefaultClient.Do(req)
	utils.CheckError(err)

	// Check response status code; log error if not successful
	if resp.StatusCode != 200 {
		log.Fatalf("failed to send payload with status: %d", resp.StatusCode)
	}

	// Stop the spinner after completing the request
	s.Stop()

	// If debug mode is enabled, process and print the response in a pretty JSON format
	if debug {
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		utils.CheckError(err)
		var result WeblinkResponse
		if err := json.Unmarshal(body, &result); err != nil {
			log.Println("cannot unmarshal response JSON")
		}
		utils.PrettyPrintJSON(result)
	}
}

// SaveWeblink saves a web link with the provided details by making a request to the "/save-weblink" endpoint.
func SaveWeblink(capacitiesSpaceId, capacitiesToken, url, titleOverwrite, descriptionOverwrite, userInput string, tags []string, debug bool) {
	weblink := Weblink{
		SpaceID:              capacitiesSpaceId,
		URL:                  url,
		TitleOverwrite:       titleOverwrite,
		DescriptionOverwrite: descriptionOverwrite,
		MdText:               userInput,
		Tags:                 tags,
	}
	doRequest(weblink, "/save-weblink", capacitiesToken, debug)
}

// SaveToDailyNote saves a daily note entry with the provided details by making a request to the "/save-to-daily-note" endpoint.
func SaveToDailyNote(capacitiesSpaceId, capacitiesToken, userInput, origin string, noTimeStamp, debug bool) {
	daily := DailyNoteEntry{
		SpaceID:     capacitiesSpaceId,
		MdText:      userInput,
		Origin:      origin,
		NoTimeStamp: !noTimeStamp,
	}
	doRequest(daily, "/save-to-daily-note", capacitiesToken, debug)
}
