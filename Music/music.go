package Music

// Music represents a structure for storing information about a piece of music
type Music struct {
	SongName string
	Artist   string
}

func (m Music) GetSongName() string {
	return m.SongName
}

func (m Music) GetArtist() string {
	return m.Artist
}
