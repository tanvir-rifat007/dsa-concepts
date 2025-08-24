package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

// Linear Search
func linearSearch(arr []int, target int) int {
	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return -1
}

// Binary Search
func binarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Generate sorted array
	maxN := 10000
	step := 500
	arr := make([]int, maxN)
	for i := 0; i < maxN; i++ {
		arr[i] = i
	}

	linearPoints := make(plotter.XYs, 0)
	binaryPoints := make(plotter.XYs, 0)

	for n := step; n <= maxN; n += step {
		subArr := arr[:n]
		target := subArr[rand.Intn(n)]

		// Measure linear time
		start := time.Now()
		linearSearch(subArr, target)
		linearTime := time.Since(start).Nanoseconds()

		// Measure binary time
		start = time.Now()
		binarySearch(subArr, target)
		binaryTime := time.Since(start).Nanoseconds()

		linearPoints = append(linearPoints, plotter.XY{X: float64(n), Y: math.Log(float64(linearTime + 1))})
		binaryPoints = append(binaryPoints, plotter.XY{X: float64(n), Y: math.Log(float64(binaryTime + 1))})

		fmt.Printf("N=%d, Linear=%dns, Binary=%dns\n", n, linearTime, binaryTime)
	}

	// Plot
	p := plot.New()
	p.Title.Text = "Linear Search vs Binary Search"
	p.X.Label.Text = "Input Size (N)"
	p.Y.Label.Text = "Log(Time in ns)"

	linearLine, err := plotter.NewLine(linearPoints)
	if err != nil {
		log.Fatal(err)
	}
	linearLine.Color = plotutil.Color(1)
	linearLine.Width = vg.Points(2)
	linearLine.Dashes = []vg.Length{vg.Points(5), vg.Points(5)}

	binaryLine, err := plotter.NewLine(binaryPoints)
	if err != nil {
		log.Fatal(err)
	}
	binaryLine.Color = plotutil.Color(2)
	binaryLine.Width = vg.Points(2)

	p.Add(linearLine, binaryLine)
	p.Legend.Add("Linear Search", linearLine)
	p.Legend.Add("Binary Search", binaryLine)

	if err := p.Save(8*vg.Inch, 6*vg.Inch, "search_comparison.png"); err != nil {
		log.Fatal(err)
	}
}
