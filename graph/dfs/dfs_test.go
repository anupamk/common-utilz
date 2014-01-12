package dfs

import (
	"github.com/anupamk/common-utilz/graph"
	"testing"
)

var graph_dfs_data = []struct {
	fname              string
	source_vertex      int32
	connected_vertices []int32
}{
	{"../data/graph-001.data", 0, []int32{0, 1, 2, 3, 4, 5, 6}},
	{"../data/graph-001.data", 9, []int32{9, 10, 11, 12}},
	{"../data/graph-001.data", 7, []int32{7, 8}},
	{"../data/graph-003.data", 0, []int32{0, 1, 2, 3, 4, 5}},
}

//
// this function returns true if two int32 slices are equal i.e. for
// all i, x[i] == y[i] and 0 <= i < len(x)
//
func cmp_int32_slice(x, y *[]int32) bool {
	if len(*x) != len(*y) {
		return false
	}

	for i, xv := range *x {
		if xv != (*y)[i] {
			return false
		}
	}

	return true
}

func TestDFSSearch(t *testing.T) {
	for _, graph_data := range graph_dfs_data {
		g, _ := graph.LoadGraphFromFile(graph_data.fname)
		dfs_g := New(g, graph_data.source_vertex)
		cv := dfs_g.ConnectedVertices()

		if cmp_int32_slice(&graph_data.connected_vertices, &cv) != true {
			t.Logf("expected: %v, got: %v\n", graph_data.connected_vertices, cv)
			t.Fail()
		}
	}
}

// benchmark

//
// pretty rudimentary actually, for all the vertices in a graph,
// create dfs with the given vertex as source.
//
func BenchmarkDepthFirstSearch(bench *testing.B) {
	graph_fname := "../data/graph-003.data"
	g, _ := graph.LoadGraphFromFile(graph_fname)

	for i := 0; i < bench.N; i++ {
		for v := int32(0); v < g.V(); v++ {
			tmp := New(g, v)
			if tmp.Count() == 0 {
				panic("Ooops")
			}
		}
	}
}
