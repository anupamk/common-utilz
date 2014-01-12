//
// provides an implementation of depth-first-search algorithm for
// graphs
//
package dfs

import (
	"github.com/anupamk/common-utilz/graph"
)

type DepthFirstSearch struct {
	visited []bool
	count   int32
}

func New(G *graph.Graph, source int32) (dfs *DepthFirstSearch) {
	dfs = &DepthFirstSearch{
		visited: make([]bool, G.V()),
		count:   0,
	}

	dfs.run_dfs(G, source)
	return
}

func (dfs *DepthFirstSearch) Visited(v int32) bool { return dfs.visited[v] }
func (dfs *DepthFirstSearch) Count() int32         { return dfs.count }

//
// returns the list of vertices that are connected to the source
// vertex
//
func (dfs *DepthFirstSearch) ConnectedVertices() (vertex_list []int32) {
	vertex_list = make([]int32, dfs.count)
	for i, j := 0, 0; i < len(dfs.visited); i++ {
		if dfs.visited[i] == true {
			vertex_list[j] = int32(i)
			j = j + 1
		}
	}

	return
}

//
// private unexported stuff
//

// implements the canonical recursive dfs procedure
func (dfs *DepthFirstSearch) run_dfs(G *graph.Graph, source int32) {
	dfs.visited[source] = true
	dfs.count += 1

	for _, w := range G.Adj(source) {
		if dfs.visited[w] == false {
			dfs.run_dfs(G, w)
		}
	}

	return
}
