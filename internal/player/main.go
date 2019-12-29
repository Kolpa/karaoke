package player

import "log"
import vlc "github.com/adrg/libvlc-go"

var listPlayer *vlc.Player
var playbackQueue = PlaybackQueue{}

//Initialize sets up the vlc player runtime
func Initialize() {
	playbackQueue.New()

	err := vlc.Init()
	if err != nil {
		log.Fatal(err)
	}

	player, err := vlc.NewPlayer()
	if err != nil {
		log.Fatal(err)
	}

	listPlayer = player

	manager, err := player.EventManager()
	if err != nil {
		log.Fatal(err)
	}

	_, err = manager.Attach(vlc.MediaPlayerEndReached, endOfSongCallback, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func endOfSongCallback(event vlc.Event, userData interface{}) {
	fetchAndPlay()
}

func fetchAndPlay() {
	song := playbackQueue.Dequeue()
	if song == nil {
		return
	}

	listPlayer.LoadMediaFromPath(*song)
	listPlayer.Play()
}

//AddSongToQueue adds a song to the playback buffer
func AddSongToQueue(path string) {
	playbackQueue.Enqueue(path)

	if !listPlayer.IsPlaying() {
		fetchAndPlay()
	}
}

//SkipSong forces playback of next Song
func SkipSong() {
	fetchAndPlay()
}

//GetCurrentQueue returns the current queue
func GetCurrentQueue() []string {
	return playbackQueue.songs
}
