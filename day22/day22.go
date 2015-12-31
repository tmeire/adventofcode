package main

import (
	"fmt"
	"math"
)

const (
	WIN  = iota
	TIE  = iota
	LOSS = iota
)

const (
	MAGIC_MISSILE = iota
	DRAIN         = iota
	SHIELD        = iota
	POISON        = iota
	RECHARGE      = iota
)

type Player struct {
	health int
	armor  int
	damage int
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func evaluate(spells []int) int {
	m := make([]int, len(spells)*2+6)
	d := make([]int, len(spells)*2+6)
	a := make([]int, len(spells)*2+6)
	h := make([]int, len(spells)*2+6)

	for i, s := range spells {
		turn := i * 2

		switch s {
		case MAGIC_MISSILE:
			m[turn] -= 53
			d[turn] += 4
		case DRAIN:
			m[turn] -= 73
			d[turn] += 2
			h[turn] += 2
		case SHIELD:
			m[turn] -= 113
			for j := 1; j < 7; j += 1 {
				a[turn+j] += 7
				if j > 1 && j < 6 && (turn+j)/2 < len(spells) && spells[(turn+j)/2] == SHIELD {
					return LOSS
				}
			}
		case POISON:
			m[turn] -= 173
			for j := 1; j < 7; j += 1 {
				d[turn+j] += 3
				if j > 1 && j < 6 && (turn+j)/2 < len(spells) && spells[(turn+j)/2] == POISON {
					return LOSS
				}
			}
		case RECHARGE:
			m[turn] -= 229
			for j := 1; j < 6; j += 1 {
				m[turn+j] += 101
				if j > 1 && j < 5 && (turn+j)/2 < len(spells) && spells[(turn+j)/2] == RECHARGE {
					return LOSS
				}
			}
		}
	}

	p := Player{50, 0, 0}
	b := Player{71, 0, 10}

	mana := 500
	for i := 0; i < len(spells)*2; i += 1 {
		mana += m[i]
		if mana < 0 {
			return LOSS
		}

		// solution b
		//p.health -= 1
		//if p.health <= 0 {
		//	return LOSS
		//}

		p.health += h[i]
		b.health -= d[i]
		if b.health <= 0 {
			return WIN
		}

		i += 1

		mana += m[i]
		if mana < 0 {
			return LOSS
		}

		p.health += h[i]
		b.health -= d[i]
		if b.health <= 0 {
			return WIN
		}

		p.health -= max(1, (b.damage - a[i]))
		if p.health <= 0 {
			return LOSS
		}
	}
	return TIE
}

func cost(spells []int) int {
	c := 0
	for _, spell := range spells {
		switch spell {
		case MAGIC_MISSILE:
			c += 53
		case DRAIN:
			c += 73
		case SHIELD:
			c += 113
		case POISON:
			c += 173
		case RECHARGE:
			c += 229
		}
	}
	return c
}

func test(spells []int, best int) int {
	c := cost(spells)
	if c >= best {
		return -1
	}

	switch evaluate(spells) {
	case WIN:
		fmt.Println(spells, c)
		return c
	case LOSS:
		return -1
	case TIE:
		x := test(append(spells, MAGIC_MISSILE), best)
		if x > -1 && x < best {
			fmt.Println("new best", x, spells)
			best = x
		}
		x = test(append(spells, DRAIN), best)
		if x > -1 && x < best {
			fmt.Println("new best", x, spells)
			best = x
		}
		x = test(append(spells, SHIELD), best)
		if x > -1 && x < best {
			fmt.Println("new best", x, spells)
			best = x
		}
		x = test(append(spells, POISON), best)
		if x > -1 && x < best {
			fmt.Println("new best", x, spells)
			best = x
		}
		x = test(append(spells, RECHARGE), best)
		if x > -1 && x < best {
			fmt.Println("new best", x, spells)
			best = x
		}
		return best
	}
	return -1
}

func main() {
	fmt.Println(test(make([]int, 0), math.MaxInt32))
}
