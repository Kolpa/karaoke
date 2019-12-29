package commands

import "karaoke/internal/player"

//GetQueueResponse Describes a api response listing songs
type GetQueueResponse struct {
	Command string
	Status  string
	Songs   []string
}

//GetQueue returns a GetQueueResponse object containing all songs in the current queue
func GetQueue() GetQueueResponse {
	queue := player.GetCurrentQueue()

	if queue == nil {
		return GetQueueResponse{
			Command: "get_queue",
			Status:  "error",
			Songs:   nil,
		}
	}

	return GetQueueResponse{
		Command: "get_queue",
		Status:  "ok",
		Songs:   queue,
	}
}
