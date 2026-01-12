package api

type Artist struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Image        string   `json:"image"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type Locations struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}

type Dates struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Relation struct {
	ID            int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

// Modèle fusionné pour la GUI
type ArtistFull struct {
	Artist    Artist
	Locations Locations
	Dates     Dates
	Relation  Relation
}
