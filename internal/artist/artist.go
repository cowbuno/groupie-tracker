package artist

import (
	"fmt"
	"net/http"
)

var client *http.Client

type Artist struct {
	ID           int      `json:"id"`
	ImageURL     string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Location     Location
	Date         Date
	Relation     Relation
}

type Location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}

type Date struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

func (artist *Artist) PrintArtistDetails() {
	fmt.Printf("ID: %d\n", artist.ID)
	fmt.Printf("Name: %s\n", artist.Name)
	fmt.Printf("Image URL: %s\n", artist.ImageURL)
	fmt.Printf("Members: %v\n", artist.Members)
	fmt.Printf("Creation Date: %d\n", artist.CreationDate)
	fmt.Printf("First Album: %s\n", artist.FirstAlbum)
	fmt.Printf("Locations: %v\n", artist.Location.Locations)
	fmt.Printf("Dates: %v\n", artist.Date.Dates)
	fmt.Printf("Relations: %v\n", artist.Relation.DatesLocations)
	fmt.Println("--------------------------------------------------")
}

// func main() {
// 	artists := GetArtists()

// 	for i := range *artists {
// 		artist := &(*artists)[i]
// 		id := artist.ID
// 		artist.Location = *GetLocationByID(id)

// 		artist.Date = *GetDatesByID(id)

// 		artist.Relation = *GetRelationByID(id)
// 		PrintArtistDetails(artist)
// 	}

// }
