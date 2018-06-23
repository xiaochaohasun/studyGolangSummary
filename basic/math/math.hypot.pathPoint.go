package main

import (
	"fmt"
	"math"
)

type Point struct {
	X float64
	Y float64
}

type Path []Point

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func pathDistance(p Path) float64 {
	sum := 0.0

	for i := range p {
		if i > 0 {
			sum += p[i-1].Distance(p[i])
		}
	}
	return sum
}

func main() {
	//计算多个点之间的距离
	pp := Path{Point{-1, 0}, {0, 0}, {3, 4}}
	fmt.Println(pathDistance(pp))
}
