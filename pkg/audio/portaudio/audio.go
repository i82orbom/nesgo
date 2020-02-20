package portaudio

import(
	"fmt"
	"github.com/i82orbom/nesgo/pkg/audio"
	"github.com/gordonklaus/portaudio"
)

type audioDevice struct {
	*portaudio.Stream
	sampleChannel chan float32
}

// NewDevice creates a new portaudio device
func NewDevice(audioSource audio.Source) (audio.Device, error) {
	err := portaudio.Initialize()
	if err != nil {
		return nil, err
	}
	api, err := portaudio.DefaultHostApi()
	if err != nil {
		return nil, err
	}
	sampleChannel := audioSource.AudioChannel()
	params := portaudio.HighLatencyParameters(nil, api.DefaultOutputDevice)
	stream, err := portaudio.OpenStream(params, createCallback(sampleChannel))
	if err != nil {
		return nil, err
	}
	return &audioDevice{
		Stream: stream,
		sampleChannel: sampleChannel,
	}, nil
}

func (d *audioDevice) Close() error {
	return d.Stream.Close()
}

func createCallback(inputCh chan float32) func([]float32) {
	return func (output []float32)  {
		for i := range output {
			fmt.Printf("FFFF")
			select {
			case val := <- inputCh:
				output[i] = val
			default:
				output[i] = 0
			}
		}
	}
}