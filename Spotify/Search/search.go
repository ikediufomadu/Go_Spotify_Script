package Search

import (
	"context"
	"fmt"
	"github.com/ikediufomadu/Go_Spotify_Script/Music"
	"github.com/ikediufomadu/Go_Spotify_Script/Spotify/Authenticator"
	"github.com/ikediufomadu/Go_Spotify_Script/Spotify/Playlist"
	"github.com/zmb3/spotify/v2"
	"log"
)

// Query is an interface representing a generic Spotify search query
type Query interface {
	QuerySpotify()
}

type TrackQuery struct {
	songName   string
	artistName string
	searchType spotify.SearchType
}

// NewTrackQuery creates a new TrackQuery instance
func NewTrackQuery(music Music.Music, searchType spotify.SearchType) (*TrackQuery, error) {
	return &TrackQuery{
		songName:   music.SongName,
		artistName: music.Artist,
		searchType: searchType,
	}, nil
}

// QuerySpotify performs a Spotify search based on the TrackQuery
func (tq TrackQuery) QuerySpotify() error {
	client, clientError := Authenticator.GetClient()
	if clientError != nil {
		log.Fatal("Client is nil", clientError)
	}

	// Formulate the search query
	query := fmt.Sprintf("%s artist %s", tq.songName, tq.artistName)

	// Perform the search
	search, searchError := client.Search(context.Background(), query, tq.searchType)
	if searchError != nil {
		log.Println("There is no such track.", query, searchError)
		return searchError
	}

	// If tracks are found, add the first track to the playlist
	tracks := search.Tracks.Tracks
	if len(tracks) > 0 {
		firstTrack := tracks[0].ID
		// Change the playlist name to your own
		// Make sure the playlist's name consists of regular unicode characters.
		Playlist.AddTrackToPlaylist(query, firstTrack, "test")
	}

	return nil
}
