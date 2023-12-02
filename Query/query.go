package Query

import (
	"fmt"
	"github.com/ikediufomadu/Go_Spotify_Script/Music"
	"github.com/ikediufomadu/Go_Spotify_Script/Spotify/Search"

	"github.com/zmb3/spotify/v2"
)

func ScrapedMusic(musicSlice []Music.Music) {
	for _, song := range musicSlice {
		query, err := formQuery(song)
		if err != nil {
			formQueryError(song)
			continue
		}

		noTrackError := query.QuerySpotify()
		if noTrackError != nil {
			continue
		}
	}
}

func formQuery(song Music.Music) (*Search.TrackQuery, error) {
	return Search.NewTrackQuery(song, spotify.SearchType(spotify.SearchTypeTrack))
}

func formQueryError(song Music.Music) {
	fmt.Printf("Error forming a query for this song and artist %s %s.", song.GetSongName(), song.GetArtist())
}
