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
// this file implements various traversal orders for a given
// graph. although mostly applicable for digraphs, we dont enforce it
// for now...
//
package traversal

import (
	"github.com/anupamk/common-utilz/graph"
	"github.com/anupamk/common-utilz/queue"
	"github.com/anupamk/common-utilz/stack"
)

type DFSTraversalOrder struct {
	pre     []int32
	post    []int32
	revpost []int32
}

func DoDFSTraversals(G graph.GraphOps) *DFSTraversalOrder {
	var do_dfs func(graph.GraphOps, int32)

	preq := queue.New()
	postq := queue.New()
	revpost := stack.New()
	visited := make([]bool, G.V())

	// the dfs procedure
	do_dfs = func(G graph.GraphOps, v int32) {
		preq.Push(v)

		// the real thang
		visited[v] = true
		for _, w := range G.Adj(v) {
			if !visited[w] {
				do_dfs(G, w)
			}
		}

		postq.Push(v)
		revpost.Push(v)
	}

	// run on the whole graph
	for v := int32(0); v < G.V(); v++ {
		if !visited[v] {
			do_dfs(G, v)
		}
	}

	retval := &DFSTraversalOrder{
		pre:     make([]int32, preq.Len()),
		post:    make([]int32, preq.Len()),
		revpost: make([]int32, preq.Len()),
	}

	// convert each to appropriate type...
	for i := 0; !preq.Empty(); i++ {
		retval.pre[i] = preq.Pop().(int32)
	}

	for i := 0; !postq.Empty(); i++ {
		retval.post[i] = postq.Pop().(int32)
	}

	for i := 0; !revpost.Empty(); i++ {
		retval.revpost[i] = revpost.Pop().(int32)
	}

	return retval
}

func (T *DFSTraversalOrder) PreOrder() []int32    { return T.pre }
func (T *DFSTraversalOrder) PostOrder() []int32   { return T.post }
func (T *DFSTraversalOrder) ReversePost() []int32 { return T.revpost }
