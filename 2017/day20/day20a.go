package day20

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const MAX_INT64 = int64(^uint64(0) >> 1)

func abs(a int64) int64 {
	if a < 0 {
		return -a
	}
	return a
}

func move(x, v, a, t int64) int64 {
	return x + t*v + t*(t+1)/2*a
}

type triple struct {
	x, y, z int64
}

type particle struct {
	p, v, a triple
	dead    bool
}

func (p particle) Position(t int64) (int64, int64, int64) {
	return move(p.p.x, p.v.x, p.a.x, t), move(p.p.y, p.v.y, p.a.y, t), move(p.p.z, p.v.z, p.a.z, t)
}

func (p particle) Distance(t int64) int64 {
	return abs(move(p.p.x, p.v.x, p.a.x, t)) + abs(move(p.p.y, p.v.y, p.a.y, t)) + abs(move(p.p.z, p.v.z, p.a.z, t))
}

func (p particle) Collides(po particle, t int64) bool {
	if p.dead || po.dead {
		return false
	}
	x1, y1, z1 := p.Position(t)
	x2, y2, z2 := po.Position(t)

	return x1 == x2 && y1 == y2 && z1 == z2
}

func load(fname string) []particle {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	particles := []particle{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var p particle
		fmt.Sscanf(scanner.Text(), "p=<%d,%d,%d>, v=<%d,%d,%d>, a=<%d,%d,%d>", &(p.p.x), &(p.p.y), &(p.p.z), &(p.v.x), &(p.v.y), &(p.v.z), &(p.a.x), &(p.a.y), &(p.a.z))
		//fmt.Print(p)
		particles = append(particles, p)
	}
	return particles
}

func Solve() {
	if len(os.Args) < 2 {
		panic("Must pass input filename as commandline argument.")
	}

	particles := load(os.Args[1])

	lastp := 0
	iters := 0

	minp := 0

	// The number of particles that didn't collide
	alive := len(particles)
	lastalive := alive
	aliveiters := 0

	for t := int64(0); iters < 1000 || aliveiters < 1000; t++ {
		min := MAX_INT64
		for i, p := range particles {
			if dist := p.Distance(t); dist < min {
				min = dist
				minp = i
			}

			// Part 2, check if the particle collides with others
			if !p.dead {
				collision := false
				for j := i + 1; j < len(particles); j++ {
					if !particles[j].dead && p.Collides(particles[j], t) {
						collision = true
						particles[j].dead = true
						alive--
						fmt.Println("Collision!", i, j)
					}
				}
				if collision {
					p.dead = true
					alive--
					fmt.Println("Collision!", i)
				}
			}
		}

		if minp != lastp {
			lastp = minp
			iters = 0
		} else {
			iters++
		}

		if alive != lastalive {
			lastalive = alive
			aliveiters = 0
		} else {
			aliveiters++
		}
		//fmt.Println(iters, aliveiters)
	}
	fmt.Println("Part A:", minp)
	fmt.Println("Part B:", alive)
}
