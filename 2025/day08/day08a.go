package day08

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
)

func Solve() {
	points := read("./2025/day08/input.txt")
	edges := pair(points)
	sort.Slice(edges, func(i, j int) bool { return edges[i].dist < edges[j].dist })

	clusters := make([]*cluster, 0)

	for idx, e := range edges {
		if idx == 1000 {
			sort.Slice(clusters, func(i, j int) bool { return len(clusters[i].points) > len(clusters[j].points) })
			println(len(clusters[0].points) * len(clusters[1].points) * len(clusters[2].points))
		}

		c1 := find(clusters, e.p1)
		c2 := find(clusters, e.p2)
		if c1 == nil && c2 == nil {
			clusters = append(clusters, &cluster{points: map[*point]struct{}{
				e.p1: {},
				e.p2: {},
			}})
		} else if c1 != nil && c2 != nil {
			c1.merge(c2)
			if len(c1.points) == len(points) {
				println(e.p1.x * e.p2.x)
				return
			}
		} else if c1 != nil {
			c1.add(e.p2)
			if len(c1.points) == len(points) {
				println(e.p1.x * e.p2.x)
				return
			}
		} else {
			c2.add(e.p1)
			if len(c2.points) == len(points) {
				println(e.p1.x * e.p2.x)
				return
			}
		}
	}
}

func find(clusters []*cluster, p *point) *cluster {
	for _, c := range clusters {
		if c.contains(p) {
			return c
		}
	}
	return nil
}

func pair(points []*point) []edge {
	edges := make([]edge, 0, len(points)*len(points))
	for i, p1 := range points {
		for j := i + 1; j < len(points); j++ {
			edges = append(edges, edge{p1, points[j], dist(p1, points[j])})
		}
	}
	return edges
}

func pow(i int) float64 {
	return math.Pow(float64(i), 2.0)
}

func dist(p1, p2 *point) float64 {
	return math.Sqrt(pow(p1.x-p2.x) + pow(p1.y-p2.y) + pow(p1.z-p2.z))
}

type cluster struct {
	points map[*point]struct{}
}

func (c *cluster) add(p *point) {
	c.points[p] = struct{}{}
}

func (c *cluster) merge(c1 *cluster) {
	if c == c1 {
		return
	}
	for p := range c1.points {
		c.points[p] = struct{}{}
	}
	c1.points = nil
}

func (c *cluster) contains(p *point) bool {
	_, ok := c.points[p]
	return ok
}

type edge struct {
	p1, p2 *point
	dist   float64
}

type point struct {
	x, y, z int
}

func read(fn string) []*point {
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var points []*point
	bf := bufio.NewReader(f)
	for {
		l, _, err := bf.ReadLine()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return points
			}
			panic(err.Error())
		}
		var p point
		_, err = fmt.Sscanf(string(l), "%d,%d,%d", &(p.x), &(p.y), &(p.z))
		if err != nil {
			panic(err)
		}
		fmt.Printf("%#v\n", p)
		points = append(points, &p)
	}
}
