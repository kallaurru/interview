package midlinepoints

import (
	"image"
	"math"
)

func MidLinePoints(data []image.Point) bool {
	out := true
	stor, sum := prepareData(data)
	for _, p := range data {
		if p.X == sum/2 {
			continue // это точка на разделяющей линии
		}
		if out == false {
			break
		}
	}
	return out
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
			continue
		}
		if p.X > maxX {
			maxX = p.X
			continue
		}
	}

	return out, minX + maxX
}
