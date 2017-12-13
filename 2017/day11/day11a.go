package day11

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

func canReduceFront(in []string, i int) bool {
	for j := i + 1; j < len(in); j++ {
		if (in[i] == "n" && in[j] == "s") || (in[i] == "ne" && in[j] == "sw") || (in[i] == "nw" && in[j] == "se") {
			return true
		} else if in[i] == "n" && in[j] == "se" {
			return true
		} else if in[i] == "n" && in[j] == "sw" {
			return true
		} else if in[i] == "ne" && in[j] == "nw" {
			return true
		} else if in[i] == "ne" && in[j] == "s" {
			return true
		} else if in[i] == "nw" && in[j] == "s" {
			return true
		} else if in[i] == "se" && in[j] == "sw" {
			return true
		}
	}
	return false
}

func canReduceBack(in []string, j int) bool {
	for i := 0; i < j; i++ {
		if (in[i] == "n" && in[j] == "s") || (in[i] == "ne" && in[j] == "sw") || (in[i] == "nw" && in[j] == "se") {
			return true
		} else if in[i] == "n" && in[j] == "se" {
			return true
		} else if in[i] == "n" && in[j] == "sw" {
			return true
		} else if in[i] == "ne" && in[j] == "nw" {
			return true
		} else if in[i] == "ne" && in[j] == "s" {
			return true
		} else if in[i] == "nw" && in[j] == "s" {
			return true
		} else if in[i] == "se" && in[j] == "sw" {
			return true
		}
	}
	return false
}

func merge(in []string) int {
	sort.Sort(sort.StringSlice(in))

	compact := make([]string, 0, len(in))
	i := 0
	j := len(in) - 1

	for i < len(in)-1 {
		if canReduceFront(in, i) {
			break
		}
		compact = append(compact, in[i])
		i++
	}

	for j > i {
		if canReduceBack(in, j) {
			break
		}
		compact = append(compact, in[j])
		j--
	}

	for i <= j {
		if (in[i] == "n" && in[j] == "s") || (in[i] == "ne" && in[j] == "sw") || (in[i] == "nw" && in[j] == "se") {
			// CANCEL OUT
			i++
			j--
		} else if in[i] == "n" && in[j] == "se" {
			compact = append(compact, "ne")
			i++
			j--
		} else if in[i] == "n" && in[j] == "sw" {
			compact = append(compact, "nw")
			i++
			j--
		} else if in[i] == "ne" && in[j] == "nw" {
			compact = append(compact, "n")
			i++
			j--
		} else if in[i] == "ne" && in[j] == "s" {
			compact = append(compact, "se")
			i++
			j--
		} else if in[i] == "nw" && in[j] == "s" {
			compact = append(compact, "sw")
			i++
			j--
		} else if in[i] == "se" && in[j] == "sw" {
			compact = append(compact, "s")
			i++
			j--
		} else {
			// CAN'T MERGE
			compact = append(compact, in[j])
			j--
		}
	}
	if len(compact) == len(in) {
		return len(compact)
	}

	return merge(compact)
}

func findmax(in string) int {
	steps := strings.Split(in, ",")

	max := 0
	for i := 0; i <= len(steps); i++ {
		l := merge(steps[:i])
		if l > max {
			max = l
		}
	}
	return max
}

func Solve() {
	if len(os.Args) < 2 {
		fmt.Println("Expected the input file name as commandline argument.")
		os.Exit(1)
	}

	b, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	o := merge(strings.Split(string(b), ","))
	fmt.Println("Part A", o)

	fmt.Println("Part B", findmax(string(b)))

}
