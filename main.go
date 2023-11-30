package main

import (
	"github.com/ikediufomadu/Go_Spotify_Script/Handler"
	"github.com/ikediufomadu/Go_Spotify_Script/Spotify/Authenticator"
	"github.com/ikediufomadu/Go_Spotify_Script/Spotify/Playlist"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Website the program will get music data from
	scraper := "https://pitchfork.com/reviews/best/tracks/?page=2"

	// Start local server for Spotify authentication
	Authenticator.StartLocalServer()

	// Connect to the website to get music data
	Handler.WebConnector(scraper)

	// Perform a specific Spotify query based on user selection
	Handler.WhichQuery(1)
	Playlist.CreatePlaylist("Test", "", true, true)
}
