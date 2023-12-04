package api

import (
	"encoding/json"
	"fmt"
	"groupie/pkg/models"
	"net/http"
	"time"
)

var client *http.Client

const url = "https://groupietrackers.herokuapp.com/api/artists"

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

func GetArtists() (*[]models.AllArtists, error) {

	var artists []models.AllArtists
	err := GetJson(url, &artists)
	if err != nil {
		return nil, err
	}

	return &artists, nil
}

func GetArtist(id string) (*[]models.Artist, error) {
	var artists []models.Artist
	urlId := url + "/" + id
	fmt.Print(urlId)
	err := GetJson(urlId, &artists)
	if err != nil {
		return nil, err
	}

	return &artists, nil
}

// func GetLocationByID(ID int) (*Location, error) {
// 	url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/locations/%d", ID)

// 	var location Location
// 	err := GetJson(url, &location)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &location, nil
// }

// func GetDatesByID(ID int) (*Date, error) {
// 	url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/dates/%d", ID)

// 	var date Date
// 	err := GetJson(url, &date)
// 	if err != nil {
// 		return nil, err

// 	}
// 	return &date, nil

// }

// func GetRelationByID(ID int) (*Relation, error) {
// 	url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/relation/%d", ID)

// 	var relation Relation

// 	err := GetJson(url, &relation)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &relation, nil
// }
