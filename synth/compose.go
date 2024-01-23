package synth

import (
	"math"
)

// ... [SoundHarmonics, ADSR, SoundCharacter type definitions as before] ...

func ApplyADSR(sound []float64, adsr ADSR, sampleRate int) {
	totalSamples := len(sound)
	attackSamples := int(adsr.Attack * float64(sampleRate))
	decaySamples := int(adsr.Decay * float64(sampleRate))
	releaseSamples := int(adsr.Release * float64(sampleRate))

	for i := range sound {
		var amplitude float64

		if i < attackSamples {
			// Attack phase
			amplitude = float64(i) / float64(attackSamples)
		} else if i < attackSamples+decaySamples {
			// Decay phase
			decayProgress := float64(i-attackSamples) / float64(decaySamples)
			amplitude = (1.0-decayProgress)*(1.0-adsr.Sustain) + adsr.Sustain
		} else if i < totalSamples-releaseSamples {
			// Sustain phase
			amplitude = adsr.Sustain
		} else {
			// Release phase
			releaseProgress := float64(i-(totalSamples-releaseSamples)) / float64(releaseSamples)
			amplitude = (1.0 - releaseProgress) * adsr.Sustain
		}

		sound[i] *= amplitude
	}
}

func ComposeSound(fundamentalFreq float64, soundChar SoundCharacter, sampleRate int, duration float64) []float64 {
	numSamples := int(duration * float64(sampleRate))
	sound := make([]float64, numSamples)

	// Generate base sine wave for the fundamental frequency
	for i := range sound {
		sound[i] = math.Sin(2 * math.Pi * fundamentalFreq * float64(i) / float64(sampleRate))
	}

	// Add harmonics
	addHarmonics := func(harmonics [10]float32, harmonicType string) {
		for i, amp := range harmonics {
			harmonicFreq := fundamentalFreq * float64(i+1)
			if harmonicType == "previous" {
				harmonicFreq = fundamentalFreq / float64(i+1)
			}
			for j := range sound {
				sound[j] += float64(amp) * math.Sin(2*math.Pi*harmonicFreq*float64(j)/float64(sampleRate))
			}
		}
	}

	addHarmonics(soundChar.Harmonics.Next, "next")
	addHarmonics(soundChar.Harmonics.Previous, "previous")

	// Normalize the sound to prevent clipping
	maxAmplitude := 0.0
	for _, sample := range sound {
		if absSample := math.Abs(sample); absSample > maxAmplitude {
			maxAmplitude = absSample
		}
	}

	if maxAmplitude > 0 {
		for i := range sound {
			sound[i] /= maxAmplitude
		}
	}
	ApplyADSR(sound, soundChar.ADSR, sampleRate)
	return sound
}
