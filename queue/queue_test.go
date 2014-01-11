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
package queue

import (
	"testing"
)

// test out basic queue operations
func TestQueueOperations(t *testing.T) {
	var i int32
	var j int32

	nq := New()

	for i = 0; i < 16*MAX_ITEMS_PER_CHUNK; i++ {
		nq.Push(i)
	}

	if nq.Len() != i {
		t.Logf("pushed %d items, queue contains only %d", i, nq.Len())
	}

	for j = 0; nq.Len() > 0; j++ {
		nq.Pop()
	}

	if j != i {
		t.Logf("pushed %d items, popped %d items", i, j)
	}

	return
}

// benchmark the operations
func BenchmarkPush(b *testing.B) {
	nq := New()
	for i := 0; i < b.N; i++ {
		nq.Push(i)
	}
}

func BenchmarkPop(b *testing.B) {
	nq := New()

	for i := 0; i < b.N; i++ {
		nq.Push(i)
	}
	b.ResetTimer()

	for !nq.Empty() {
		nq.Pop()
	}
}
