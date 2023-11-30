package Search

import (
	"context"
	"fmt"
	"github.com/ikediufomadu/Go_Spotify_Script/Spotify/Authenticator"
	"github.com/ikediufomadu/Go_Spotify_Script/Spotify/Playlist"
	"github.com/zmb3/spotify/v2"
	"log"
)

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

func NewSongNameQuery(songName string, artistName string, searchType spotify.SearchType) *SongNameQuery {
	return &SongNameQuery{
		songName:   songName,
		artistName: artistName,
		searchType: searchType,
	}
}

func NewArtistNameQuery(artistName string, searchType spotify.SearchType) *ArtistNameQuery {
	return &ArtistNameQuery{
		artistName: artistName,
		searchType: searchType,
	}
}

func NewGenreQuery(genre string, searchType spotify.SearchType) *GenreQuery {
	return &GenreQuery{
		genre:      genre,
		searchType: searchType,
	}
}

func (snq SongNameQuery) QuerySpotify() {
	client, clientError := Authenticator.GetClient()
	if clientError != nil {
		log.Fatal("Client is nil", clientError)
	}

	query := fmt.Sprintf("%s artist %s", snq.songName, snq.artistName)

	search, searchError := client.Search(context.Background(), query, snq.searchType)
	if searchError != nil {
		log.Fatal("There was an error querying this song. ", searchError)
	}

	if len(search.Tracks.Tracks) > 0 {
		firstTrack := search.Tracks.Tracks[0].ID
		Playlist.AddTrackToPlaylist(firstTrack)
	} else {
		log.Println("No tracks found for the given query.\n", query)
	}
}

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
