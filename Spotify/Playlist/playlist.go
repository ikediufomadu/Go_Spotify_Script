package Playlist

import (
	"context"
	"fmt"
	"github.com/ikediufomadu/Go_Spotify_Script/Spotify/Authenticator"
	"github.com/zmb3/spotify/v2"
	"log"
)

// CreatePlaylist creates a new Spotify playlist for the authenticated user
func CreatePlaylist(playlistName, playlistDescription string, public, collaborative bool) {
	client, clientError := Authenticator.GetClient()
	isClientNull(clientError)

	userID, userIDError := Authenticator.GetUserID()
	if userIDError != nil {
		log.Fatal("User is nil", userIDError)
	}

	playlist, playlistError := client.CreatePlaylistForUser(context.Background(), userID, playlistName, playlistDescription, public, collaborative)
	if playlistError != nil {
		fmt.Println("There was a problem creating your playlist.", playlistError)
	} else {
		fmt.Println(playlist.ID + "\n PUT THIS STRING INTO THE getPlaylist METHOD.\nplaylist, playlistErr := client.GetPlaylist(context.Background(), \"STRING HERE\")")
	}
}

// AddTrackToPlaylist adds a track to the authenticated user's Spotify playlist
func AddTrackToPlaylist(track spotify.ID) {
	client, clientError := Authenticator.GetClient()
	isClientNull(clientError)

	fullPlaylist := getPlaylist()
	_, playlistError := client.AddTracksToPlaylist(context.Background(), fullPlaylist.ID, track)
	if playlistError != nil {
		fmt.Printf("There was a problem adding this song to the playlist %s: %v\n", fullPlaylist.Name, playlistError)
	} else {
		fmt.Printf("Song added to playlist.\n")
	}
}

func DeleteDuplicateTracks() {

}

func DeleteTrack() {

}

func getPlaylist() *spotify.FullPlaylist {
	client, clientError := Authenticator.GetClient()
	isClientNull(clientError)

	// Right now you will have to create a playlist then grab the playlist ID from your CLI, add that ID here.
	playlist, playlistErr := client.GetPlaylist(context.Background(), "6jQH6TbyhQrvb36ZBjLWQe")
	if playlistErr != nil {
		log.Fatal("Error grabbing playlist")
	}

	return playlist
}

// isClientNull checks if the Spotify client is nil and logs a fatal error if so
func isClientNull(clientError error) {
	if clientError != nil {
		log.Fatal("Client is nil", clientError)
	}
}
