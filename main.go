package main

import (
	"github.com/ikediufomadu/Go_Spotify_Script/Handler"
	"github.com/ikediufomadu/Go_Spotify_Script/Spotify/Authenticator"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	scraper := "https://pitchfork.com/reviews/best/tracks/?page=2"

	Authenticator.StartLocalServer()
	Handler.WebConnector(scraper)
	Handler.WhichQuery(1)
	//Playlist.CreatePlaylist("Test Playlist", "test", false, false)
}
