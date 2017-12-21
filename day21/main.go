package main

import "strings"
import "fmt"

func main() {
	start := makeGrid(startingGrid)
	expandedRules := expandRules(testInput)
	result := iterate(start, expandedRules)
	fmt.Println(result.key())
}

func iterate(pg pixelGrid, rules map[string]string) pixelGrid {
	var subGridLength int
	if len(pg)%2 == 0 {
		subGridLength = 2
	} else {
		subGridLength = 3
	}
	numSubGrids := len(pg) / subGridLength
	allSubgrids := []pixelGrid{}
	for i := 0; i < numSubGrids; i++ {
		for j := 0; j < numSubGrids; j++ {
			subGrid := pg.getSubGrid(i, j, subGridLength)
			result, foundMatch := rules[subGrid.key()]
			if !foundMatch {
				panic("eek no match for subgrid: " + subGrid.key())
			}
			subGrid = makeGrid(result)
			allSubgrids = append(allSubgrids, subGrid)
		}
	}
	return joinGrids(allSubgrids, numSubGrids*subGridLength)
}
func expandRules(in map[string]string) map[string]string {
	out := map[string]string{}
	for k, v := range in {
		grid := makeGrid(k)
		gridOptions := grid.makeKeys()
		for _, s := range gridOptions {
			out[s] = v
		}
	}
	return out
}

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

	for y := 0; y < size; y++ {
		out = append(out, make([]bool, size))
		for x := 0; x < size; x++ {
			subGridIndex := (y * subSize) + (x % subSize)
			subGridX := (x % subSize)
			subGridY := (y % subSize)
			fmt.Println(subGridIndex, subGridY, subGridX)
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

func (pg pixelGrid) getSubGrid(x, y, length int) pixelGrid {
	out := pixelGrid{}
	for yy := y; yy < y+length; yy++ {
		out = append(out, make([]bool, length))
		for xx := x; xx < x+length; xx++ {
			out[yy-y][xx-x] = pg[yy][xx]
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
