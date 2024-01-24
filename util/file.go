package util

import (
	"os"

	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
)

// ReadWavFile reads a WAV file and returns its data as []float64 and its sample rate.
func ReadWavFile(filename string) ([]float64, int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, 0, err
	}
	defer file.Close()

	decoder := wav.NewDecoder(file)
	buf, err := decoder.FullPCMBuffer()
	if err != nil {
		return nil, 0, err
	}

	data := make([]float64, len(buf.Data))
	for i, sample := range buf.Data {
		data[i] = float64(sample) / float64((int(1)<<uint(buf.SourceBitDepth))/2)
	}

	return data, buf.Format.SampleRate, nil
}

// SaveSoundToWav saves a sound (as a slice of float64) to a WAV file.
func SaveSoundToWav(filename string, sound []float64, sampleRate int) error {
	outFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer outFile.Close()

	enc := wav.NewEncoder(outFile, sampleRate, 16, 1, 1) // 16-bit, 1 channel (mono)

	intBuffer := &audio.IntBuffer{Data: make([]int, len(sound)), Format: &audio.Format{SampleRate: sampleRate, NumChannels: 1}}
	maxVal := float64((1 << 15) - 1)
	for i, sample := range sound {
		intBuffer.Data[i] = int(sample * maxVal)
	}

	if err := enc.Write(intBuffer); err != nil {
		return err
	}

	// Close the encoder which will finalize the WAV file
	return enc.Close()
}
