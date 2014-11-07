package main

import (
	"fmt"
	"log"
	"os"

	"code.google.com/p/plotinum/plot"
	"code.google.com/p/plotinum/vg"

	"github.com/btracey/numcsv"
	"github.com/btracey/quickplot"
)

func main() {
	// csv name
	dataname := "/Users/brendan/Documents/mygo/data/ransuq/flatplate/med/Flatplate_Re_3e_06/turb_flatplate_sol.dat"
	xHeadingName := "OmegaBar"
	yHeadingName := "Chi"

	plotname := "plot.pdf"

	datafile, err := os.Open(dataname)
	if err != nil {
		log.Fatal(err)
	}

	reader := numcsv.NewReader(datafile)
	reader.Comma = "\t"

	headings, err := reader.ReadHeading()
	if err != nil {
		log.Fatal(err)
	}

	xIdx, yIdx := -1, -1
	for i, str := range headings {
		if str == xHeadingName {
			xIdx = i
		}
		if str == yHeadingName {
			yIdx = i
		}
	}
	if xIdx == -1 || yIdx == -1 {
		log.Fatal("missing heading")
	}

	fmt.Println(headings)

	data, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)

	scatter, err := quickplot.ScatterFromColumns(data, xIdx, yIdx)
	if err != nil {
		log.Fatal(err)
	}

	scatter.GlyphStyle.Radius = vg.Centimeters(.1)
	scatter.GlyphStyle.Shape = plot.CircleGlyph{}

	p, _ := plot.New()
	p.Add(scatter)

	p.Save(4, 4, plotname)
	fmt.Println(headings)
}
