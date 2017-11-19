package main

import (
	"fmt"
	"io/ioutil"
)

type Length struct {
	count int
}

func (l *Length) Write(b []byte) {
	l.count += len(b)
}

func (l *Length) WriteByte(b byte) {
	l.count += 1
}

func (l *Length) WriteString(b string) {
	l.count += len(b)
}

func (l *Length) Add(i int) {
	l.count += i
}

func (l *Length) Len() int {
	return l.count
}

type state interface {
	process(b byte) state
	Len() int
}

type plain struct {
	*Length
}

func (p *plain) process(b byte) state {
	switch b {
	case '(':
		return &counter{p.Length, "(", 0}
	default:
		p.WriteByte(b)
		return p
	}
}

type counter struct {
	*Length
	tmp   string
	count int
}

func (c *counter) add(i int, s string) state {
	c.count = c.count*10 + i
	c.tmp += s
	return c
}

func (c *counter) process(b byte) state {
	switch b {
	case '0':
		return c.add(0, "0")
	case '1':
		return c.add(1, "1")
	case '2':
		return c.add(2, "2")
	case '3':
		return c.add(3, "3")
	case '4':
		return c.add(4, "4")
	case '5':
		return c.add(5, "5")
	case '6':
		return c.add(6, "6")
	case '7':
		return c.add(7, "7")
	case '8':
		return c.add(8, "8")
	case '9':
		return c.add(9, "9")
	case 'x':
		// Done with the count, onto the number of repeats
		return &repeater{c.Length, c.tmp + "x", c.count, 0}
	case '(':
		// The previous ( was not the start of a marker, maybe this one is
		c.WriteString(c.tmp)
		return &counter{c.Length, "(", 0}
	default:
		// Oops, nope, not a valid counter character, back to plain
		c.WriteString(c.tmp)
		c.WriteByte(b)
		return &plain{c.Length}
	}
}

type repeater struct {
	*Length
	tmp    string
	count  int
	repeat int
}

func (r *repeater) add(i int, s string) state {
	r.repeat = r.repeat*10 + i
	r.tmp += s
	return r
}

func (r *repeater) process(b byte) state {
	switch b {
	case '0':
		return r.add(0, "0")
	case '1':
		return r.add(1, "1")
	case '2':
		return r.add(2, "2")
	case '3':
		return r.add(3, "3")
	case '4':
		return r.add(4, "4")
	case '5':
		return r.add(5, "5")
	case '6':
		return r.add(6, "6")
	case '7':
		return r.add(7, "7")
	case '8':
		return r.add(8, "8")
	case '9':
		return r.add(9, "9")
	case ')':
		// Done with the repeat, time to copy next characters
		return &copier{r.Length, make([]byte, r.count), r.count, r.repeat, 0}
	case '(':
		// The previous ( was not the start of a marker, maybe this one is
		r.WriteString(r.tmp)
		return &counter{r.Length, "(", 0}
	default:
		// Oops, nope, not a valid counter character, back to plain
		r.WriteString(r.tmp)
		r.WriteByte(b)
		return &plain{r.Length}
	}
}

type copier struct {
	*Length
	tmp    []byte
	count  int
	repeat int
	seen   int
}

func (c *copier) process(b byte) state {
	c.tmp[c.seen] = b
	c.seen += 1

	if c.seen < c.count {
		// We haven't seen all characters yet
		return c
	}

	// Decompress the data within
	d := decompress(c.tmp)

	// Add the repeat times the decompressed length
	c.Add(d * c.repeat)

	// And back to plain
	return &plain{c.Length}
}

func decompressString(ss string) int {
	return decompress([]byte(ss))
}

func decompress(ss []byte) int {
	p := state(&plain{&Length{0}})

	for _, s := range ss {
		p = p.process(s)
	}
	return p.Len()
}

func main() {
	examples := []string{"advent", "A(1x5)BC", "(3x3)XYZ", "A(2x2)BCD(2x2)EFG", "(6x1)(1x3)A", "X(8x2)(3x3)ABCY"}
	for _, e := range examples {
		fmt.Println(decompressString(e))
	}

	bs, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic("No such file")
	}

	d := decompress(bs)
	fmt.Println("Decompressed length:", d)
}

/*
states:

PLAIN --- (      --> MARKER_OPEN
	  --- a-z0-9 --> PLAIN

MARKER_OPEN --- 0-9 --> MARKER_COUNT
            --- a-z --> PLAIN

MARKER_COUNT --- 0-9 --> MARKER_COUNT
             --- x   --> MARKER_REPEAT
             --- (   --> MARKER_OPEN
             --- a-z --> PLAIN

MARKER_REPEAT --- 0-9 --> MARKER_REPEAT
              --- (   --> MARKER_OPEN
              --- )   --> MARKER_CLOSED
              --- a-z --> PLAIN

*/
