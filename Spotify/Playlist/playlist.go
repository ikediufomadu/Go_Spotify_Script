package Playlist

import (
	"context"
	"fmt"
	"github.com/ikediufomadu/Go_Spotify_Script/Spotify/Authenticator"
	"github.com/zmb3/spotify/v2"
	"log"
)

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
		fmt.Println(playlist.ID)
	}
}

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

func DeleteDups() {

}

func getPlaylist() *spotify.FullPlaylist {
	client, clientError := Authenticator.GetClient()
	isClientNull(clientError)

	playlist, playlistErr := client.GetPlaylist(context.Background(), "6jQH6TbyhQrvb36ZBjLWQe")
	if playlistErr != nil {
		log.Fatal("Error grabbing playlist")
	}

	return playlist
}

func isClientNull(clientError error) {
	if clientError != nil {
		log.Fatal("Client is nil", clientError)
	}
}
