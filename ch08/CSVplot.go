package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/Arafatk/glot"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	defer func() {
		if c := recover(); c != nil {
			fmt.Println(c)
			fmt.Printf("Usage: %s <file1> [<file2> ...]\n", filepath.Base(os.Args[0]))
			return
		}
	}()
	flag.Parse()

	dimensions := 2
	persist := true
	debug := false
	style := "circle"
	for _, filename := range flag.Args() {
		_, err := os.Stat(filename)
		if err != nil {
			panic(err)
		}

		f, err := os.Open(filename)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		r := csv.NewReader(f)
		r.FieldsPerRecord = -1
		records, err := r.ReadAll()
		if err != nil {
			panic(err)
		}

		var points [][]float64
		var xs, ys []float64
		for _, record := range records {
			if x, err := strconv.ParseFloat(record[0], 64); err != nil {
				panic(err)
			} else {
				xs = append(xs, x)
			}
			if y, err := strconv.ParseFloat(record[1], 64); err != nil {
				panic(err)
			} else {
				ys = append(ys, y)
			}
		}
		points = append(points, xs, ys)
		plot, err := glot.NewPlot(dimensions, persist, debug)
		if err != nil {
			panic(err)
		}
		plot.SetTitle("Using Glot with CSV data")
		plot.SetXLabel("X")
		plot.SetYLabel("Y")
		plot.AddPointGroup("Circle:", style, points)
		if err := plot.SavePlot("output.png"); err != nil {
			panic(err)
		}

		fmt.Println(points)
		fmt.Println(plot)
	}

}
