package commands

import "karaoke/internal/player"

//GetCurrentSongResponse Describes a api response listing songs
type GetCurrentSongResponse struct {
	Command string
	Status  string
	Song    string
}

//GetCurrentSong returns a GetCurrentSongResponse object containing all songs in the current queue
func GetCurrentSong() GetCurrentSongResponse {
	song := player.GetCurrentSong()

	return GetCurrentSongResponse{
		Command: "current_song",
		Status:  "ok",
		Song:   song,
	}
}
