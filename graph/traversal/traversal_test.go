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
package traversal

import (
	"github.com/anupamk/common-utilz/graph"
	"testing"
)

func TestBFSTraversal(t *testing.T) {
	fname := "../data/graph-001.data"
	sv := int32(0)
	g, _ := graph.LoadGraphFromFile(fname)

	gw := BFSWalker(g, sv)
	for n, err := gw(); err != EOG; n, err = gw() {
		t.Logf("visit-node: %d\n", n)
	}
}

func BenchmarkBFSTraversal(bench *testing.B) {
	fname := "../data/graph-003.data"
	g, _ := graph.LoadGraphFromFile(fname)
	bench.ResetTimer()

	for i := 0; i < bench.N; i++ {
		for v := int32(0); v < g.V(); v++ {
			gw := BFSWalker(g, v)
			for _, err := gw(); err != EOG; _, err = gw() {
			}
		}
	}
}

func TestDFSTraversal(t *testing.T) {
	fname := "../data/graph-001.data"
	sv := int32(0)
	g, _ := graph.LoadGraphFromFile(fname)

	gw := DFSWalker(g, sv)

	for n, err := gw(); err != EOG; n, err = gw() {
		t.Logf("visit-node: %d\n", n)
	}
}

func BenchmarkDFSTraversal(bench *testing.B) {
	fname := "../data/graph-003.data"
	g, _ := graph.LoadGraphFromFile(fname)
	bench.ResetTimer()

	for i := 0; i < bench.N; i++ {
		for v := int32(0); v < g.V(); v++ {
			gw := DFSWalker(g, v)
			for _, err := gw(); err != EOG; _, err = gw() {
			}
		}
	}
}
