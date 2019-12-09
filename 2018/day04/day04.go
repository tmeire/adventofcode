package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type sleeper struct {
	minutes []int
	total   int
}

func collect(events []event) map[string]*sleeper {
	gID := ""
	sleepers := map[string]*sleeper{}
	var sleepStart time.Time

	for _, e := range events {
		pp := strings.Split(e.msg, " ")
		switch pp[0] {
		case "Guard":
			gID = pp[1]
			if _, ok := sleepers[gID]; !ok {
				sleepers[gID] = &sleeper{make([]int, 60), 0}
			}
		case "falls":
			sleepStart = e.t
		case "wakes":
			m := sleepStart
			for m.Before(e.t) {
				sleepers[gID].minutes[m.Minute()]++
				sleepers[gID].total++
				m = m.Add(time.Minute)
			}
		default:
			panic("Unknown event")
		}
	}
	return sleepers
}

func partA(sleepers map[string]*sleeper) int {
	var gID string
	maxMinutes := -1
	for id, sleeper := range sleepers {
		if sleeper.total > maxMinutes {
			maxMinutes = sleeper.total
			gID = id
		}
	}

	numID, err := strconv.Atoi(gID[1:])
	if err != nil {
		panic(err)
	}

	maxMinute := 0
	maxMinuteVal := 0
	for idx, e := range sleepers[gID].minutes {
		if e > maxMinuteVal {
			maxMinuteVal = e
			maxMinute = idx
		}
	}

	return numID * maxMinute
}

func partB(sleepers map[string]*sleeper) int {
	var maxgID string
	maxMinute := -1
	maxC := 0
	for id, sleeper := range sleepers {
		for minute, c := range sleeper.minutes {
			if c > maxC {
				maxC = c
				maxMinute = minute
				maxgID = id
			}
		}0
	}

	numID, err := strconv.Atoi(maxgID[1:])
	if err != nil {
		panic(err)
	}

	return numID * maxMinute
}

type event struct {
	t   time.Time
	msg string
}

var eventRegex = regexp.MustCompile(`\[([0-9\:\- ]+)\] (.+)`)

func parseEvent(s string) event {
	m := eventRegex.FindAllStringSubmatch(s, -1)
	if len(m) != 1 && len(m[0]) != 3 {
		panic(fmt.Sprintf("%#v\n", m))
	}

	t, err := time.Parse("2006-01-02 15:04", m[0][1])
	if err != nil {
		panic(err)
	}

	return event{t, m[0][2]}
}

func main() {
	f, err := os.Open("input-sorted.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	events := []event{}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		events = append(events, parseEvent(scanner.Text()))
	}

	sleepers := collect(events)

	fmt.Println(partA(sleepers))
	fmt.Println(partB(sleepers))
}
