package helpers

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Education struct {
	SchoolName        string `json:"title"`
	Degree            string `json:"degree"`
	FieldOfStudy      string `json:"field"`
	SchoolDescription string `json:"description,omitempty"`
	// StartDate         string `json:"start_year"`
	// EndDate           string `json:"end_year"`
}

// Position represents a specific position within an experience.
type Position struct {
	Subtitle        string `json:"subtitle"`
	Meta            string `json:"meta"`
	Description     string `json:"description"`
	DescriptionHTML string `json:"description_html"`
	Duration        string `json:"duration"`
	StartDate       string `json:"start_date"`
	EndDate         string `json:"end_date"`
	DurationShort   string `json:"duration_short"`
	Title           string `json:"title"`
}

// Experience represents a single experience entry.
type Experience struct {
	Title           string     `json:"title"`
	DescriptionHTML *string    `json:"description_html"`
	Duration        string     `json:"duration"`
	StartDate       string     `json:"start_date"`
	EndDate         string     `json:"end_date"`
	DurationShort   string     `json:"duration_short"`
	Company         string     `json:"company"`
	CompanyID       string     `json:"company_id"`
	URL             string     `json:"url"`
	Location        *string    `json:"location"`
	Description     *string    `json:"description"`
	Positions       []Position `json:"positions,omitempty"`
}

// CurrentCompany represents the current company information.
type CurrentCompany struct {
	Name  string `json:"name"`
	Title string `json:"title"`
}

type Activity struct {
	Interaction string `json:"interaction"`
	Title       string `json:"title"`
}

type Post struct {
	Title string `json:"title"`
}

type LinkedInProfile struct {
	ID             string         `json:"id"`
	Name           string         `json:"name"`
	Location       string         `json:"city"`
	Title          string         `json:"position"`
	About          string         `json:"about"`
	CurrentCompany CurrentCompany `json:"current_company"`
	Followers      int            `json:"followers"`
	Connections    int            `json:"connections"`

	Posts      []Post       `json:"posts,omitempty"`
	Activity   []Activity   `json:"activity,omitempty"`
	Experience []Experience `json:"experience,omitempty"`
	Education  []Education  `json:"education,omitempty"`
}

func ParseLinkedInDataForAI(data []byte) (LinkedInProfile, error) {

	if data == nil {
		return LinkedInProfile{}, fmt.Errorf("data is nil")
	}

	// Unmarshal the JSON data into a struct
	var linkedinProfile LinkedInProfile

	err := json.Unmarshal(data, &linkedinProfile)
	if err != nil {
		return LinkedInProfile{}, err
	}

	//  limit the experience to latest 3
	linkedinProfile.Experience = linkedinProfile.Experience[:3] // TODO: make this dynamic

	//  limit the education to latest 2
	linkedinProfile.Education = linkedinProfile.Education[:2] // TODO: make this dynamic

	//  run loop to remove "by {name}" from the interaction & limit to 5
	linkedinProfile.Activity = linkedinProfile.Activity[:5] // TODO: make this dynamic
	for i, activity := range linkedinProfile.Activity {
		linkedinProfile.Activity[i].Interaction = strings.Replace(activity.Interaction, "by "+linkedinProfile.Name, "", 1)
	}

	return linkedinProfile, nil
}
