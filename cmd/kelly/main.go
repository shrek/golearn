package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("#args: %d, args: %v\n", len(os.Args), os.Args)
}
