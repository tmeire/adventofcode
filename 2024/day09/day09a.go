package day09

import (
	"fmt"
	"os"
)

func explode(diskmap []byte) []int {
	var blocks []int
	isFile := true
	index := 0
	for _, e := range diskmap {
		l := int(e - '0')
		v := -1
		if isFile {
			v = index
			index += 1
		}
		for i := 0; i < l; i++ {
			blocks = append(blocks, v)
		}
		isFile = !isFile
	}
	return blocks
}

func compress(blocks []int) []int {
	j := len(blocks) - 1
	for i := 0; i < len(blocks); i++ {
		if blocks[i] != -1 {
			continue
		}
		for blocks[j] == -1 {
			j--
		}
		if i >= j {
			break
		}
		blocks[i] = blocks[j]
		blocks[j] = -1
	}
	return blocks
}

func checksum(blocks []int) int {
	chksm := 0
	for i, b := range blocks {
		if b != -1 {
			chksm += i * b
		} else {
			break
		}
	}
	return chksm
}

type region struct {
	n, p   *region
	length int
	fileID int
}

func (r *region) printAll() {
	fmt.Printf("(%d, %d, %p, %p)\n", r.fileID, r.length, r.p, r.n)
	if r.n != nil {
		r.n.printAll()
	}
}

func regions(diskmap []byte) (*region, *region) {
	var head, tail *region

	isFile := true
	index := 0
	for _, e := range diskmap {
		v := -1
		if isFile {
			v = index
			index += 1
		}

		r := &region{
			p:      tail,
			length: int(e - '0'),
			fileID: v,
		}
		if tail != nil {
			tail.n = r
			tail = r
		} else {
			head = r
			tail = r
		}
		isFile = !isFile
	}
	return head, tail
}

func compressRegion(head, tail *region) *region {
	reg := tail
	for reg != head {
		if reg.fileID == -1 {
			reg = reg.p
			continue
		}
		r := head
		for (r.fileID != -1 || r.length < reg.length) && r != reg {
			r = r.n
		}
		if r == reg {
			// skip this block, can't move
			reg = reg.p
			continue
		}
		if r.length > reg.length {
			r.n = &region{
				p:      r,
				n:      r.n,
				fileID: -1,
				length: r.length - reg.length,
			}
			r.n.n.p = r.n
		}
		r.length = reg.length
		r.fileID = reg.fileID
		reg.fileID = -1
		reg = reg.p
		//printRegions(head)
	}
	return head
}

func checksumRegions(head *region) int {
	chksm := 0
	r := head
	index := 0
	for r != nil {
		if r.fileID == -1 {
			index += r.length
		} else {
			for i := 0; i < r.length; i++ {
				chksm += index * r.fileID
				index++
			}
		}
		r = r.n
	}
	return chksm
}

func printRegions(head *region) {
	r := head
	index := 0
	for r != nil {
		if r.fileID == -1 {
			for i := 0; i < r.length; i++ {
				print(".")
				index++
			}
		} else {
			for i := 0; i < r.length; i++ {
				print(r.fileID)
				index++
			}
		}
		r = r.n
	}
	println()
}

func Solve() {
	diskmap, err := os.ReadFile("./2024/day09/input.txt")
	if err != nil {
		panic(err)
	}

	blocks := explode(diskmap)
	compressed := compress(blocks)
	chksm := checksum(compressed)
	println(chksm)

	first, last := regions(diskmap)
	first = compressRegion(first, last)
	chksmRegions := checksumRegions(first)
	println(chksmRegions)
}
