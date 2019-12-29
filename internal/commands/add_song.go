package commands

import "karaoke/internal/player"
import "log"

//AddSongResponse Describes a api response listing songs
type AddSongResponse struct {
	Command string
	Status  string
}

//AddSong adds a song to the playlist
func AddSong(arguments interface{}) AddSongResponse {
	songName, ok := arguments.(string)

	log.Printf("%T", arguments)

	if !ok {
		return AddSongResponse{
			Command: "add_song",
			Status: "error",
		}
	}

	player.AddSongToQueue(songName)

	return AddSongResponse{
		Command: "add_song",
		Status:  "ok",
	}
}
