package util

import (
	"math"

	"github.com/mjibson/go-dsp/fft"
)


func AnalyzeSound(sound []float64, sampleRate int) (float64, [50]float64, error) {
    var harmonics [50]float64

    // Perform FFT
    fftResult := fft.FFTReal(sound)

    // Find the peak amplitude in a given frequency range
    findPeakAmplitude := func(startIndex, endIndex int) float64 {
        peakAmplitude := 0.0
        for i := startIndex; i <= endIndex && i < len(fftResult)/2; i++ {
            amplitude := math.Sqrt(real(fftResult[i])*real(fftResult[i]) + imag(fftResult[i])*imag(fftResult[i]))
            if amplitude > peakAmplitude {
                peakAmplitude = amplitude
            }
        }
        return peakAmplitude
    }

    // Estimate the fundamental frequency index
    fundamentalIndex := 0
    fundamentalAmplitude := findPeakAmplitude(0, len(fftResult)/20)

    // Find the actual fundamental frequency index
    for i, complexValue := range fftResult {
        amplitude := math.Sqrt(real(complexValue)*real(complexValue) + imag(complexValue)*imag(complexValue))
        if amplitude == fundamentalAmplitude {
            fundamentalIndex = i
            break
        }
    }

    // Calculate the fundamental frequency
    fundamentalFreq := float64(fundamentalIndex) * float64(sampleRate) / float64(len(sound))

    // Calculate the amplitudes of the next 10 harmonics
    for i := 1; i <= 50; i++ { // Start from the second harmonic
        startIndex := fundamentalIndex*i - 5
        endIndex := fundamentalIndex*i + 5
        harmonicAmplitude := findPeakAmplitude(startIndex, endIndex)
        harmonics[i-1] = harmonicAmplitude / fundamentalAmplitude // Normalize by fundamental amplitude
    }

    return fundamentalFreq, harmonics, nil
}
