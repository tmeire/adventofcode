package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

//const l1 = "L2,R2"
//const l2 = "R2"

const l1 = "R1008,D256,L88,U390,R429,D828,R2,D452,L644,D942,R387,U221,L274,D837,R437,U664,R952,U126,L840,U425,L749,D199,L48,U394,L623,D562,L760,D856,L648,U666,R756,U396,L588,U217,R208,D492,L230,U60,L178,D211,L806,U423,L399,D159,L176,D555,R173,D946,L360,U415,L734,D441,L146,D332,R135,D529,L364,U742,L862,D790,L399,D392,R706,D740,L839,D950,R822,D27,R108,D873,L492,D465,L635,U771,L586,U66,R703,U943,R141,U396,R641,D339,R460,U295,L397,U799,R479,U963,L211,U933,R158,U248,R443,U807,R115,U885,R670,U116,L24,D980,R349,U363,L413,U444,L453,D497,L202,U300,L122,D895,L210,U218,R456,U293,L576,U968,L612,D225,L732,D34,R800,U925,R731,U520,R686,D181,L102,D824,R832,D527,L614,D624,R734,U552,L911,D352,R157,D70,R958,U317,L43,U902,R265,U986,R305,U264,L957,U888,R66,D413,L73,D642,R14,D559,R414,D985,R679,D965,R333,D332,L261,D446,L479,U430,L730,D37,L936,D615,L344,D215,R912,D95,L691,U383,L328,U560,R806,U711,R515,U448,R403,D109,L589,U458,L240,D375,L88,D479,R93,U794,L303,U783,L833,U500,R406,D589,L694,U504,L484,U695,R228,U813,R646,U768,L60,D326,L580,U840,L387,U147,L50,U155,L454,D574,L885,D705,R727,D827,R409,U335,L271,D388,R897,D563,L360,U70,R777,U903,R363,D202,R855,D159,R35,U585,L384,D540,R78,U13,R979,D702,L868,D868,R508,D552,L735,U923,R840,U133,L355,U282,R626,D699,L560,D26,R902,D873,L333,U492,L96,U461,R261,U784,L793,D723,R887,U836,R790,D400,L331,U389,L107,U534,L377,D831,R181,U325,L328,U778,L498,D109,L692,U185,R284,U930,R784,D843,L261,U119,L751,U313,R197,U911,L21,D201,L881,U119,R210,D700,R93,U208,R116"
const l2 = "L1009,D700,L634,U294,R898,D947,R650,U988,L623,D968,R761,U490,R525,U76,R633,D139,R348,D855,L983,U553,L454,D211,L240,D465,R260,U285,R653,D734,L346,U434,R813,U599,R98,D779,L58,D6,R309,U437,L712,U896,R262,U911,R400,D247,R297,U915,L223,D591,L755,D398,L980,U177,R186,U882,R418,U763,L741,D60,L942,U648,L430,D401,R30,D157,L901,D179,L615,U535,L586,D613,L606,U239,L133,D251,L549,D579,R612,U307,L127,D343,L43,D288,R245,U122,R352,D527,R476,U24,L805,U291,R953,D469,L941,U577,L384,U345,L463,D50,R311,D649,L746,D902,R644,U913,R147,D649,R848,D673,R93,U65,R363,U734,L289,U599,R738,U45,R128,D508,L93,D201,R51,U239,R17,D496,L661,D912,R165,U291,L207,D308,R241,D388,L910,D821,R714,D327,L605,U880,L682,D934,R334,D1,R602,D54,L51,D913,L575,U168,R614,D603,R452,U718,R689,D505,R83,D385,R636,D692,R424,D573,R686,D572,L467,D698,L21,U510,L497,U329,R286,U733,R584,U919,R499,U971,L558,U511,L565,D623,L502,U536,L483,U372,L686,D420,L900,U316,L37,U372,L915,D641,R165,D927,L137,U231,R813,U416,L131,D530,R486,U795,L507,U757,L208,U308,L521,U583,L758,U654,R554,D467,R381,U155,R47,U829,R92,D158,R54,D500,R17,D471,R748,U571,L194,D55,L921,U271,L730,U207,L204,U806,R19,D33,R218,D911,L106,U220,R551,U308,L268,U5,L374,U657,R639,U705,R294,U962,L927,U892,L477,U290,R378,D193,L154,U859,L618,D690,L769,D779,R752,D915,L693,D586,L558,D864,L523,U354,R386,U236,R888,U302,L75,U628,R132,D306,L939,U73,L687,D488,R21,D760,L856,U96,L116,U557,L639,U812,L389,D364,L807,U696,R781,D625,R565,U728,R134,D406,R785,U583,R60,D819,L939"

