package nes

// APU represents the NES Audio Processing Unit
type APU struct {
	audioChannel chan float32
}

// NewAPU creates a new NES APU
func NewAPU() *APU {
	return &APU{
		audioChannel: make(chan float32, 44100),
	}
}

// AudioChannel returns an audio channel to output the emulated audio
func (a *APU) AudioChannel() chan float32 {
	return a.audioChannel
}

// CPUWrite writes data to the APU to update its status
func (a *APU) CPUWrite(address uint16, data uint8) {

}
