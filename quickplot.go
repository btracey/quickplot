package quickplot

import (
	"errors"
	"sort"

	"code.google.com/p/plotinum/plotter"
	"github.com/btracey/myplot"
	"github.com/gonum/floats"
	"github.com/gonum/matrix/mat64"
)

// ScatterFromColumns takes two columns of data and returns a scatter of them
// to add to a plot
func ScatterFromColumns(data *mat64.Dense, x, y int) (*plotter.Scatter, error) {
	r, c := data.Dims()
	if x >= c {
		return nil, errors.New("x variable greater than number of columns")
	}
	if y >= c {
		return nil, errors.New("x variable greater than number of columns")
	}
	// Construct the xys with the two columns
	xys := make(plotter.XYs, r)
	for i := range xys {
		xys[i].X = data.At(i, x)
		xys[i].Y = data.At(i, y)
	}
	return plotter.NewScatter(xys)
}

// CDF returns a plotter for a CDF of points. Points must be sorted. Clips the data
// at min and max cdf in case there are extreme outliers
func CDF(data, weights []float64, min, max float64) *plotter.Line {
	if !sort.Float64sAreSorted(data) {
		panic("quickplot: data not sorted")
	}

	if weights == nil {
		weights = make([]float64, len(data))
		for i := range weights {
			weights[i] = 1
		}
	}
	if len(weights) != len(data) {
		panic("quickplot: slice length mismatch")
	}

	cumsum := make([]float64, len(data))

	floats.CumSum(cumsum, weights)
	floats.Scale(cumsum[len(cumsum)-1], cumsum)

	pts := make(plotter.XYs, len(data))
	for i := range pts {
		pts[i].X = data[i]
		pts[i].Y = cumsum[i]
	}
	line, err := plotter.NewLine(pts)
	if err != nil {
		panic(err)
	}
	return line
}

// Scatter makes a sactter plot of the two vectors
func Scatter(x, y []float64) (*plotter.Scatter, error) {
	if len(x) != len(y) {
		panic("quickplot: slice length mismatch")
	}
	pts := make(plotter.XYs, len(x))
	for i := range pts {
		pts[i].X = x[i]
		pts[i].Y = y[i]
	}
	return plotter.NewScatter(pts)
}

func Contour(x, y, z []float64) (*myplot.ColoredScatter, error) {
	if len(x) != len(y) {
		panic("quickplot: slice length mismatch")
	}
	if len(y) != len(z) {
		panic("quickplot: slice length mismatch")
	}
	pts := make(plotter.XYZs, len(x))
	for i, v := range x {
		pts[i].X = v
		pts[i].Y = y[i]
		pts[i].Z = z[i]
	}
	return myplot.NewColoredScatter(pts)
}
