package handlers

import (
	"groupie/internal/artist"
	"log"
	"net/http"
)

func mainPage(result http.ResponseWriter, call *http.Request) {

}

func StartPage() {
	arts, err := artist.GetArtists()
	if err != nil {
		log.Fatal(err)
	}
	aplication := &
}
