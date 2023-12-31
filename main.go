package main

import (
	"github.com/ikediufomadu/Go_Spotify_Script/Handler/Web"
	"github.com/ikediufomadu/Go_Spotify_Script/Spotify/Authenticator"
)

func main() {
	// Website the program will get music data from
	scraper := "https://www.billboard.com/charts/billboard-global-200/"

	// Start local server for Spotify authentication
	Authenticator.StartLocalServer()

	// Put the name of the playlist you want to add tracks to
	programCore()

	// Connect to the website to get music data
	if finished {
		Web.WebConnector(scraper)
	}

}
