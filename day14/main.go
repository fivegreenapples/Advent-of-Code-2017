package main

import (
	"fmt"

	"github.com/fivegreenapples/adventofcode2017/day14/knothash"
)

func main() {
	// fmt.Println("Part1", bitsSetForMemory(input))

	testMemory := makeMemory(testInput)
	fmt.Println("Test (bits set):", testMemory.numberBitsSet())
	fmt.Println("Test (regions):", testMemory.numberContigiousRegions())

	answerMemory := makeMemory(input)
	fmt.Println("Answer (bits set):", answerMemory.numberBitsSet())
	fmt.Println("Answer (regions):", answerMemory.numberContigiousRegions())
}

type memBlock [4]bool
type memoryRow [128]bool
type memory [128]memoryRow
type memIndex [2]int

var blocksByVal = map[int]memBlock{
	0:  memBlock{false, false, false, false},
	1:  memBlock{false, false, false, true},
	2:  memBlock{false, false, true, false},
	3:  memBlock{false, false, true, true},
	4:  memBlock{false, true, false, false},
	5:  memBlock{false, true, false, true},
	6:  memBlock{false, true, true, false},
	7:  memBlock{false, true, true, true},
	8:  memBlock{true, false, false, false},
	9:  memBlock{true, false, false, true},
	10: memBlock{true, false, true, false},
	11: memBlock{true, false, true, true},
	12: memBlock{true, true, false, false},
	13: memBlock{true, true, false, true},
	14: memBlock{true, true, true, false},
	15: memBlock{true, true, true, true},
}

func makeMemory(key string) memory {
	mem := memory{}
	for r := 0; r < 128; r++ {
		rowKey := fmt.Sprintf("%s-%d", key, r)
		rowHash := knothash.Calculate(rowKey)
		mem[r] = hashToMemoryRow(rowHash)
	}
	return mem
}

func hashToMemoryRow(hash [16]int) memoryRow {
	row := memoryRow{}
	subBlock := memBlock{}
	j := 0
	for i := 0; i < 16; i++ {
		// bits for bits 5-8
		subBlock = blocksByVal[(hash[i]>>4)&0xf]
		row[j] = subBlock[0]
		row[j+1] = subBlock[1]
		row[j+2] = subBlock[2]
		row[j+3] = subBlock[3]
		// bits for LSB 4 bits
		subBlock = blocksByVal[hash[i]&0xf]
		row[j+4] = subBlock[0]
		row[j+5] = subBlock[1]
		row[j+6] = subBlock[2]
		row[j+7] = subBlock[3]

		j += 8
	}
	return row
}

func (mr *memoryRow) numberBitsSet() int {
	bits := 0
	for r := 0; r < 128; r++ {
		if mr[r] {
			bits++
		}
	}
	return bits
}

func (m *memory) numberBitsSet() int {
	bits := 0
	for r := 0; r < 128; r++ {
		bits += m[r].numberBitsSet()
	}
	return bits
}
func (m *memory) numberContigiousRegions() int {
	regionsByIndex := map[memIndex]int{}
	indexesByRegion := map[int][]memIndex{}
	nextRegion := 1

	for r := 0; r < 128; r++ {
		for c := 0; c < 128; c++ {

			// If not set, move on
			if !m[r][c] {
				continue
			}

			useRegion := 0

			if c > 0 && m[r][c-1] {
				// If previous bit is set, use its region
				useRegion = regionsByIndex[memIndex{r, c - 1}]
			} else if r > 0 && m[r-1][c] {
				// If bit above is set, use its region
				useRegion = regionsByIndex[memIndex{r - 1, c}]
			} else {
				// Use new region
				useRegion = nextRegion
				nextRegion++
				indexesByRegion[useRegion] = []memIndex{}
			}
			regionsByIndex[memIndex{r, c}] = useRegion
			indexesByRegion[useRegion] = append(indexesByRegion[useRegion], memIndex{r, c})

			// Now check if bit above is set and has a different region
			if r > 0 && m[r-1][c] {
				oldRegion := regionsByIndex[memIndex{r - 1, c}]
				if oldRegion != useRegion {
					// convert this region to our one
					for _, idx := range indexesByRegion[oldRegion] {
						regionsByIndex[idx] = useRegion
						indexesByRegion[useRegion] = append(indexesByRegion[useRegion], idx)
					}
					// and delete the old region
					delete(indexesByRegion, oldRegion)
				}
			}
		}
	}

	// the length of indexesByRegion is the nmber of contigious regions
	return len(indexesByRegion)
}
