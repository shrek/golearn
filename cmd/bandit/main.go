package main

import (
	"fmt"
	"os"
	"strconv"
	//"math"
	//"sort"
	//"log"

	//"gorgonia.org/gorgonia"
	//"gorgonia.org/tensor"

	//"gonum.org/v1/gonum/stat"
	"github.com/shrek/golearn/pkg/bandit"
)

func Usage() {
	fmt.Println("<command> numRuns gradientAlpha")
	fmt.Println("    bin/bandit 100 0.1")	
}

func main() {
	fmt.Println("hello bandit")
	if len(os.Args) != 3 {
		Usage()
		os.Exit(1)
	}
	//bandit.PlotNormal()
	k, err := strconv.Atoi(os.Args[1])
	if err != nil {
		Usage()
		os.Exit(1)
	}
	alpha, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		Usage()
		os.Exit(1)
	}
	bandit.Run(k, alpha)
}
