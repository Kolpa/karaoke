package commands

import "os"
import "path/filepath"

//ListSongsResponse Describes a api response listing songs
type ListSongsResponse struct {
	Command string
	Status  string
	Songs   []string
}

//ListSongs returns a SongResponse object containing all songs in current folder
func ListSongs() ListSongsResponse {
	var fileRoot = os.Getenv("SONG_ROOT")

	matches, err := filepath.Glob(fileRoot + "/*.mp3")

	if err != nil {
		return ListSongsResponse{
			Command: "list_songs",
			Status:  "error",
			Songs:   nil,
		}
	}

	return ListSongsResponse{
		Command: "list_songs",
		Status:  "ok",
		Songs:   matches,
	}
}
