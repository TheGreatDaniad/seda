// Package synth provides functions to generate basic waveforms used in synthesizers.
package synth

import (
    "math"
)

// GenerateSineWave creates a sine wave of a specified frequency, sample rate, and duration.
// The sine wave is a smooth periodic oscillation and is fundamental in sound synthesis.
// freq: Frequency of the sine wave in Hertz.
// sampleRate: Sampling rate in samples per second.
// duration: Duration of the wave in seconds.
// Returns a slice of float64 representing the waveform.
func GenerateSineWave(freq float64, sampleRate int, duration float64) []float64 {
    numSamples := int(duration * float64(sampleRate))
    wave := make([]float64, numSamples)
    for i := range wave {
        wave[i] = math.Sin(2.0 * math.Pi * freq * float64(i) / float64(sampleRate))
    }
    return wave
}

// GenerateSquareWave creates a square wave of a specified frequency, sample rate, and duration.
// Square waves alternate between a maximum and a minimum value, creating a distinct sound.
// freq: Frequency of the square wave in Hertz.
// sampleRate: Sampling rate in samples per second.
// duration: Duration of the wave in seconds.
// Returns a slice of float64 representing the waveform.
func GenerateSquareWave(freq float64, sampleRate int, duration float64) []float64 {
    numSamples := int(duration * float64(sampleRate))
    wave := make([]float64, numSamples)
    for i := range wave {
        t := float64(i) / float64(sampleRate)
        if math.Mod(t, 1.0/freq) < 1.0/(2*freq) {
            wave[i] = 1.0
        } else {
            wave[i] = -1.0
        }
    }
    return wave
}

// GenerateSawtoothWave creates a sawtooth wave of a specified frequency, sample rate, and duration.
// Sawtooth waves have a characteristic "buzzing" sound and contain all harmonic components.
// freq: Frequency of the sawtooth wave in Hertz.
// sampleRate: Sampling rate in samples per second.
// duration: Duration of the wave in seconds.
// Returns a slice of float64 representing the waveform.
func GenerateSawtoothWave(freq float64, sampleRate int, duration float64) []float64 {
    numSamples := int(duration * float64(sampleRate))
    wave := make([]float64, numSamples)
    for i := range wave {
        wave[i] = 2.0 * (float64(i)/float64(sampleRate)*freq - math.Floor(0.5+float64(i)/float64(sampleRate)*freq))
    }
    return wave
}

// GenerateTriangleWave creates a triangle wave of a specified frequency, sample rate, and duration.
// Triangle waves are non-sinusoidal and named for their triangular shape.
// freq: Frequency of the triangle wave in Hertz.
// sampleRate: Sampling rate in samples per second.
// duration: Duration of the wave in seconds.
// Returns a slice of float64 representing the waveform.
func GenerateTriangleWave(freq float64, sampleRate int, duration float64) []float64 {
    numSamples := int(duration * float64(sampleRate))
    wave := make([]float64, numSamples)
    for i := range wave {
        wave[i] = math.Abs(2.0*(float64(i)/float64(sampleRate)*freq-math.Floor(0.5+float64(i)/float64(sampleRate)*freq))) - 1.0
    }
    return wave
}
