package main

import (
	"fmt"
	"math"
)

type point struct {
	X float64
	Y float64
}


func Distance(p, q point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}


func main() {
	p := point{0, 0}
	q := point{3, 4}

	fmt.Println(Distance(q, p))

}
