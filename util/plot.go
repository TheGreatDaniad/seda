package util

import (
	"image/color"
	"math"

	"github.com/mjibson/go-dsp/fft"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func PlotFFT(sound []float64, sampleRate int, filename string) error {
	// Perform FFT
	fftResult := fft.FFTReal(sound)
	// Prepare data for plotting
	pts := make(plotter.XYs, len(fftResult)/2)
	for i := range pts {
		pts[i].X = float64(i) * float64(sampleRate) / float64(len(sound))
		amplitude := math.Sqrt(real(fftResult[i])*real(fftResult[i]) + imag(fftResult[i])*imag(fftResult[i]))
		pts[i].Y = amplitude
	}

	// Create a new plot
	p := plot.New()

	p.Title.Text = "FFT of Sound"
	p.X.Label.Text = "Frequency (Hz)"
	p.Y.Label.Text = "Amplitude"

	// Add the line plot to the plot
	line, err := plotter.NewLine(pts)
	if err != nil {
		return err
	}
	line.Color = color.RGBA{R: 255, A: 255}
	p.Add(line)

	// Save the plot to a PNG file
	if err := p.Save(6*vg.Inch, 4*vg.Inch, filename); err != nil {
		return err
	}

	return nil
}
