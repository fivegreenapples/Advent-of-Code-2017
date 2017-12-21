package main

import "strings"

type pixelGrid [][]bool

func makeGrid(in string) pixelGrid {
	bits := strings.Split(in, "/")
	grid := pixelGrid{}
	for _, b := range bits {
		grid = append(grid, pixelsToBools(b))
	}
	return grid
}
func joinGrids(in []pixelGrid, size int) pixelGrid {

	out := pixelGrid{}
	subSize := len(in[0])
	numSubsPerSide := size / subSize

	for y := 0; y < size; y++ {
		out = append(out, make([]bool, size))
		for x := 0; x < size; x++ {
			subGridIndex := ((y / subSize) * numSubsPerSide) + (x / subSize)
			subGridX := (x % subSize)
			subGridY := (y % subSize)
			out[y][x] = in[subGridIndex][subGridY][subGridX]
		}
	}

	return out
}

func (pg pixelGrid) key() string {
	pixels := make([]string, len(pg))
	for i, row := range pg {
		pixels[i] = boolsToPixels(row)
	}
	return strings.Join(pixels, "/")
}
func (pg pixelGrid) String() string {
	out := make([]string, len(pg))
	for i, row := range pg {
		out[i] = boolsToPixels(row)
	}
	return strings.Join(out, "\n")
}
func (pg pixelGrid) numPixelsOn() int {
	out := 0
	for _, row := range pg {
		for _, val := range row {
			if val {
				out++
			}
		}
	}
	return out
}

func (pg pixelGrid) getSubGrid(x, y, length int) pixelGrid {
	out := pixelGrid{}
	for yy := y * length; yy < (y*length)+length; yy++ {
		out = append(out, make([]bool, length))
		for xx := x * length; xx < (x*length)+length; xx++ {
			out[yy-(y*length)][xx-(x*length)] = pg[yy][xx]
		}
	}
	return out
}

func (pg pixelGrid) makeKeys() []string {
	one := pg.rotate90()
	two := one.rotate90()
	three := two.rotate90()

	flipped := pg.flipHorizontal()
	flippedOne := flipped.rotate90()
	flippedTwo := flippedOne.rotate90()
	flippedThree := flippedTwo.rotate90()

	allKeys := []string{}
	allKeys = append(allKeys, pg.key())
	allKeys = append(allKeys, one.key())
	allKeys = append(allKeys, two.key())
	allKeys = append(allKeys, three.key())
	allKeys = append(allKeys, flipped.key())
	allKeys = append(allKeys, flippedOne.key())
	allKeys = append(allKeys, flippedTwo.key())
	allKeys = append(allKeys, flippedThree.key())

	return allKeys
}

func (pg pixelGrid) rotate90() pixelGrid {
	out := pixelGrid{}
	size := len(pg)
	for y, row := range pg {
		for x, val := range row {
			if y == 0 {
				out = append(out, make([]bool, size))
			}

			out[x][size-1-y] = val
		}
	}
	return out
}

func (pg pixelGrid) flipHorizontal() pixelGrid {
	out := make([][]bool, len(pg))
	for i := range pg {
		out[i] = reverseBools(pg[i])
	}
	return out
}

func reverseBools(in []bool) []bool {
	out := make([]bool, len(in))
	for i, j := 0, len(in)-1; j >= 0; i, j = i+1, j-1 {
		out[i] = in[j]
	}
	return out
}

func pixelsToBools(in string) []bool {
	out := []bool{}
	for _, r := range in {
		out = append(out, r == '#')
	}
	return out
}
func boolsToPixels(in []bool) string {
	out := ""
	for _, r := range in {
		if r {
			out += "#"
		} else {
			out += "."
		}
	}
	return out
}
