/*
Copyright © 2024 Guilherme Thomazi Bonicontro <thomazi@linux.com>
*/
package capacities

const (
	baseURL = "https://api.capacities.io"
)

// DailyNoteEntry represents a daily note entry.
type DailyNoteEntry struct {
	SpaceID     string `json:"spaceId"`
	MdText      string `json:"mdText"`
	Origin      string `json:"origin"`
	NoTimeStamp bool   `json:"noTimeStamp"`
}

// Weblink represents a web link entry.
type Weblink struct {
	SpaceID              string   `json:"spaceId"`
	URL                  string   `json:"url"`
	TitleOverwrite       string   `json:"titleOverwrite"`
	DescriptionOverwrite string   `json:"descriptionOverwrite"`
	Tags                 []string `json:"tags"`
	MdText               string   `json:"mdText"`
}

// WeblinkResponse represents a response for a web link entry.
type WeblinkResponse struct {
	SpaceID     string   `json:"spaceId"`
	ID          string   `json:"id"`
	StructureID string   `json:"structureId"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}
