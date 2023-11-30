package Playlist

import (
	"context"
	"fmt"
	"github.com/ikediufomadu/Go_Spotify_Script/Spotify/Authenticator"
	"github.com/zmb3/spotify/v2"
	"log"
)

func CreatePlaylist(playlistName, playlistDescription string, public, collaborative bool) {
	client, err := Authenticator.GetClient()
	if err != nil {
		log.Fatal("Client is nil", err)
	}
	userID, err2 := Authenticator.GetUserID()
	if err2 != nil {
		log.Fatal("User is nil", err)
	}
	playlist, err := client.CreatePlaylistForUser(context.Background(), userID, playlistName, playlistDescription, public, collaborative)
	if err != nil {
		fmt.Println("There was a problem creating your playlist.", err)
	} else {
		fmt.Println(playlist.ID)
	}
}

func AddTrackToPlaylist(track spotify.ID) {
	client, err := Authenticator.GetClient()
	if err != nil {
		log.Fatal("Client is nil", err)
	}
	fullPlaylist := getPlaylist()
	snapshotID, err := client.AddTracksToPlaylist(context.Background(), fullPlaylist.ID, track)
	if err != nil {
		fmt.Printf("There was a problem adding this song to the playlist %s: %v\n", fullPlaylist.Name, err)
	} else {
		fmt.Printf("Added %s to playlist. %v", track.String(), snapshotID)
	}
}

func getPlaylist() *spotify.FullPlaylist {
	client, err := Authenticator.GetClient()
	if err != nil {
		log.Fatal("Client is nil", err)
	}

	playlist, playlistErr := client.GetPlaylist(context.Background(), "6jQH6TbyhQrvb36ZBjLWQe")
	if playlistErr != nil {
		fmt.Println("Error grabbing playlist")
	}

	return playlist
}
