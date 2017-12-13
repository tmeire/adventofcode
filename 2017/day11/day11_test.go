package day11

import "testing"

func TestPartA(t *testing.T) {
	fixtures := []struct {
		input  string
		output int
	}{
		{"ne,ne,ne", 3},
		{"ne,ne,sw,sw", 0},
		{"ne,ne,s,s", 2},
		{"se,sw,se,sw,sw", 3},
		{"nw,s,se", 1},
		{"se,ne,ne,n,n,n,n,n,n,nw,nw,nw,se", 8},
		{"nw,s,sw", 2},
	}

	for i, f := range fixtures {
		out := collapse(f.input)
		if out != f.output {
			t.Errorf("Case %d failed: got %d, wanted %d", i, out, f.output)
		}
	}
}

func TeestPartB(t *testing.T) {
	fixtures := []struct {
		input  string
		output int
	}{
		{"ne,ne,ne", 3},
		{"ne,ne,sw,sw", 2},
		{"ne,sw,ne,sw", 1},
		{"ne,ne,s,s", 2},
		{"se,sw,se,sw,sw", 3},
		{"se,sw,se,sw", 2},
		{"se,ne,ne,n,n,n,n,n,n", 8},
		{"se,ne,ne,n,n,n,n,n,n,nw", 8},
		{"se,ne,ne,n,n,n,n,n,n,nw,nw,nw", 8},
		{"se,ne,ne,n,n,n,n,n,n,nw,nw,nw,se", 8},
	}

	for i, f := range fixtures {
		out := findmax(f.input)
		if out != f.output {
			t.Errorf("Case %d failed: got %d, wanted %d", i, out, f.output)
		}
	}
}

