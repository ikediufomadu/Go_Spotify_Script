package Music

type Music struct {
	SongName   string `json:"song_name"`
	Delimiter  string
	Artist     string `json:"artist"`
	Delimiter2 string
	Genre      string `json:"genre"`
	Delimiter3 string
	Published  string `json:"published"`
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
