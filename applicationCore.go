package main

import (
	"fmt"
	"github.com/ikediufomadu/Go_Spotify_Script/Spotify/Playlist"
	"strings"
)

var finished = false

func programCore() {
	for !finished {
		fmt.Println("Do you want to create a new playlist? (y/n)")
		var choice string
		var playlistName string
		fmt.Scanln(&choice)

		if strings.ToLower(choice) == "y" {
			var playlistDescription string
			var isPublic string
			var isCollaborative string

			fmt.Println("Enter playlist name")
			fmt.Scanln(&playlistName)

			fmt.Println("Enter playlist description")
			fmt.Scanln(&playlistDescription)

			fmt.Println("Do you want playlist to be public? (y/n)")
			fmt.Scanln(&isPublic)

			if strings.ToLower(isPublic) == "y" {
				fmt.Println("Do you want playlist to be collaborative (y/n)")
				fmt.Scanln(&isCollaborative)

				if strings.ToLower(isCollaborative) == "y" {
					finished = !finished
					Playlist.CreatePlaylist(playlistName, playlistDescription, true, true)
					Playlist.SetPlaylist(playlistName)
				} else if strings.ToLower(isCollaborative) == "n" {
					finished = !finished
					Playlist.CreatePlaylist(playlistName, playlistDescription, true, false)
					Playlist.SetPlaylist(playlistName)
				} else {
					fmt.Println("Please enter a valid option")
				}
			} else if strings.ToLower(isPublic) == "n" {
				finished = !finished
				Playlist.CreatePlaylist(playlistName, playlistDescription, false, false)
				Playlist.SetPlaylist(playlistName)
			}
		} else if strings.ToLower(choice) == "n" {
			fmt.Println("Enter the name of the playlist you want to add tracks to")
			fmt.Scanln(&playlistName)
			finished = !finished
			Playlist.GetPlaylist(playlistName)
		} else {
			fmt.Println("Please enter a valid option")
		}
	}
}
