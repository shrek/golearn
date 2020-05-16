package bandit

import (
	//"fmt"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func getPoints(rewards []float64) plotter.XYs {
	pts := make(plotter.XYs, len(rewards))
	ind := 1
	for i := range pts {
		pts[i].X = float64(ind)
		pts[i].Y = rewards[ind-1]
		ind++
	}
	return pts
}

func plotSeries(fname string, names []string, series ...[]float64) error {
	p, err := plot.New()
	if err != nil {
		return err
	}
	p.Title.Text = "Functions"
	p.X.Label.Text = "iteration"
	p.Y.Label.Text = "average reward"

	//for i, n := range names {
	err = plotutil.AddLinePoints(p,
		names[0], getPoints(series[0]),
		names[1], getPoints(series[1]),
		names[2], getPoints(series[2]),
		names[3], getPoints(series[3]),
		names[4], getPoints(series[4]),
	)

	if err != nil {
		panic(err)
	}
	//}

	// Save the plot to a PNG file.
	if err := p.Save(8*vg.Inch, 4*vg.Inch, fname); err != nil {
		panic(err)
	}
	return err
}
