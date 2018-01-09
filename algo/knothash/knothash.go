package knothash

var HEXCHAR = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}

type List []byte

func (l List) reverse(pos, length int) {
	for i := 0; i < length/2; i++ {
		a := (pos + i) % len(l)
		b := (pos + length - 1 - i) % len(l)

		l[a], l[b] = l[b], l[a]
	}
}

func (l List) Sparse(lengths []int, rounds int) List {
	pos := 0
	skip := 0

	for i := 0; i < rounds; i++ {
		for _, length := range lengths {
			l.reverse(pos, length)
			pos += length + skip
			skip++
		}
	}
	return l
}

func (l List) Dense() List {
	d := make(List, 16)
	idx := 0
	for i := 0; i < 16; i++ {
		var b byte
		for j := 0; j < 16; j++ {
			b ^= l[idx]
			idx++
		}
		d[i] = b
	}
	return d
}

func (l List) Hex() string {
	b := make([]byte, 32)
	for i, v := range l {
		b[2*i] = HEXCHAR[v/16]
		b[2*i+1] = HEXCHAR[v%16]
	}
	return string(b)
}

func Hash(input string) string {
	l := make(List, 256)
	for i := 0; i < 256; i++ {
		l[i] = byte(i)
	}

	lengths := make([]int, len(input), len(input)+5)
	lengths = append(lengths, 17, 31, 73, 47, 23)

	for i, b := range input {
		lengths[i] = int(b)
	}

	return l.Sparse(lengths, 64).Dense().Hex()
}
