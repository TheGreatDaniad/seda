package main

import (
	"fmt"
	"seda/synth"
	"seda/util"
	"time"
)

func main() {
	// const (
	// 	sampleRate = 44100 // Standard CD-quality sample rate
	// 	freq       = 440.0 // A4 note
	// 	duration   = 10.0  // 2 seconds
	// )

	sound, sampleRate, err := util.ReadWavFile("violin.wav")
	if err != nil {
		panic(err)
	}

	// util.PlotFFT(sound, sampleRate, "violin.png")
	_, harmonics, err := util.AnalyzeSound(sound, sampleRate)
	if err != nil {
		panic(err)
	}
	t1 := time.Now()
	waveform := synth.ComposeSound(440, synth.SoundCharacter{
		ADSR:      synth.ADSR{Attack: 2, Decay: 8, Sustain: 0.0, Release: 0.1},
		Harmonics: synth.SoundHarmonics{Next: harmonics},
	}, sampleRate, 1)
	t2 := time.Since(t1)
    fmt.Println(t2)
	// util.PlotFFT(waveform, sampleRate, "violin-a4-2.png")

	util.SaveSoundToWav("violin-a4-synth.wav", waveform, sampleRate)
	// synth.PlayRawSound(waveform, sampleRate)
}
