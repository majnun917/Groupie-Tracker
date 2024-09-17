package api

import "encoding/json"

// Aggregates data from all API structures.
type Data struct {
	A Artist
	R Relation
	L Location
	D Date
}

// Saves data retrieved from the artist API.
type Artist struct {
	Id           uint     `json:"id"`
	Name         string   `json:"name"`
	Image        string   `json:"image"`
	Members      []string `json:"members"`
	CreationDate uint     `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

// Saves data retrieved from the location API.
type Location struct {
	Locations []string `json:"locations"`
}

// Saves data retrieved from the date API.
type Date struct {
	Dates []string `json:"dates"`
}

// Saves data retrieved from the relation API.
type Relation struct {
	DatesLocations map[string][]string `json:"datesLocations"`
}

// Uses slices of structs to index each artist's data from the APIs.
// Uses map[string]json.RawMessage for parsing nested layers when multiple levels are involved.
var (
	artistInfo   []Artist
	locationMap  map[string]json.RawMessage
	locationInfo []Location
	datesMap     map[string]json.RawMessage
	datesInfo    []Date
	relationMap  map[string]json.RawMessage
	relationInfo []Relation
)