//const l1 = "R75,D30,R83,U83,L12,D49,R71,U7,L72"
//const l2 = "U62,R66,U55,R34,D71,R55,D58,R83"

//const l1 = "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51"
//const l2 = "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"

//const l1 = "R8,U5,L5,D3"
//const l2 = "U7,R6,D4,L4"

type direction byte

const (
	UP direction = iota
	LEFT
	DOWN
	RIGHT
)

type step struct {
	d direction
	l int
}

func parse(l string) []step {
	rawsteps := strings.Split(l, ",")

	steps := make([]step, 0, len(rawsteps))
	for _, rawstep := range rawsteps {
		var d direction
		switch rawstep[0] {
		case 'U':
			d = UP
		case 'L':
			d = LEFT
		case 'D':
			d = DOWN
		case 'R':
			d = RIGHT
		}

		l, err := strconv.Atoi(rawstep[1:len(rawstep)])
		if err != nil {
			fmt.Println(err)
			return nil
		}

		steps = append(steps, step{d, l})
	}
	return steps
}

type intgrid struct {
	g map[int]map[int]int
}

func (grid *intgrid) inc(x, y, flag int) {
	if _, ok := grid.g[x]; !ok {
		grid.g[x] = make(map[int]int)
	}
	grid.g[x][y] |= flag
}

func (grid *intgrid) set(x, y, v int) {
	if _, ok := grid.g[x]; !ok {
		grid.g[x] = make(map[int]int)
	}
	grid.g[x][y] = v
}

func (grid *intgrid) get(x, y int) int {
	if col, ok := grid.g[x]; ok {
		if v, ok := col[y]; ok {
			return v
		}
	}
	return -1
}

func traceWithFlag(wires *intgrid, steps []step, flag int) {
	x, y := 0, 0
	for _, step := range steps 3
		switch step.d 
		case UP:
			for i := 0; i < step.l; i++ {
				y++
				wires.inc(x, y, flag)
			}
		case LEFT:
			for i := 0; i < step.l; i++ {
				x--
				wires.inc(x, y, flag)
			}
		case DOWN:
			for i := 0; i < step.l; i++ {
				y--
				wires.inc(x, y, flag)
			}
		case RIGHT:
			for i := 0; i < step.l; i++ {
				x++
				wires.inc(x, y, flag)
			}
		}
	}
}

func traceWithCount(wires *intgrid, steps []step) {
	x, y := 0, 0
	stepcount := 0
	for _, step := range steps {
		switch step.d {
		case UP:
			for i := 0; i < step.l; i++ {
				stepcount++
				y++
				wires.set(x, y, stepcount)
			}
		case LEFT:
			for i := 0; i < step.l; i++ {
				stepcount++
				x--
				wires.set(x, y, stepcount)
			}
		case DOWN:
			for i := 0; i < step.l; i++ {
				stepcount++
				y--
				wires.set(x, y, stepcount)
			}
		case RIGHT:
			for i := 0; i < step.l; i++ {
				stepcount++
				x++
				wires.set(x, y, stepcount)
			}
		}
	}
}

func main() {
	wires := &intgrid{make(map[int]map[int]int)}

	l1steps := parse(l1)
	traceWithFlag(wires, l1steps, 1<<1)

	l2steps := parse(l2)
	traceWithFlag(wires, l2steps, 1<<2)

	v := 1 << 1
	v |= 1 << 2

	closest := math.MaxInt64
	for i, col := range wires.g {
		for j, count := range col {
			if count == v {
				dist := abs(i) + abs(j)
				if dist < closest {
					closest = dist
				}
			}
		}
	}
	println(closest)

	wire1 := &intgrid{make(map[int]map[int]int)}
	traceWithCount(wire1, l1steps)
	wire2 := &intgrid{make(map[int]map[int]int)}
	traceWithCount(wire2, l2steps)

	shortest := math.MaxInt64
	for i, col := range wire1.g {
		for j, w1count := range col {
			w2count := wire2.get(i, j)
			if w2count != -1 && w1count+w2count < shortest {
				shortest = w1count + w2count
			}
		}
	}
	println(shortest)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
