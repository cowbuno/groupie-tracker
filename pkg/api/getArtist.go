package api

import (
	"encoding/json"
	"fmt"
	"groupie/pkg/models"
	"net/http"
	"time"
)

var client *http.Client

const urlArtist = "https://groupietrackers.herokuapp.com/api/artists"
const urlLocation = "https://groupietrackers.herokuapp.com/api/locations"
const urlRelation = "https://groupietrackers.herokuapp.com/api/relation"

func init() {
	client = &http.Client{Timeout: 10 * time.Second}
}

func GetJson(url string, target interface{}) error {
	resp, err := client.Get(url)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}

func GetAllArtists() (*[]models.Artist, error) {

	type ArtistSubset struct {
		ID           int    `json:"id"`
		Image        string `json:"image"`
		Name         string `json:"name"`
		CreationDate int    `json:"creationDate"`
	}

	var artistsSubset []ArtistSubset
	err := GetJson(urlArtist, &artistsSubset)
	if err != nil {
		return nil, err
	}
	var artists []models.Artist

	for _, subset := range artistsSubset {
		artist := models.Artist{
			ID:           subset.ID,
			Image:        subset.Image,
			Name:         subset.Name,
			CreationDate: subset.CreationDate,
			// Initialize other fields as needed
		}
		artists = append(artists, artist)
	}

	return &artists, nil
}

func GetArtistByID(ID string) (*models.Artist, error) {
	var artist models.Artist
	urlId := urlArtist + "/" + ID
	err := GetJson(urlId, &artist)
	if err != nil {
		return nil, err
	}

	location, err := GetLocationByID(ID)
	if err != nil {
		return nil, err
	}
	artist.Location = location

	relation, err := GetRelationByID(ID)
	fmt.Println(relation)
	if err != nil {
		return nil, err
	}
	artist.Relation = relation

	return &artist, nil
}

func GetLocationByID(ID string) (*models.Location, error) {
	url := urlLocation + "/" + ID
	fmt.Println(url)
	var location models.Location
	err := GetJson(url, &location)
	if err != nil {
		return nil, err
	}
	return &location, nil
}

// func GetDatesByID(ID int) (*Date, error) {
// 	url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/dates/%d", ID)

// 	var date Date
// 	err := GetJson(url, &date)
// 	if err != nil {
// 		return nil, err

// 	}
// 	return &date, nil

// }

func GetRelationByID(ID string) (*models.Relation, error) {
	url := urlRelation + "/" + ID
	var relation models.Relation

	err := GetJson(url, &relation)
	if err != nil {
		return nil, err
	}
	return &relation, nil
}
