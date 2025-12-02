package day02

import (
	"fmt"
	"strconv"
	"strings"
)

const example = "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"

const input = "269194394-269335492,62371645-62509655,958929250-958994165,1336-3155,723925-849457,4416182-4470506,1775759815-1775887457,44422705-44477011,7612653647-7612728309,235784-396818,751-1236,20-36,4-14,9971242-10046246,8796089-8943190,34266-99164,2931385381-2931511480,277-640,894249-1083306,648255-713763,19167863-19202443,62-92,534463-598755,93-196,2276873-2559254,123712-212673,31261442-31408224,421375-503954,8383763979-8383947043,17194-32288,941928989-941964298,3416-9716"

type Range struct {
	low  int
	high int
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err.Error())
	}
	return i
}

func ranges(input string) []Range {
	var r []Range
	for _, s := range strings.Split(input, ",") {
		ss := strings.Split(s, "-")
		r = append(r, Range{atoi(ss[0]), atoi(ss[1])})
	}
	return r
}

func Solve() {
	r := ranges(input)

	sum := 0
	// for each range
	for _, rr := range r {
		// for each id in the range
		for i := rr.low; i <= rr.high; i++ {
			id := fmt.Sprintf("%d", i)
			// part 1
			//if len(id)%2 == 0 {
			//	if id[:(len(id)/2)+1] == id[len(id)/2:] {
			//		sum += i
			//	}
			//}
			// for each (subsequently shorter) repeating substring length
			// start with the long substrings first as they require fewer iterations to check
			for n := len(id) / 2; n > 0; n-- {
				// skip the string manipulation if it's not divisible by n anyway
				if len(id)%n != 0 {
					continue
				}
				match := true
				for j := 1; j < len(id)/n; j++ {
					if id[:n] != id[j*n:j*n+n] {
						match = false
						break
					}
				}
				// if it matches, add the id to the sum and stop iterating
				if match {
					sum += i
					break
				}
			}
		}
	}
	println(sum)
}
