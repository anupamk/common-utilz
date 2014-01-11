package graph

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

var graph_data_files = [...]string{
	"data/graph-001.data",
	"data/graph-002.data",
}

//
// save a newly created graph, and see if the serialized-output
// matches the expected one
//
func Example_001_CreateAndDumpGraph() {
	tmp, _ := LoadGraphFromFile(graph_data_files[0])
	tmp_str := tmp.SerializeGraph()
	fmt.Println(tmp_str)

	// Output:
	// 13
	// 13
	// 0 5
	// 0 1
	// 0 2
	// 0 6
	// 3 4
	// 3 5
	// 4 6
	// 4 5
	// 7 8
	// 9 12
	// 9 10
	// 9 11
	// 11 12
}

func Example_002_CreateAndDumpGraph() {
	tmp, _ := LoadGraphFromFile(graph_data_files[1])
	tmp_str := tmp.SerializeGraph()
	fmt.Println(tmp_str)

	// Output:
	// 3
	// 3
	// 0 1
	// 0 2
	// 1 2
}

//
// create a new graph from serialized graph, and compare the two for
// equality.
//
func TestLoadGraph(t *testing.T) {
	for _, fname := range graph_data_files {
		g1, _ := LoadGraphFromFile(fname)
		g1_str := g1.SerializeGraph()

		g1_reader := strings.NewReader(g1_str)
		graph_reader := bufio.NewReader(g1_reader)

		g2, _ := LoadGraphFromReader(graph_reader)
		g2_str := g2.SerializeGraph()

		if g1_str != g2_str {
			t.Log("Error: Unequal Graphs")
			t.Fail()
		}
	}
}
