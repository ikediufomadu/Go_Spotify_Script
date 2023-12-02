package Playlist

import (
	"context"
	"fmt"
	"github.com/ikediufomadu/Go_Spotify_Script/Spotify/Authenticator"
	"github.com/zmb3/spotify/v2"
	"log"
	"strings"
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
func AddTrackToPlaylist(query string, track spotify.ID, playlistName string) {
	client, clientError := Authenticator.GetClient()
	isClientNull(clientError)

	userPlaylist, getPlaylistError := getPlaylist(playlistName)
	if getPlaylistError != nil || userPlaylist == nil {
		log.Fatal("Playlist does not exist.")
	}

	_, playlistError := client.AddTracksToPlaylist(context.Background(), userPlaylist.ID, track)
	if playlistError != nil {
		fmt.Printf("There was a problem adding this track to the playlist. %s.\n %s: %s\n", query, userPlaylist.Name, playlistError)
	} else {
		fmt.Printf("%s added to playlist.\n", query)
	}
}

func DeleteDuplicateTracks() {

}

func DeleteTrack() {

}

func getPlaylist(playlistName string) (*spotify.FullPlaylist, error) {
	client, clientError := Authenticator.GetClient()
	isClientNull(clientError)

	userID, userIDError := Authenticator.GetUserID()
	if userIDError != nil {
		return nil, userIDError
	}

	// Get all playlists for the user
	playlists, playlistsError := client.GetPlaylistsForUser(context.Background(), userID)
	if playlistsError != nil {
		return nil, playlistsError
	}

	// Find the playlist by name
	for _, playlist := range playlists.Playlists {

		if strings.ToLower(playlist.Name) == strings.ToLower(playlistName) && playlist.Owner.ID == userID {
			// Found the playlist by name
			return client.GetPlaylist(context.Background(), playlist.ID)
		}
	}

	// Playlist not found
	return nil, playlistsError
}

// isClientNull checks if the Spotify client is nil and logs a fatal error if so
func isClientNull(clientError error) {
	if clientError != nil {
		log.Fatal("Client is nil", clientError)
	}
}
