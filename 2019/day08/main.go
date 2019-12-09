package main

import (
	"io/ioutil"
	"math"
)

const (
	w = 25
	h = 6
	//w = 3
	//h = 2
)

func main() {
	//data := []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0', '1', '2'}
	data, err := ioutil.ReadFile("data.txt")
	if err != nil {
		panic(err.Error())
	}

	layers := make([][]byte, 0, len(data)/(w*h))
	for i := 0; i < len(data); i += w * h {
		layers = append(layers, data[i:i+(w*h)])
	}
	println(len(layers))

	maxZ := math.MaxInt64
	maxZL := 0
	for i := 0; i < len(layers); i++ {
		zeros := 0
		for a := 0; a < w*h; a++ {
			if layers[i][a] == '0' {
				zeros++
			}
		}
		if zeros < maxZ {
			maxZ = zeros
			maxZL = i
			println(maxZL, maxZ)
		}
	}
	println(maxZL)
	ones, twos := 0, 0
	for a := 0; a < w*h; a++ {
		switch layers[maxZL][a] {
		case '1':
			ones++
		case '2':
			twos++
		}
	}
	println(ones * twos)

	img := make([]byte, w*h)
	for x := len(layers) - 1; x >= 0; x-- {
		for i, p := range layers[x] {
			if p != '2' {
				img[i] = p
			}
		}
	}

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if img[y*w+x] == '0' {
				print("▓")
			} else {
				print(" ")
			}
		}
		println()
	}
}
