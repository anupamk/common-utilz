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
// couple of commonly used algorithms for graphs are provided in this
// package.
//
// this file provides test+benchmark routines for connected-components
package algorithms

import (
	"fmt"
	"github.com/anupamk/common-utilz/graph"
	"testing"
)

func ExampleConnectedComponents() {
	fname := "../data/graph-001.data"
	g, _ := graph.LoadFromFile(fname)

	cc := New(g)
	fmt.Println(cc)

	// Output:
	// total-components: 3
	// 0: 1
	// 1: 1
	// 2: 1
	// 3: 1
	// 4: 1
	// 5: 1
	// 6: 1
	// 7: 2
	// 8: 2
	// 9: 3
	// 10: 3
	// 11: 3
	// 12: 3
}

func TestCCCount(t *testing.T) {
	fname := "../data/graph-001.data"
	g, _ := graph.LoadFromFile(fname)

	cc := New(g)
	cc_count := cc.Count()

	if cc_count != 3 {
		t.Logf("failed: expected-count: 3, got: %d\n", cc_count)
		t.Fail()
	}
}

func TestCCIsConnected(t *testing.T) {
	fname := "../data/graph-001.data"
	g, _ := graph.LoadFromFile(fname)

	cc := New(g)

	if connected, _ := cc.IsConnected(0, 10); connected {
		t.Logf("failed: expected: false, got: true")
		t.Fail()
	}

	if connected, _ := cc.IsConnected(10, 12); !connected {
		t.Logf("failed: expected: true, got: false")
		t.Fail()
	}
}

func BenchmarkConnectedComponents(bench *testing.B) {
	fname := "../data/graph-004.data"
	g, _ := graph.LoadFromFile(fname)
	bench.ResetTimer()

	for i, _ := 0, New(g); i < bench.N; i++ {
	}
}
