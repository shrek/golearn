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
}

func main() {
	fmt.Println("hello bandit\n")
	//bandit.PlotNormal()
	k, err := strconv.Atoi(os.Args[1])
	if err != nil {
		Usage()
		panic("expected arg 1 to be int")
	}
	alpha, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		Usage()
		panic("expected arg w to be float")
	}
	bandit.Run(k, alpha)
}
