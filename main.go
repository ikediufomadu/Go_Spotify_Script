package main

import (
	"github.com/ikediufomadu/Go_Spotify_Script/Handler"
	"github.com/ikediufomadu/Go_Spotify_Script/Spotify/Authenticator"
)

func main() {
	// Website the program will get music data from
	scraper := "https://www.billboard.com/charts/hot-100/"

	// Start local server for Spotify authentication
	Authenticator.StartLocalServer()

	// Connect to the website to get music data
	Handler.WebConnector(scraper)
}
