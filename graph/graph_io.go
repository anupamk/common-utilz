//
// commonly used functions for reading/writing graph structures
//
package graph

import (
	"bufio"
	"fmt"
	"github.com/anupamk/common-utilz/line_parser"
	"io"
	"os"
)

//
// this function is called to create a graph from it's serialized
// definition.
//
func LoadGraphFromReader(src *bufio.Reader) (new_graph *Graph, err error) {
	var V, E int32
	var edges [][2]int32

	V, E, edges, err = parse_graph_input(src)
	if err != nil && err != io.EOF {
		goto all_done
	}

	// create the graph, and setup the connections
	new_graph = New(V)
	for i := int32(0); i < E; i++ {
		v, w := edges[i][0], edges[i][1]

		// skip edges which are obviouzly bogus
		if V < v || v < 0 || V < w || w < 0 {
			fmt.Printf("skipping bogus connection: %d %d\n", v, w)
			continue
		}
		new_graph.AddEdge(v, w)
	}

all_done:
	return
}

//
// this is a convenience interface over LoadGraph(...) to create a
// graph from its serialized definition stored in a file identified by
// 'fname'
//
func LoadGraphFromFile(fname string) (g *Graph, err error) {
	var f *os.File

	if f, err = os.Open(fname); err != nil {
		return
	}
	defer f.Close()

	file_reader := bufio.NewReader(f)
	if g, err = LoadGraphFromReader(file_reader); err != nil {
		return
	}

	return
}

//
// dump the string representation of the graph, in the following
// format :
//     <line-001> V vertices, E edges
//     <line-002> vertex-1 : adj-list-of(vertex-1)
//     <line-003> vertex-2 : adj-list-of(vertex-2)
//     <line-004> vertex-3 : adj-list-of(vertex-3)
//
func (G *Graph) String() string {
	str := fmt.Sprintf("%d vertices, %d edges\n", G.v, G.e)

	for v, l := range G.adj {
		str += fmt.Sprintf("%d : ", v)

		for w := l.Front(); w != nil; w = w.Next() {
			str += fmt.Sprintf("%d ", w.Value.(int32))
		}

		str += fmt.Sprintf("\n")
	}

	return str
}

//
// this function emits the graph structure in a format suitable for
// subsequent loading via a LoadGraph(...) invokation
//
func (G *Graph) SerializeGraph() string {
	str := ""

	// vertex and edge count
	str += fmt.Sprintf("%d\n", G.V())
	str += fmt.Sprintf("%d\n", G.E())

	// vertex-specific adjacency-list dump
	for v, adj_list := range G.adj {
		for node := adj_list.Front(); node != nil; node = node.Next() {
			w := node.Value.(int32)

			// for undirected graphs, don't dump both v-w,
			// and w-v edges
			if int32(v) > w {
				continue
			}
			str += fmt.Sprintf("%d %d\n", v, w)
		}
	}

	return str
}

/*
 * private un-exported stuff
**/

//
// parse the input data file which is expected to be in the following
// format:
//     <line-001> number-of-vertices (V)
//     <line-002> number-of-edges (E)
//     <line-003> vertex-i vertex-j (edges[][2])
//     ....................
//     ....................
//     <line-N>   vertex-i vertex-j (edges[][2])
//
// commented-lines (starting with '#') / non-empty lines are ignored
//
func parse_graph_input(in *bufio.Reader) (V int32, E int32, Edges [][2]int32, err error) {
	var values []int32
	var comment_char byte = '#'

	if in == nil {
		return
	}

	// parse vertex-count
	values, err = line_parser.Int32sFromReader(in, comment_char)
	if err != nil {
		return
	}
	V = values[0]

	// parse edge-count
	values, err = line_parser.Int32sFromReader(in, comment_char)
	if err != nil {
		return
	}
	E = values[0]

	// parse edge-list
	Edges = make([][2]int32, E)
	for i := int32(0); i < E; i++ {
		values, err = line_parser.Int32sFromReader(in, comment_char)
		if err != nil {
			break
		}
		Edges[i][0], Edges[i][1] = values[0], values[1]
	}

	return
}
