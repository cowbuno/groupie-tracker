package main

import (
	"fmt"
	"groupie/internal/artist"
)

func main() {
	arts, err := artist.GetArtists()
	fmt.Println(err)
	for _, artist := range *arts {
		artist.PrintArtistDetails()
	}
}
