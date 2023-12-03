package main

import (
	"bufio"
	"fmt"
	"github.com/ikediufomadu/Go_Spotify_Script/Spotify/Playlist"
	"os"
	"strings"
)

var finished = false

func programCore() {
	for !finished {
		fmt.Println("Do you want to create a new playlist? (y/n)")
		var choice string
		var playlistName string
		var err error
		fmt.Scanln(&choice)
		reader := bufio.NewReader(os.Stdin)

		if strings.ToLower(choice) == "y" {
			fmt.Println("Enter playlist name")
			playlistName, err = reader.ReadString('\n')

			if err != nil {
				fmt.Println("Error reading input:", err)
				return
			}
			playlistName = playlistName[:len(playlistName)-1]

			finished = !finished
			Playlist.CreatePlaylist(playlistName)
			Playlist.SetPlaylist(playlistName)

		} else if strings.ToLower(choice) == "n" {
			fmt.Println("Enter the name of the playlist you want to add tracks to")
			playlistName, err = reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading input:", err)
				return
			}
			playlistName = playlistName[:len(playlistName)-1]
			finished = !finished
			Playlist.GetPlaylist(playlistName)
		} else {
			fmt.Println("Please enter a valid option")
		}
	}
}
