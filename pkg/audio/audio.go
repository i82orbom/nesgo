package audio

// Device represents an audio device
type Device interface {
	Close() error
}

// Source represents an audio source
type Source interface {
	AudioChannel() chan float32
}
