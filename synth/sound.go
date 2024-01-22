package synth

import (
	"math"

	"github.com/hajimehoshi/oto"
)

// PlayRawSound plays raw sound data from a float64 slice.
func PlayRawSound(data []float64, sampleRate int) error {
	// Initialize audio context
	player, err := oto.NewContext(sampleRate, 2, 2, 4096) // 2 channels (stereo), 2 bytes per sample
	if err != nil {
		return err
	}
	defer player.Close()

	// Convert the float64 slice to a byte slice (PCM data)
	pcm := make([]byte, len(data)*2) // *2 because we're using 16-bit samples (2 bytes)
	for i, sample := range data {
		val := int16(sample * math.MaxInt16) // Convert float64 to int16
		pcm[2*i] = byte(val)                 // Lower 8 bits
		pcm[2*i+1] = byte(val >> 8)          // Higher 8 bits
	}

	// Play the sound
	p := player.NewPlayer()
	_, err = p.Write(pcm)
	if err != nil {
		return err
	}

	return nil
}
