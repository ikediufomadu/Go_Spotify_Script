package Playlist

import (
	"context"
	"fmt"
	"github.com/ikediufomadu/Go_Spotify_Script/Spotify/Authenticator"
	"github.com/zmb3/spotify/v2"
	"log"
	"strings"
)

var userPlaylist *spotify.FullPlaylist

// createPlaylist creates a new Spotify playlist for the user
func CreatePlaylist(playlistName, playlistDescription string, isPublic, isCollaborative bool) {
	fmt.Println(playlistName, playlistDescription, isPublic, isCollaborative)
	client := Authenticator.GetClient()
	userID := Authenticator.GetUserID()

	_, playlistError := client.CreatePlaylistForUser(context.Background(), userID, playlistName, playlistDescription, isPublic, isCollaborative)
	if playlistError != nil {
		fmt.Println("There was a problem creating your playlist.", playlistError)
	}
}

// AddTrackToPlaylist adds a track to the authenticated user's Spotify playlist
func AddTrackToPlaylist(query string, track spotify.ID) {
	client := Authenticator.GetClient()

	if checkDuplicates(track, userPlaylist) {
		fmt.Printf("Found duplicate track '%s' in your playlist.\n", track.String())
	} else {
		_, playlistError := client.AddTracksToPlaylist(context.Background(), userPlaylist.ID, track)
		if playlistError != nil {
			fmt.Print("There was a problem adding a track to the playlist.")
		} else {
			fmt.Printf("%s added to playlist.\n", query)
		}
	}
}

func checkDuplicates(trackID spotify.ID, userPlaylist *spotify.FullPlaylist) bool {
	tracksInPlaylist := userPlaylist.Tracks.Tracks
	for _, track := range tracksInPlaylist {
		if track.Track.ID == trackID {
			return true
		}
	}
	return false
}

func GetPlaylist(playlistName string) {
	client := Authenticator.GetClient()
	userID := Authenticator.GetUserID()
	var playlistFound bool
	// Get all playlists for the user
	playlists, playlistsError := client.GetPlaylistsForUser(context.Background(), userID)
	if playlistsError != nil {
		log.Fatal("Could not retrieve user playlists.", playlistsError)
	}

	// Find the playlist by name
	for _, playlist := range playlists.Playlists {
		// Checks playlist names and also checks if it's owned by the user
		if strings.ToLower(playlist.Name) == strings.ToLower(playlistName) && playlist.Owner.ID == userID {
			var err error
			userPlaylist, err = client.GetPlaylist(context.Background(), playlist.ID)
			if err != nil {
				log.Fatalf("Could not access the playlist: %s", userPlaylist.Name)
			}
			playlistFound = true
		}
	}

	if !playlistFound {
		log.Fatal("Playlist not found")
	}
}
func SetPlaylist(playlistName string) {
	GetPlaylist(playlistName)
}
