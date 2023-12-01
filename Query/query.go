package Query

import (
	"fmt"
	"github.com/ikediufomadu/Go_Spotify_Script/Music"
	"github.com/ikediufomadu/Go_Spotify_Script/Spotify/Search"

	"github.com/zmb3/spotify/v2"
)

// WhichQuery performs Spotify queries based on user selection
func WhichQuery(selection string) {
	switch selection {
	case "rap":
		loopMusicGenre("rap", Music.Rap)
	case "hiphop":
		loopMusicGenre("hiphop", Music.Hiphop)
	case "rb":
		loopMusicGenre("rb", Music.RB)
	case "jazz":
		loopMusicGenre("jazz", Music.Jazz)
	case "rock":
		loopMusicGenre("rock", Music.Rock)
	case "pop":
		loopMusicGenre("pop", Music.Pop)
	case "electronic":
		loopMusicGenre("electronic", Music.Electronic)
	}
}

func loopMusicGenre(genreName string, genre []Music.Music) {
	for _, song := range genre {
		query, err := formQuery(song)
		if err != nil {
			formQueryError(song, genreName)
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

func formQueryError(song Music.Music, genre string) {
	fmt.Printf("Error forming a query for this song and artist %s %s. Current genre slice: %s", song.GetSongName(), song.Artist, genre)
}
