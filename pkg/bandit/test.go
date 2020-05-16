// test distributions and plotting
package bandit

import (
	"image/color"
	//"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"

	"gonum.org/v1/gonum/stat/distuv"
)

func PlotNormal() {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = "Functions"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	// A quadratic function x^2
	quad := plotter.NewFunction(func(x float64) float64 { return x * x })
	quad.Color = color.RGBA{B: 255, A: 255}

	normal := distuv.Normal{
		Mu:    0,
		Sigma: 0.5,
	}
	norm := plotter.NewFunction(func(x float64) float64 {
		return normal.Prob(x)
	})
	norm.Color = color.RGBA{G: 255, A: 255}
	normSample := plotter.NewFunction(func(x float64) float64 {
		return normal.Rand()
	})
	normSample.Color = color.RGBA{R: 255, A: 255}

	// Add the functions and their legend entries.
	p.Add(quad, norm, normSample)
	p.Legend.Add("x^2", quad)
	p.Legend.Add("norm(x)", norm)
	p.Legend.Add("normSample", normSample)
	p.Legend.ThumbnailWidth = 0.5 * vg.Inch

	// Set the axis ranges.  Unlike other data sets,
	// functions don't set the axis ranges automatically
	// since functions don't necessarily have a
	// finite range of x and y values.
	p.X.Min = -5
	p.X.Max = 5
	p.Y.Min = 0
	p.Y.Max = 1

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "/tmp/normal.png"); err != nil {
		panic(err)
	}
}