func BenchmarkFindmax(b *testing.B) {
	input := "se,ne,ne,n,n,n,n,n,n,nw,nw,nw,se,nw,nw,sw,se,nw,sw,nw,se,sw,s,sw,s,s,sw,s,sw,sw,ne,sw,s,sw,s,sw,nw,s,s,s,s,s,s,sw,s,s,n,nw,s,s,se,se,se,s,nw,se,s,se,nw,se,se,n,se,se,se,se,se,se,se,ne,sw,nw,ne,n,se,sw,nw,ne,se,se,ne,sw,se,ne,ne,se,n,ne,ne,ne,ne,ne,ne,sw,ne,ne,ne,ne,ne,ne,sw,ne,nw,s,se,ne,ne,ne,nw,ne,sw,ne,ne,ne,ne,ne,ne,nw,sw,sw,ne,ne,ne,n,s,ne,n,s,n,ne,n,n,n,n,n,n,n,ne,n,n,s,s,ne,n,n,sw,n,sw,n,n,se,n,n,n,n,n,n,n,n,s,sw,ne,nw,nw,n,n,se,se,s,n,nw,n,n,n,nw,n,nw,n,n,nw,se,ne,nw,nw,n,nw,n,nw,s,n,nw,n,ne,n,n,nw,n,nw,se,se,nw,nw,nw,nw,n,nw,ne,ne,n,s,nw,nw,nw,nw,n,nw,sw,se,nw,nw,nw,sw,nw,s,nw,nw,nw,nw,nw,se,n,ne,nw,ne,sw,sw,nw,sw,s,nw,nw,ne,nw,sw,nw,sw,nw,n,n,sw,nw,nw,n,ne,s,ne,sw,ne,nw,nw,sw,nw,nw,s,nw,nw,nw,nw,nw,nw,sw,sw,nw,nw,se,nw,sw,nw,ne,sw,sw,nw,nw,nw,nw,nw,s,sw,nw,sw,sw,nw,nw,sw,sw,sw,sw,sw,sw,sw,ne,sw,nw,sw,sw,nw,sw,ne,ne,sw,se,se,sw,sw,sw,sw,sw,n,n,sw,nw,n,sw,sw,n,se,sw,sw,sw,sw,sw,sw,sw,sw,sw,se,sw,ne,sw,n,sw,sw,s,nw,s,s,sw,sw,sw,s,sw,sw,sw,ne,nw,s,sw,s,sw,s,ne,s,ne,sw,sw,sw,s,sw,s,sw,s,s,sw,sw,sw,n,sw,s,sw,sw,s,sw,sw,sw,sw,s,sw,sw,sw,se,sw,ne,sw,sw,se,sw,sw,sw,sw,s,s,s,s,sw,s,s,sw,sw,s,nw,n,s,s,sw,s,s,n,se,s,sw,sw,se,se,s,s,sw,sw,nw,s,nw,s,s,s,s,s,s,sw,sw,s,s,s,n,sw,s,sw,s,s,s,s,s,n,s,se,s,ne,ne,s,s,ne,s,s,sw,s,s,s,s,s,s,se,s,n,s,s,s,s,s,se,s,se,s,se,ne,s,s,s,sw,se,se,se,s,s,se,nw,s,s,s,sw,se,s,s,se,s,ne,s,s,se,s,s,s,s,s,s,s,se,se,nw,s,s,se,n,s,se,ne,s,s,s,ne,s,se,nw,s,s,s,s,n,ne,s,se,se,s,s,ne,s,sw,s,s,s,se,s,s,s,se,ne,s,s,s,s,se,s,s,s,se,s,s,s,se,s,s,nw,s,s,nw,s,se,sw,s,s,se,s,s,se,se,s,se,se,s,s,se,sw,se,se,se,se,se,s,s,s,s,se,se,ne,se,s,ne,nw,se,se,s,se,nw,se,se,n,se,s,n,se,n,s,se,se,se,sw,ne,se,se,s,n,s,se,se,se,s,se,s,se,n,se,se,sw,se,se,s,nw,sw,se,nw,se,se,se,se,s,se,se,se,se,se,se,se,se,se,se,n,s,se,se,se,se,ne,nw,se,se,nw,se,se,n,s,se,se,se,s,se,se,se,se,se,se,ne,se,se,ne,se,se,nw,se,ne,se,se,se,se,s,se,ne,se,se,se,ne,se,nw,ne,se,s,se,se,se,se,se,nw,se,se,sw,sw,se,se,se,sw,se,ne,s,se,nw,ne,se,ne,nw,se,se,se,se,se,se,sw,se,ne,se,se,se,se,se,se,se,ne,se,se,se,se,se,s,se,ne,sw,se,s,se,ne,s,se,se,se,se,sw,sw,se,se,ne,ne,se,se,nw,ne,ne,ne,se,se,se,se,se,s,ne,se,ne,ne,se,ne,ne,s,n,ne,ne,se,ne,ne,se,se,se,se,nw,se,se,ne,se,ne,se,se,s,se,se,ne,se,sw,se,se,se,ne,ne,se,se,sw,se,ne,ne,se,se,n,ne,ne,ne,se,ne,ne,ne,se,ne,ne,ne,ne,ne,ne,se,se,se,ne,se,ne,ne,ne,ne,nw,ne,nw,se,ne,ne,ne,se,ne,ne,ne,se,ne,ne,se,n,n,ne,nw,se,se,ne,ne,se,ne,n,ne,ne,ne,se,se,se,nw,nw,se,ne,ne,ne,ne,nw,ne,ne,ne,ne,ne,se,ne,ne,ne,ne,ne,ne,ne,ne,se,ne,s,sw,ne,ne,n,se,ne,ne,se,ne,ne,ne,ne,ne,ne,se,ne,sw,ne,ne,ne,s,ne,ne,ne,ne,ne,ne,ne,ne,sw,ne,ne,ne,ne,ne,nw,sw,se,nw,ne,ne,ne,ne,se,ne,s,ne,ne,ne,ne,ne,s,ne,ne,ne,ne,ne,n,ne,n,ne,ne,ne,nw,se,ne,ne,ne,ne,n,nw,ne,ne,ne,ne,sw,ne,n,n,ne,ne,ne,ne,ne,ne,ne,ne,ne,ne,ne,n,ne,n,ne,ne,se,n,n,nw,ne,ne,ne,ne,ne,ne,ne,ne,ne,ne,se,ne,n,n,ne,ne,se,ne,n,s,ne,ne,ne,s,n,ne,s,n,ne,s,n,n,s,n,ne,ne,ne,nw,n,n,n,ne,nw,s,ne,n,ne,ne,n,ne,ne,n,n,n,ne,ne,n,ne,n,ne,ne,sw,ne,ne,ne,ne,ne,n,n,ne,ne,se,ne,se,nw,ne,sw,ne,ne,s,ne,n,ne,n,nw,s,s,ne,ne,nw,ne,se,se,ne,ne,se,ne,n,n,n,sw,ne,s,n,ne,n,n,n,ne,sw,n,ne,ne,ne,n,ne,ne,n,ne,ne,nw,n,n,nw,ne,n,n,se,ne,n,ne,ne,n,n,ne,n,se,ne,nw,ne,n,sw,se,ne,n,n,se,nw,sw,n,ne,s,n,n,ne,n,ne,se,ne,n,n,ne,n,n,n,n,n,ne,ne,sw,ne,se,ne,ne,n,n,ne,ne,ne,n,ne,ne,ne,n,se,n,ne,se,sw,sw,n,ne,n,n,n,n,n,n,ne,n,n,ne,ne,n,n,sw,nw,nw,ne,n,ne,n,se,se,ne,n,ne,n,ne,ne,ne,ne,n,n,n,n,ne,n,n,n,ne,n,ne,n,n,n,se,ne,ne,nw,n,ne,ne,nw,n,ne,n,n,n,n,ne,n,n,n,ne,n,n,n,ne,nw,n,ne,n,n,n,n,sw,ne,n,sw,n,n,ne,n,n,ne,n,ne,nw,n,n,n,ne,n,n,n,sw,n,n,n,n,n,n,n,n,n,nw,n,n,n,sw,n,n,n,n,n,n,n,n,se,ne,n,n,n,n,nw,n,ne,n,se,se,n,n,s,ne,n,n,n,n,n,sw,s,n,n,n,n,n,n,ne,sw,n,s,n,n,n,nw,n,n,n,n,n,n,n,ne,n,n,n,n,n,se,n,n,n,nw,n,n,n,n,sw,n,n,n,n,se,n,n,nw,n,n,n,n,n,n,s,n,n,n,n,n,se,n,nw,n,n,n,n,n,n,nw,n,n,n,n,n,n,ne,nw,n,nw,n,n,n,n,n,n,n,n,se,nw,n,n,n,n,nw,n,nw,n,se,n,n,n,n,s,ne,sw,n,n,se,se,n,n,n,sw,n,n,n,nw,n,n,ne,n,nw,n,nw,n,n,nw,n,n,n,n,n,n,nw,n,se,nw,s,nw,n,se,se,n,n,n,n,n,n,ne,ne,n,n,n,n,n,n,n,se,n,n,n,sw,n,sw,n,nw,nw,s,nw,n,nw,n,n,n,nw,n,n,s,n,n,n,nw,nw,se,n,nw,nw,n,nw,n,nw,n,n,se,nw,n,n,n,n,n,n,n,ne,nw,n,nw,nw,n,nw,ne,n,n,n,n,s,se,s,n,n,n,n,n,n,se,n,s,n,nw,nw,nw,nw,n,n,n,sw,se,s,n,nw,n,n,nw,n,n,n,n,n,nw,n,s,n,se,n,ne,n,nw,nw,nw,nw,n,n,n,n,n,se,nw,sw,nw,n,nw,n,n,n,n,nw,nw,se,nw,n,nw,n,nw,nw,sw,n,n,n,n,nw,n,n,nw,nw,sw,n,s,n,nw,n,n,n,nw,n,nw,se,nw,n,nw,ne,ne,s,nw,se,nw,nw,n,nw,nw,n,nw,nw,n,n,ne,s,nw,nw,nw,n,nw,n,n,nw,n,nw,sw,se,se,nw,n,ne,n,n,se,nw,n,s,ne,sw,nw,n,nw,nw,n,sw,nw,se,n,se,n,n,nw,ne,nw,nw,nw,nw,nw,s,nw,nw,nw,nw,nw,ne,n,nw,nw,nw,nw,n,n,n,n,n,sw,nw,n,n,n,nw,nw,s,nw,n,se,nw,n,se,n,n,nw,nw,nw,nw,n,nw,nw,nw,nw,nw,nw,n,nw,n,s,nw,nw,nw,nw,n,nw,n,se,nw,n,ne,nw,n,nw,n,se,n,n,nw,nw,nw,nw,nw,sw,nw,nw,sw,nw,nw,nw,ne,nw,n,s,s,nw,ne,nw,n,nw,n,ne,nw,n,nw,nw,ne,nw,n,nw,nw,se,nw,n,nw,n,sw,se,n,n,nw,ne,nw,nw,n,nw,nw,nw,nw,nw,se,se,se,ne,nw,n,nw,nw,s,nw,nw,nw,se,n,nw,nw,se,n,nw,nw,nw,n,nw,sw,nw,se,n,nw,nw,nw,n,nw,ne,nw,n,nw,nw,sw,nw,nw,n,nw,nw,n,nw,ne,sw,nw,n,n,nw,sw,nw,se,nw,nw,nw,s,nw,nw,n,n,sw,n,s,ne,sw,s,nw,nw,nw,nw,nw,nw,nw,nw,nw,nw,nw,nw,nw,nw,nw,n,n,nw,nw,nw,ne,n,nw,nw,s,nw,ne,s,ne,nw,nw,nw,nw,s,nw,nw,nw,nw,se,s,nw,nw,sw,nw,nw,nw,nw,nw,nw,nw,nw,n,s,s,nw,ne,nw,se,nw,se,nw,nw,nw,ne,se,nw,ne,nw,nw,nw,nw,n,se,nw,nw,nw,ne,nw,nw,s,nw,nw,sw,ne,s,nw,nw,nw,sw,n,nw,nw,nw,nw,sw,s,se,nw,nw,nw,se,ne,nw,nw,n,nw,sw,nw,sw,nw,nw,n,nw,nw,nw,n,nw,nw,nw,nw,nw,nw,sw,nw,sw,nw,nw,nw,n,nw,s,sw,nw,sw,nw,nw,nw,nw,sw,nw,nw,nw,nw,n,nw,sw,nw,nw,sw,nw,s,nw,s,nw,sw,sw,nw,nw,se,nw,nw,nw,sw,nw,nw,sw,nw,nw,n,s,ne,nw,nw,nw,nw,se,nw,n,nw,nw,sw,nw,nw,nw,nw,nw,nw,nw,sw,nw,nw,nw,nw,sw,nw,se,nw,s,nw,nw,sw,nw,nw,sw,sw,nw,nw,nw,nw,nw,nw,nw,nw,s,nw,nw,nw,nw,s,nw,nw,nw,sw,nw,nw,s,nw,nw,nw,nw,nw,nw,nw,nw,nw,sw,nw,nw,ne,nw,s,sw,nw,nw,nw,sw,nw,ne,se,sw,nw,nw,nw,se,nw,sw,nw,nw,sw,nw,sw,sw,sw,nw,nw,sw,nw,nw,nw,sw,ne,nw,nw,n,nw,nw,sw,nw,sw,nw,nw,sw,sw,nw,ne,sw,nw,sw,se,sw,n,sw,n,ne,nw,nw,n,s,nw,nw,nw,ne,nw,nw,nw,nw,s,nw,s,nw,nw,sw,sw,nw,nw,nw,nw,nw,nw,nw,nw,n,nw,nw,se,nw,n,nw,n,ne,n,nw,sw,sw,sw,nw,nw,ne,nw,nw,nw,ne,n,nw,sw,sw,nw,sw,nw,sw,nw,sw,n,sw,nw,nw,sw,nw,nw,ne,nw,nw,nw,nw,sw,sw,nw,nw,sw,nw,sw,nw,nw,s,ne,s,sw,nw,nw,nw,sw,sw,nw,nw,sw,nw,s,sw,sw,nw,se,nw,nw,sw,nw,nw,sw,sw,nw,sw,sw,nw,s,nw,nw,nw,sw,sw,nw,nw,nw,se,sw,s,nw,nw,nw,nw,nw,nw,sw,n,s,nw,sw,nw,se,n,sw,ne,se,sw,nw,nw,ne,sw,nw,sw,sw,nw,ne,ne,sw,sw,nw,se,sw,nw,nw,sw,sw,sw,sw,se,sw,n,nw,nw,sw,sw,sw,sw,sw,nw,sw,nw,nw,sw,sw,nw,n,sw,sw,sw,nw,nw,nw,nw,sw,s,sw,nw,sw,sw,nw,sw,nw,n,nw,s,nw,sw,nw,se,sw,nw,ne,ne,sw,n,nw,nw,nw,nw,nw,nw,nw,sw,sw,sw,nw,nw,nw,nw,nw,nw,sw,nw,nw,nw,nw,sw,nw,sw,se,sw,nw,sw,nw,ne,nw,sw,sw,sw,s,nw,sw,nw,nw,nw,nw,sw,sw,s,n,s,sw,sw,se,sw,nw,ne,nw,sw,nw,sw,sw,nw,nw,s,sw,sw,ne,sw,se,n,sw,sw,nw,nw,sw,se,n,nw,sw,nw,sw,n,nw,n,sw,sw,sw,sw,nw,s,sw,nw,sw,nw,sw,sw,s,nw,sw,sw,sw,sw,sw,sw,se,sw,ne,nw,sw,se,sw,sw,nw,sw,sw,sw,se,n,sw,nw,sw,se,sw,se,s,nw,sw,nw,sw,sw,nw,n,s,nw,nw,sw,sw,nw,sw,sw,sw,n,nw,n,ne,se,s,se,s,sw,sw,se,n,sw,sw,sw,sw,sw,sw,sw,nw,sw,sw,sw,nw,sw,sw,sw,nw,sw,nw,sw,sw,sw,sw,sw,sw,sw,sw,sw,sw,sw,nw,sw,sw,nw,sw,nw,sw,nw,sw,sw,sw,s,nw,nw,n,sw,sw,nw,nw,sw,sw,ne,s,sw,nw,s,sw,nw,s,nw,se,sw,sw,sw,n,sw,n,sw,ne,n,sw,sw,ne,sw,sw,sw,se,sw,sw,sw,se,sw,sw,sw,sw,sw,sw,nw,sw,ne,ne,sw,sw,sw,nw,sw,sw,sw,sw,sw,ne,nw,sw,nw,sw,sw,se,sw,sw,sw,sw,ne,sw,sw,sw,sw,sw,sw,sw,se,sw,sw,nw,sw,sw,sw,nw,nw,sw,sw,se,nw,sw,n,sw,sw,sw,sw,sw,sw,sw,sw,s,sw,sw,se,sw,s,s,sw,n,sw,sw,sw,sw,sw,sw,sw,sw,n,sw,sw,sw,sw,sw,nw,nw,nw,s,se,sw,sw,sw,s,s,se,sw,sw,sw,se,sw,n,sw,sw,sw,sw,sw,sw,sw,sw,sw,sw,sw,sw,sw,sw,ne,sw,sw,sw,sw,sw,n,n,sw,sw,sw,sw,nw,n,sw,n,sw,sw,sw,sw,sw,sw,sw,nw,sw,sw,sw,ne,sw,sw,sw,ne,n,se,s,sw,sw,sw,sw,ne,sw,sw,ne,sw,sw,sw,sw,nw,sw,s,se,se,sw,sw,sw,s,sw,sw,sw,sw,sw,s,n,sw,sw,sw,sw,sw,sw,sw,sw,sw,nw,sw,sw,sw,sw,sw,se,sw,sw,sw,ne,s,sw,sw,sw,sw,sw,s,sw,sw,sw,ne,sw,sw,s,sw,sw,s,sw,sw,sw,sw,sw,sw,sw,sw,se,sw,sw,sw,sw,sw,sw,sw,nw,sw,sw,sw,sw,sw,sw,sw,s,sw,sw,nw,sw,sw,sw,ne,ne,sw,n,sw,sw,s,sw,sw,sw,sw,s,sw,sw,sw,s,sw,sw,sw,sw,nw,sw,sw,sw,ne,sw,sw,s,sw,sw,sw,s,sw,sw,sw,sw,nw,sw,sw,sw,sw,s,sw,sw,sw,s,sw,sw,sw,s,sw,n,nw,sw,nw,s,nw,sw,sw,sw,s,sw,sw,sw,s,s,ne,ne,sw,sw,sw,s,sw,sw,sw,sw,sw,nw,s,sw,sw,sw,sw,sw,sw,sw,s,sw,sw,sw,s,sw,n,sw,nw,sw,nw,sw,s,sw,sw,s,sw,n,sw,sw,s,sw,s,s,n,sw,sw,sw,sw,se,ne,s,s,sw,s,sw,sw,sw,ne,sw,ne,sw,se,sw,sw,s,sw,sw,nw,s,s,sw,ne,s,sw,sw,nw,n,s,ne,s,s,s,n,se,sw,s,s,sw,sw,s,ne,sw,n,s,sw,s,n,se,sw,sw,sw,sw,sw,sw,s,sw,s,nw,s,sw,s,sw,n,sw,s,sw,sw,s,sw,sw,ne,sw,se,sw,sw,nw,ne,sw,sw,s,ne,sw,sw,sw,sw,sw,sw,sw,ne,sw,ne,s,ne,sw,s,s,s,n,n,nw,sw,sw,se,se,sw,ne,n,n,sw,ne,sw,se,sw,sw,s,sw,sw,nw,sw,sw,sw,s,sw,s,sw,s,sw,s,sw,sw,s,sw,sw,s,sw,nw,s,sw,sw,s,sw,sw,sw,s,s,sw,s,sw,sw,n,s,sw,s,sw,s,sw,s,s,sw,sw,sw,sw,nw,s,sw,s,sw,sw,ne,sw,sw,s,s,nw,sw,sw,sw,sw,sw,sw,sw,sw,nw,sw,s,s,sw,sw,sw,s,sw,sw,nw,sw,sw,sw,sw,ne,sw,sw,sw,s,s,ne,s,sw,s,sw,sw,sw,sw,sw,sw,s,s,s,nw,se,se,ne,s,s,sw,ne,s,n,s,nw,sw,sw,s,sw,sw,n"
	for i := 0; i < b.N; i++ {
		findmax(input)
	}
}
