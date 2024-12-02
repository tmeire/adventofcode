package main

type point struct {
	x, y int
}

func main() {
	data,err := ioutil.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}

	points := make([]point, 0, len(data))
	x, y := 0, 0
	for i, b := range data {
		switch b {
		case ' ':
			break
		case '#':
			//
		case '\n':
			y++
		}
		x++
		points
	}


}
