//
// Copyright (c) 2014, Anupam Kapoor. All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:

// 1. Redistributions of source code must retain the above copyright notice, this
//    list of conditions and the following disclaimer.
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution.

// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR
// ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
// Author: anupam.kapoor@gmail.com (Anupam Kapoor)
//
//
// couple of commonly used algorithms for graphs are provided in this
// package.
//
// this file implements a cycle-detector which detects and returns a
// detected cycle in digraphs
//
package algorithms

import (
	"fmt"
	"github.com/anupamk/common-utilz/graph"
	"github.com/anupamk/common-utilz/stack"
)

type DirectedCycle struct {
	cycle []int32
}

func (dc *DirectedCycle) String() string {
	return fmt.Sprintf("%v\n", dc.cycle)
}

//
// this function return a DirectedCycle on a given digraph
//
func DigraphCyle(G graph.GraphOps) (dc *DirectedCycle) {
	var do_dfs func(graph.GraphOps, int32)

	visited := make([]bool, G.V())
	edge_to := make([]int32, G.V())
	vertex_stack := make([]bool, G.V())
	cycle_stack := stack.New()

	// the dfs procedure
	do_dfs = func(G graph.GraphOps, v int32) {
		vertex_stack[v] = true // v is on the path
		visited[v] = true

		for _, w := range G.Adj(v) {
			switch {
			case cycle_stack.Len() > 0:
				return

			case !visited[w]:
				edge_to[w] = v
				do_dfs(G, w)

			case vertex_stack[w] == true:
				// detected a cycle
				for x := v; x != w; x = edge_to[x] {
					cycle_stack.Push(x)
				}
				cycle_stack.Push(w)
				cycle_stack.Push(v)
			}
		}
		vertex_stack[v] = false // v is not on path anymore
	}

	// run on the whole graph
	for v := int32(0); v < G.V(); v++ {
		if !visited[v] {
			do_dfs(G, v)
		}
	}

	// nothing more to do
	if cycle_stack.Len() == 0 {
		return
	}

	// ok found a cycle, store it for future reference
	dc = &DirectedCycle{
		cycle: make([]int32, cycle_stack.Len()),
	}

	for i := 0; !cycle_stack.Empty(); i++ {
		dc.cycle[i] = cycle_stack.Pop().(int32)
	}

	return
}

func (dc *DirectedCycle) IsAcyclic() bool { return len(dc.cycle) == 0 }
func (dc *DirectedCycle) Cycle() []int32  { return dc.cycle }
