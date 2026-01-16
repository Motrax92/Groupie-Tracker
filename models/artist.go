package models

type Artist struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Image        string   `json:"image"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Members      []string `json:"members"`

	LocationsURL string `json:"locations"`
	DatesURL     string `json:"concertDates"`
	RelationURL  string `json:"relations"`
}
