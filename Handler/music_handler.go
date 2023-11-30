package Handler

import (
	"github.com/ikediufomadu/Go_Spotify_Script/Music"
	"github.com/ikediufomadu/Go_Spotify_Script/Spotify/Search"
	"github.com/zmb3/spotify/v2"
	"time"
)

var (
	songName   []string
	songArtist []string
	songGenre  []string
)

// Appends music data to the global slices
func storeMusicData(musicData []Music.Music) {
	for _, music := range musicData {
		songName = append(songName, music.GetSongName())
		songArtist = append(songArtist, music.GetArtist())
		songGenre = append(songGenre, music.GetGenre())
	}
}

// WhichQuery performs Spotify queries based on user selection
func WhichQuery(selection int) {
	switch selection {
	case 1:
		// Search for each song by name and artist
		for i, songName := range GetSongNameSlice() {
			artistName := songArtist[i]
			nameQuery := Search.NewSongNameQuery(songName, artistName, spotify.SearchType(spotify.SearchTypeTrack))
			nameQuery.QuerySpotify()

			time.Sleep(3 * time.Second)
		}
	case 2:
		// Search for each artist
		for _, artistName := range GetSongArtistSlice() {
			artistQuery := Search.NewArtistNameQuery(artistName, spotify.SearchType(spotify.SearchTypeArtist))
			artistQuery.QuerySpotify()
		}
	case 3:
		// Search for each genre
		for _, genreType := range GetSongGenreSlice() {
			genreQuery := Search.NewGenreQuery(genreType, spotify.SearchType(spotify.SearchTypePlaylist))
			genreQuery.QuerySpotify()
		}
	}
}

func GetSongNameSlice() []string {
	return songName
}

func GetSongArtistSlice() []string {
	return songArtist
}

func GetSongGenreSlice() []string {
	return songGenre
}
