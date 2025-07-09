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
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/guitmz/capable/utils"
	"github.com/olekukonko/tablewriter"
)

// doRequest sends an HTTP POST request with the provided payload to a specified endpoint.
// It handles the response, logging any errors, and optionally prints the result in a pretty JSON format if debug is true.
func doRequest(payload any, endpoint, token string, debug bool) []byte {
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

	defer resp.Body.Close()

	if resp.Body != nil {
		bodyBytes, err := io.ReadAll(resp.Body)
		utils.CheckError(err)
		return bodyBytes
	}

	return nil
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
	weblinkResponse := &WeblinkResponse{}
	result := doRequest(weblink, "/save-weblink", capacitiesToken, debug)
	if err := json.Unmarshal(result, weblinkResponse); err != nil {
		log.Println("cannot unmarshal response JSON to WeblinkResponse:", err)
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"ID", "StructureID", "Title", "Description", "Tags"})
	table.Append([]string{weblinkResponse.ID, weblinkResponse.StructureID, weblinkResponse.Title, weblinkResponse.Description})
	table.Render()
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

// SearchContent searches for content based on user input, mode, and optional filters making a request to the "/search" endpoint.
func SearchContent(capacitiesSpaceId, capacitiesToken, userInput, mode string, filter []string, debug bool) {
	search := SearchContentRequest{
		SpaceIDs:           []string{capacitiesSpaceId},
		Mode:               mode,
		SearchTerm:         userInput,
		FilterStructureIds: filter,
	}
	searchResult := &SearchResult{}
	result := doRequest(search, "/search", capacitiesToken, debug)
	if err := json.Unmarshal(result, searchResult); err != nil {
		log.Println("cannot unmarshal response JSON to SearchResult:", err)
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"ID", "Title", "StructureID", "Highlight Snippet"})
	for _, result := range searchResult.Results {
		snippet := ""
		if len(result.Highlights) > 0 && len(result.Highlights[0].Snippets) > 0 {
			snippet = result.Highlights[0].Snippets[0]
		}
		table.Append([]string{result.ID, result.Title, result.StructureID, snippet})
	}
	table.Render()
}
