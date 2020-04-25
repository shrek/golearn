
package main

import (
	"fmt"
	"math"
	"sort"
	"log"

        "gorgonia.org/gorgonia"
        "gorgonia.org/tensor"
	
	"gonum.org/v1/gonum/stat"
)

func testTensor() {
	fmt.Println("-----Testing tensor----\n")	
	// Create a (2, 2)-Matrix of integers
	a := tensor.New(tensor.WithShape(2, 2), tensor.WithBacking([]int{1, 2, 3, 4}))
	fmt.Printf("a:\n%v\n", a)
	
	// Create a (2, 3, 4)-tensor of float32s
	b := tensor.New(tensor.WithBacking(tensor.Range(tensor.Float32, 0, 24)), tensor.WithShape(2, 3, 4))
	fmt.Printf("b:\n%1.1f", b)
	
	// Accessing data
	x, _ := b.At(0, 1, 2) // in Numpy syntax: b[0,1,2]
	fmt.Printf("x: %1.1f\n\n", x)
	
	// Setting data
	b.SetAt(float32(1000), 0, 1, 2)
	fmt.Printf("b:\n%v", b)
	
}

func testGorgonia() {
	fmt.Println("-----Starting gorgonia----\n")
	g := gorgonia.NewGraph()

        var x, y, z *gorgonia.Node
        var err error

        // define the expression
        x = gorgonia.NewScalar(g, gorgonia.Float64, gorgonia.WithName("x"))
        y = gorgonia.NewScalar(g, gorgonia.Float64, gorgonia.WithName("y"))
        if z, err = gorgonia.Add(x, y); err != nil {
                log.Fatal(err)
        }

        // create a VM to run the program on
        machine := gorgonia.NewTapeMachine(g)
        defer machine.Close()

        // set initial values then run
        gorgonia.Let(x, 2.0)
        gorgonia.Let(y, 2.5)
        if err = machine.RunAll(); err != nil {
                log.Fatal(err)
        }

        fmt.Printf("%v", z.Value())
}

func testGonum() {
	fmt.Println("------ testing gonum ------\n")

	xs := []float64{
		32.32, 56.98, 21.52, 44.32,
		55.63, 13.75, 43.47, 43.34,
		12.34,
	}

	fmt.Printf("data: %v\n", xs)

	sort.Float64s(xs)
	fmt.Printf("data: %v (sorted)\n", xs)

	// computes the weighted mean of the dataset.
	// we don't have any weights (ie: all weights are 1)
	// so we just pass a nil slice.
	mean := stat.Mean(xs, nil)

	// computes the median of the dataset.
	// here as well, we pass a nil slice as weights.
	median := stat.Quantile(0.5, stat.Empirical, xs, nil)

	variance := stat.Variance(xs, nil)
	stddev := math.Sqrt(variance)

	fmt.Printf("mean=     %v\n", mean)
	fmt.Printf("median=   %v\n", median)
	fmt.Printf("variance= %v\n", variance)
	fmt.Printf("std-dev=  %v\n", stddev)	
}

func main() {
	testTensor()
	testGonum()
	testGorgonia()
}
