package day12

import (
	"github.com/tmeire/adventofcode/io"
)

type plot struct {
	plant byte
	n     int // other neighbours
	r     int // index of the region

	ts, bs, ls, rs int
}

func color(land [][]*plot, i, j, r int) {
	if land[i][j].r != -1 {
		return
	}

	land[i][j].r = r
	if i > 0 && land[i-1][j].plant == land[i][j].plant {
		color(land, i-1, j, r)
	}
	if i < len(land)-1 && land[i+1][j].plant == land[i][j].plant {
		color(land, i+1, j, r)
	}
	if j > 0 && land[i][j-1].plant == land[i][j].plant {
		color(land, i, j-1, r)
	}
	if j < len(land[i])-1 && land[i][j+1].plant == land[i][j].plant {
		color(land, i, j+1, r)
	}
}

func Solve() {
	grid, err := io.ReadByteLinesFromFile("./2024/day12/input.txt")
	if err != nil {
		panic(err)
	}

	land := make([][]*plot, 0, len(grid))
	for _, row := range grid {
		var l []*plot
		for _, cell := range row {
			l = append(l, &plot{plant: cell, r: -1})
		}
		land = append(land, l)
	}

	areas := make([]int, len(grid)*len(grid[0]))

	// Add an index to track each of the regions
	var nextIndex int
	for i := 0; i < len(land); i++ {
		for j := 0; j < len(land[i]); j++ {
			if land[i][j].r == -1 {
				color(land, i, j, nextIndex)
				nextIndex++
			}
		}
	}

	for i := 0; i < len(land); i++ {
		for j := 0; j < len(land[i]); j++ {
			areas[land[i][j].r]++
		}
	}

	fences := make([]int, len(grid)*len(grid[0]))

	// Calculate the number of neighbours with other plants
	for i := 0; i < len(land); i++ {
		for j := 0; j < len(land[i]); j++ {
			if i > 0 {
				if land[i-1][j].plant != land[i][j].plant {
					land[i][j].n++
				}
			} else {
				land[i][j].n++
			}
			if i < len(land)-1 {
				if land[i+1][j].plant != land[i][j].plant {
					land[i][j].n++
				}
			} else {
				land[i][j].n++
			}
			if j > 0 {
				if land[i][j-1].plant != land[i][j].plant {
					land[i][j].n++
				}
			} else {
				land[i][j].n++
			}
			if j < len(land[i])-1 {
				if land[i][j+1].plant != land[i][j].plant {
					land[i][j].n++
				}
			} else {
				land[i][j].n++
			}
			fences[land[i][j].r] += land[i][j].n
		}
	}
	price := 0
	for i := 0; i < nextIndex; i++ {
		price += areas[i] * fences[i]
	}
	println(price)

	// Calculate the number of sides for each area
	var nextSide int = 1
	// Top
	for i := 0; i < len(land); i++ {
		for j := 0; j < len(land[i]); j++ {
			if i > 0 && land[i-1][j].r == land[i][j].r {
				// if the plot above is in the same region, this is not a border
				continue
			}

			if j == 0 {
				// we're on the edge of the land, let's consider it a new border
				land[i][j].ts = nextSide
				nextSide++
			} else if land[i][j-1].r != land[i][j].r {
				// we're in a different region, lets' consider it a new border
				land[i][j].ts = nextSide
				nextSide++
			} else if land[i][j-1].ts == 0 {
				// we're in the same region, but the neighbour goes up while this one doesn't
				land[i][j].ts = nextSide
				nextSide++
			} else {
				// fill in the same border id as the neighbour
				land[i][j].ts = land[i][j-1].ts
			}
		}
	}
	// Bottom
	for i := 0; i < len(land); i++ {
		for j := 0; j < len(land[i]); j++ {
			if i < len(land)-1 && land[i+1][j].r == land[i][j].r {
				// if the plot below is in the same region, this is not a border
				continue
			}

			if j == 0 {
				// we're on the edge of the land, let's consider it a new border
				land[i][j].bs = nextSide
				nextSide++
			} else if land[i][j-1].r != land[i][j].r {
				// we're in a different region, lets' consider it a new border
				land[i][j].bs = nextSide
				nextSide++
			} else if land[i][j-1].bs == 0 {
				// we're in the same region, but the neighbour goes down while this one doesn't
				land[i][j].bs = nextSide
				nextSide++
			} else {
				// fill in the same border id as the neighbour
				land[i][j].bs = land[i][j-1].bs
			}
		}
	}
	// Left
	for i := 0; i < len(land); i++ {
		for j := 0; j < len(land[i]); j++ {
			if j > 0 && land[i][j-1].r == land[i][j].r {
				// if the plot to the left is in the same region, this is not a border
				continue
			}

			if i == 0 {
				// we're on the edge of the land, let's consider it a new border
				land[i][j].ls = nextSide
				nextSide++
			} else if land[i-1][j].r != land[i][j].r {
				// we're in a different region, lets' consider it a new border
				land[i][j].ls = nextSide
				nextSide++
			} else if land[i-1][j].ls == 0 {
				// we're in the same region, but the neighbour goes left while this one doesn't
				land[i][j].ls = nextSide
				nextSide++
			} else {
				// fill in the same border id as the neighbour
				land[i][j].ls = land[i-1][j].ls
			}
		}
	}
	// Right
	for i := 0; i < len(land); i++ {
		for j := 0; j < len(land[i]); j++ {
			if j < len(land[i])-1 && land[i][j+1].r == land[i][j].r {
				// if the plot to the right is in the same region, this is not a border
				continue
			}

			if i == 0 {
				// we're on the edge of the land, let's consider it a new border
				land[i][j].rs = nextSide
				nextSide++
			} else if land[i-1][j].r != land[i][j].r {
				// we're in a different region, lets' consider it a new border
				land[i][j].rs = nextSide
				nextSide++
			} else if land[i-1][j].rs == 0 {
				// we're in the same region, but the neighbour goes right while this one doesn't
				land[i][j].rs = nextSide
				nextSide++
			} else {
				// fill in the same border id as the neighbour
				land[i][j].rs = land[i-1][j].rs
			}
		}
	}

	sides := make([]map[int]struct{}, nextIndex)
	for i := 0; i < nextIndex; i++ {
		sides[i] = make(map[int]struct{})
	}

	for i := 0; i < len(land); i++ {
		for j := 0; j < len(land[i]); j++ {
			if land[i][j].ts != 0 {
				sides[land[i][j].r][land[i][j].ts] = struct{}{}
			}
			if land[i][j].bs != 0 {
				sides[land[i][j].r][land[i][j].bs] = struct{}{}
			}
			if land[i][j].ls != 0 {
				sides[land[i][j].r][land[i][j].ls] = struct{}{}
			}
			if land[i][j].rs != 0 {
				sides[land[i][j].r][land[i][j].rs] = struct{}{}
			}
		}
	}
	newprice := 0
	for i := 0; i < nextIndex; i++ {
		newprice += areas[i] * len(sides[i])
	}
	println(newprice)

}
