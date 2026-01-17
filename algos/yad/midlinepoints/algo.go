package midlinepoints

import (
	"image"
	"math"
)

func MidLinePoints(data []image.Point) bool {
	stor, sum := prepareData(data)
	if len(data) == 1 {
		return false
	}
	for _, p := range data {
		if p.X == sum/2 {
			continue // это точка на разделяющей линии
		}
		sim := image.Pt(sum-p.X, p.Y)
		countSim := stor[sim]
		count := stor[p]
		if count == countSim {
			continue
		}
		return false
	}
	return true
}

func prepareData(data []image.Point) (map[image.Point]int, int) {
	out := make(map[image.Point]int, len(data))
	minX := math.MaxInt
	maxX := math.MinInt
	for _, p := range data {
		_, ok := out[p]
		if ok {
			out[p] += 1
		} else {
			out[p] = 1
		}
		if p.X < minX {
			minX = p.X
		}
		if p.X > maxX {
			maxX = p.X
			continue
		}
	}

	return out, minX + maxX
}
