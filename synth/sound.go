package synth

import (
	"math"

	"github.com/hajimehoshi/oto"
)

// sound harmonics amplitude of the previous and next 10 harmonics of a sound, value between 0-1
type SoundHarmonics struct {
	Next     [50]float64
	Previous [50]float64
}

// ADSR represents the Attack, Decay, Sustain, and Release envelope parameters for a synthesizer.
type ADSR struct {
	Attack  float64 // Time (in seconds) for the sound to reach peak amplitude
	Decay   float64 // Time (in seconds) for the sound to decay to the sustain level
	Sustain float64 // Sustain level (a value between 0 and 1 representing the steady state amplitude)
	Release float64 // Time (in seconds) for the sound to decay to zero amplitude after the key is released
}

type SoundCharacter struct {
	Harmonics SoundHarmonics
	ADSR      ADSR
}

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
