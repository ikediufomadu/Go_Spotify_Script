package Music

import "strings"

// Slices for each music genre
var (
	Rap        []Music
	Hiphop     []Music
	RB         []Music
	Jazz       []Music
	Rock       []Music
	Pop        []Music
	Electronic []Music
)

// Music represents a structure for storing information about a piece of music
type Music struct {
	SongName   string `json:"song_name"`
	Delimiter  string
	Artist     string `json:"artist"`
	Delimiter2 string
	Genre      string `json:"genre"`
	Delimiter3 string
	Published  string `json:"published"`
}

// StoreMusicData Appends music data to the global slices based on genre
func StoreMusicData(musicData []Music) {
	for _, music := range musicData {
		if strings.Contains(strings.ToLower(music.Genre), strings.ToLower("rap")) {
			Rap = append(Rap, music)
		} else if strings.Contains(strings.ToLower(music.Genre), strings.ToLower("hiphop")) {
			Hiphop = append(Hiphop, music)
		} else if strings.Contains(strings.ToLower(music.Genre), strings.ToLower("rb")) {
			RB = append(RB, music)
		} else if strings.Contains(strings.ToLower(music.Genre), strings.ToLower("jazz")) {
			Jazz = append(Jazz, music)
		} else if strings.Contains(strings.ToLower(music.Genre), strings.ToLower("rock")) {
			Rock = append(Rock, music)
		} else if strings.Contains(strings.ToLower(music.Genre), strings.ToLower("pop")) {
			Pop = append(Pop, music)
		} else if strings.Contains(strings.ToLower(music.Genre), strings.ToLower("electronic")) {
			Electronic = append(Electronic, music)
		}
	}
}

func (m Music) GetSongName() string {
	return m.SongName
}

func (m Music) GetArtist() string {
	return m.Artist
}

func (m Music) GetGenre() string {
	return m.Genre
}

func (m Music) GetDatePublished() string {
	return m.Published
}
