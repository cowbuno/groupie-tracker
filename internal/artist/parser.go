package artist

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func GetJson(url string, target interface{}) error {
	resp, err := client.Get(url)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}

func GetArtists() (*[]Artist, error) {
	client = &http.Client{Timeout: 10 * time.Second}

	url := "https://groupietrackers.herokuapp.com/api/artists"
	var artists []Artist
	err := GetJson(url, &artists)
	if err != nil {
		log.Fatal(err)
	}

	for i := range artists {
		artist := artists[i]
		id := artist.ID
		Location, err := GetLocationByID(id)
		if err != nil {
			return nil, err
		}
		artist.Location = *Location

		Date, err := GetDatesByID(id)
		if err != nil {
			return nil, err
		}
		artist.Date = *Date

		Relation, err := GetRelationByID(id)
		if err != nil {
			return nil, err
		}
		artist.Relation = *Relation

	}
	return &artists, nil

}

func GetLocationByID(ID int) (*Location, error) {
	url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/locations/%d", ID)

	var location Location
	err := GetJson(url, &location)
	if err != nil {
		return nil, err
	}
	return &location, nil
}

func GetDatesByID(ID int) (*Date, error) {
	url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/dates/%d", ID)

	var date Date
	err := GetJson(url, &date)
	if err != nil {
		return nil, err

	}
	return &date, nil

}

func GetRelationByID(ID int) (*Relation, error) {
	url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/relation/%d", ID)

	var relation Relation

	err := GetJson(url, &relation)
	if err != nil {
		return nil, err
	}
	return &relation, nil
}
