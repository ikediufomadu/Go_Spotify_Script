package main

import (
	"github.com/ikediufomadu/Go_Spotify_Script/Handler"
	"github.com/ikediufomadu/Go_Spotify_Script/Query"
	"github.com/ikediufomadu/Go_Spotify_Script/Spotify/Authenticator"
)

func main() {
	// Website the program will get music data from
	scraper := "https://pitchfork.com/reviews/best/tracks/?page=1"

	// Start local server for Spotify authentication
	Authenticator.StartLocalServer()

	// Connect to the website to get music data
	Handler.WebConnector(scraper)

	// Perform a specific Spotify query based on user selection
	Query.GenreQuery("rock")
}
