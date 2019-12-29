package commands

import "karaoke/internal/player"

//SkipSongResponse Describes a api response listing songs
type SkipSongResponse struct {
	Command string
	Status  string
}

//SkipSong skips the currently playing song
func SkipSong() SkipSongResponse {
	player.SkipSong()

	return SkipSongResponse{
		Command: "skip_song",
		Status:  "ok",
	}
}
