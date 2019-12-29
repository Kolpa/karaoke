package player

import "sync"

//PlaybackQueue a queue of songs to play
type PlaybackQueue struct {
    songs []string
    lock  sync.RWMutex
}

// New creates a new PlaybackQueue
func (s *PlaybackQueue) New() *PlaybackQueue {
    s.songs = make([]string, 0)
    return s
}

// Enqueue adds an Item to the end of the queue
func (s *PlaybackQueue) Enqueue(t string) {
    s.lock.Lock()
    s.songs = append(s.songs, t)
    s.lock.Unlock()
}

// Dequeue removes an Item from the start of the queue
func (s *PlaybackQueue) Dequeue() *string {
    s.lock.Lock()
    item := s.songs[0]
    s.songs = s.songs[1:len(s.songs)]
    s.lock.Unlock()
    return &item
}

// Front returns the item next in the queue, without removing it
func (s *PlaybackQueue) Front() *string {
    s.lock.RLock()
    item := s.songs[0]
    s.lock.RUnlock()
    return &item
}

// IsEmpty returns true if the queue is empty
func (s *PlaybackQueue) IsEmpty() bool {
    return len(s.songs) == 0
}

// Size returns the number of songs in the queue
func (s *PlaybackQueue) Size() int {
    return len(s.songs)
}