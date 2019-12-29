package commands

import "path/filepath"

//ListSongsResponse Describes a api response listing songs
type ListSongsResponse struct {
	Command string
	Status  string
	Songs   []string
}

//ListSongs returns a SongResponse object containing all songs in current folder
func ListSongs() ListSongsResponse {
	matches, err := filepath.Glob("*.mp3")

	if err != nil {
		return ListSongsResponse{
			Command: "list_songs",
			Status:  "error",
			Songs:   nil,
		}
	}

	for i, songFile := range matches {
		matches[i] = filepath.Base(songFile)
	}

	return ListSongsResponse{
		Command: "list_songs",
		Status:  "ok",
		Songs:   matches,
	}
}
