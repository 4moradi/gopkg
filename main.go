package main

import (
	"fmt"

	"github.com/4moradi/gopkg/formatter"
	"github.com/4moradi/gopkg/math"
)

func main() {
	num := math.Double(2)
	output := print.Format(num)
	fmt.Println(output)
}
