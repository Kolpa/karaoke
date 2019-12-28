package commands

import "path/filepath"

type SongResponse struct {
	Command string
	Status  string
	Songs   []string
}

func ListSongs() SongResponse {
	matches, err := filepath.Glob("*.mp3")

	if err != nil {
		return SongResponse{
			Command: "list_songs",
			Status:  "error",
			Songs:   nil,
		}
	}

	for i, songFile := range matches {
		matches[i] = filepath.Base(songFile)
	}

	return SongResponse{
		Command: "list_songs",
		Status:  "ok",
		Songs:   matches,
	}
}
