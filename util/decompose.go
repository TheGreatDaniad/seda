package util

import (
	"fmt"
	"math"
	"sort"

	"github.com/mjibson/go-dsp/fft"
)

func AnalyzeSound(sound []float64, sampleRate int) (float64, [50]float64, []struct {
	FrequencyRatio float64
	MagnitudeRatio float64
}, error) {
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

	topFreqs, err := TopFrequencies(fundamentalFreq, fftResult, sampleRate)
	if err != nil {
		return 0, [50]float64{}, nil, err
	}

	return fundamentalFreq, harmonics, topFreqs, nil
}

type FrequencyData struct {
	FrequencyRatio float64 // Frequency / Fundamental Frequency
	MagnitudeRatio float64 // Magnitude / Fundamental Magnitude
}
type FrequencyMagnitude struct {
	Frequency float64
	Magnitude float64
}

func TopFrequencies(fundamentalFreq float64, fftResult []complex128, sampleRate int) ([]struct {
	FrequencyRatio float64
	MagnitudeRatio float64
}, error) {
	freqMagnitudes := make([]FrequencyMagnitude, 0)
	fundamentalMagnitude := 0.0

	// Add fundamental frequency first
	fundamentalIndex := int(fundamentalFreq * float64(len(fftResult)) / float64(sampleRate))
	if fundamentalIndex < len(fftResult) {
		fundamentalMagnitude = math.Sqrt(real(fftResult[fundamentalIndex])*real(fftResult[fundamentalIndex]) + imag(fftResult[fundamentalIndex])*imag(fftResult[fundamentalIndex]))
		freqMagnitudes = append(freqMagnitudes, FrequencyMagnitude{Frequency: fundamentalFreq, Magnitude: fundamentalMagnitude})
	}

	// Add other frequencies within the specified range
	for i := range fftResult {
		frequency := float64(i) * float64(sampleRate) / float64(len(fftResult))

		if frequency >= 20 && frequency <= 20000 && frequency != fundamentalFreq {

			magnitude := math.Sqrt(real(fftResult[i])*real(fftResult[i]) + imag(fftResult[i])*imag(fftResult[i]))
			freqMagnitudes = append(freqMagnitudes, FrequencyMagnitude{Frequency: frequency, Magnitude: magnitude})
		}
	}

	if fundamentalMagnitude == 0 {
		return nil, fmt.Errorf("fundamental frequency magnitude is zero")
	}

	// Sort by magnitude
	sort.Slice(freqMagnitudes, func(i, j int) bool {
		return freqMagnitudes[i].Magnitude > freqMagnitudes[j].Magnitude
	})

	numTopFreqs := len(freqMagnitudes)
	if numTopFreqs > 2000 {
		numTopFreqs = 2000
	}

	// Select top frequencies and calculate ratios
	topFrequencies := make([]struct {
		FrequencyRatio float64
		MagnitudeRatio float64
	}, numTopFreqs)

	for i := 0; i < numTopFreqs; i++ {
		fm := freqMagnitudes[i]
		topFrequencies[i] = struct {
			FrequencyRatio float64
			MagnitudeRatio float64
		}{
			FrequencyRatio: fm.Frequency / fundamentalFreq,
			MagnitudeRatio: fm.Magnitude / fundamentalMagnitude,
		}
	}

	return topFrequencies, nil
}
func isDivisible(num1, num2 float64) bool {
	const epsilon = 1e-9 // Small value to account for floating-point precision issues
	remainder := math.Mod(num1, num2)
	return math.Abs(remainder) < epsilon
}
