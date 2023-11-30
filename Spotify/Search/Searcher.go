package Search

import (
	"context"
	"fmt"
	"github.com/ikediufomadu/Go_Spotify_Script/Spotify/Authenticator"
	"github.com/ikediufomadu/Go_Spotify_Script/Spotify/Playlist"
	"github.com/zmb3/spotify/v2"
	"log"
)

// Query is an interface representing a generic Spotify search query
type Query interface {
	QuerySpotify()
}

type SongNameQuery struct {
	songName   string
	artistName string
	searchType spotify.SearchType
}

type ArtistNameQuery struct {
	artistName string
	searchType spotify.SearchType
}

type GenreQuery struct {
	genre      string
	searchType spotify.SearchType
}

// NewSongNameQuery creates a new SongNameQuery instance
func NewSongNameQuery(songName string, searchType spotify.SearchType, artistName ...string) *SongNameQuery {
	var artist string
	if len(artistName) > 0 {
		artist = artistName[0]
	}

	return &SongNameQuery{
		songName:   songName,
		artistName: artist,
		searchType: searchType,
	}
}

// NewArtistNameQuery creates a new ArtistNameQuery instance
func NewArtistNameQuery(artistName string, searchType spotify.SearchType) *ArtistNameQuery {
	return &ArtistNameQuery{
		artistName: artistName,
		searchType: searchType,
	}
}

// NewGenreQuery creates a new GenreQuery instance
func NewGenreQuery(genre string, searchType spotify.SearchType) *GenreQuery {
	return &GenreQuery{
		genre:      genre,
		searchType: searchType,
	}
}

// QuerySpotify performs a Spotify search based on the SongNameQuery
func (snq SongNameQuery) QuerySpotify() {
	client, clientError := Authenticator.GetClient()
	if clientError != nil {
		log.Fatal("Client is nil", clientError)
	}

	// Formulate the search query
	query := fmt.Sprintf("%s artist %s", snq.songName, snq.artistName)

	// Perform the search
	search, searchError := client.Search(context.Background(), query, snq.searchType)
	if searchError != nil {
		log.Fatal("There was an error querying this song. ", searchError)
	}

	// If tracks are found, add the first track to the playlist
	if len(search.Tracks.Tracks) > 0 {
		firstTrack := search.Tracks.Tracks[0].ID
		Playlist.AddTrackToPlaylist(firstTrack)
	} else {
		log.Println("No tracks found for the given query.\n", query)
	}
}

// QuerySpotify performs a Spotify search based on the ArtistNameQuery
func (anq ArtistNameQuery) QuerySpotify() {
	client, err := Authenticator.GetClient()
	if err != nil {
		log.Fatal("Client is nil", err)
	}
	search, err := client.Search(context.Background(), anq.artistName, anq.searchType)
	if err != nil {
		log.Fatal("There was an error querying this artist. ", err)
	}
	fmt.Print(search.Artists)
}

// QuerySpotify performs a Spotify search based on the GenreQuery
func (gq GenreQuery) QuerySpotify() {
	client, err := Authenticator.GetClient()
	if err != nil {
		log.Fatal("Client is nil", err)
	}
	search, err := client.Search(context.Background(), gq.genre, gq.searchType)
	if err != nil {
		log.Fatal("There was an error querying this genre. ", err)
	}
	fmt.Print(search.Albums)
}
