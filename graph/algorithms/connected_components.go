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
// this file provides the "connected-component" implementation which
// separates vertices of a graph into equivalece classes
//
package algorithms

import (
	"fmt"
	"github.com/anupamk/common-utilz/graph"
	"github.com/anupamk/common-utilz/traversal"
)

type ConnectedComponent struct {
	id    []int32
	count int32
}

//
// find all connected-components in a graph using breadth-first-search
//
func New(G graph.GraphOps) (CC *ConnectedComponent) {
	CC = &ConnectedComponent{
		id:    make([]int32, G.V()),
		count: 0,
	}
	marked := make([]bool, G.V())

	// each vertex is it's own component for starters
	for i, _ := range CC.id {
		CC.id[i] = int32(i)
	}

	for i, seen := range marked {
		if !seen {
			CC.count += 1
			bfs_walker := traversal.BFSGraphSubsetWalker(G, int32(i))
			for E, err := bfs_walker(); err != traversal.EOGS; E, err = bfs_walker() {
				CC.id[E.Dst] = CC.count
				marked[E.Dst] = true
			}
		}
	}

	return
}

//
// returns true if v is connected to w false otherwise. signals an
// error if either v/w are invalid vertices for the graph
//
func (CC *ConnectedComponent) IsConnected(v, w int32) (yesno bool, err error) {
	if (v > int32(len(CC.id))) || (w > int32(len(CC.id))) {
		err = fmt.Errorf("bogus vertex: %d or %d\n", v, w)
		return
	}

	yesno = (CC.id[v] == CC.id[w])
	return
}

//
// returns the number of connected components in the source graph
//
func (CC *ConnectedComponent) Count() int32 { return CC.count }

// regular stuff

func (CC *ConnectedComponent) String() string {
	str := ""

	str += fmt.Sprintf("total-components: %d\n", CC.count)
	for i, cc := range CC.id {
		str += fmt.Sprintf("%d: %d\n", i, cc)
	}

	return str
}
