package main

import (
	"seda/synth"
)

func main() {
	const (
		sampleRate = 44100 // Standard CD-quality sample rate
		freq       = 440.0 // A4 note
		duration   = 5.0   // 2 seconds
	)

	waveform := synth.GenerateSquareWave(freq, sampleRate, duration)

	synth.PlayRawSound(waveform, sampleRate)

}
