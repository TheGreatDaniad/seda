package main

import (
	"seda/synth"
)

func main() {
	const (
		sampleRate = 44100 // Standard CD-quality sample rate
		freq       = 440.0 // A4 note
		duration   = 10.0  // 2 seconds
	)

	// waveform := synth.GenerateSineWave(freq, sampleRate, duration)
	waveform := synth.ComposeSound(440, synth.SoundCharacter{
		ADSR:      synth.ADSR{Attack: 0.1, Decay: 3, Sustain: 0.0, Release: 0.1},
		Harmonics: synth.SoundHarmonics{Next: [10]float32{0.8,0,7,0.6,0.5,0.4,0.3,0.2,0.1,0.05}, Previous: [10]float32{0.8,0,7,0.6,0.5,0.4,0.3,0.2,0.1,0.05}},
	}, sampleRate, duration)
	synth.PlayRawSound(waveform, sampleRate)

}
