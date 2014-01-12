package dfs

import (
	"fmt"
	"github.com/anupamk/common-utilz/graph"
)

func Example_001_DFSSearch() {
	graph_fname := "../data/graph-003.data"

	g, _ := graph.LoadGraphFromFile(graph_fname)
	dfs_g := New(g, 0)

	fmt.Println(dfs_g.Visited(0))

	// Output:
	// true
}
