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
package stack

import (
	"fmt"
	"testing"
)

func ExamplePushPop() {
	st := New()

	st.Push(1)
	st.Push(2)
	st.Push(3)
	st.Push(4)

	v1 := st.Pop()
	v2 := st.Pop()
	v3 := st.Pop()
	v4 := st.Pop()

	fmt.Printf("%d %d %d %d\n", v1, v2, v3, v4)

	// Output:
	// 4 3 2 1
}

func TestPushPopPush(t *testing.T) {
	st := New()

	// stack is empty after Pop
	st.Push(1)
	st.Pop()

	// can we push stuff on an empty stack ?
	st.Push(2)
	st.Pop()
}

func TestStackOperations(t *testing.T) {
	var i, j int32

	ns := New()

	for i = 0; i < 16*MAX_ITEMS_PER_CHUNK; i++ {
		ns.Push(i)
	}

	if ns.Len() != i {
		t.Logf("pushed %d items, stack contains only %d", i, ns.Len())
	}

	for j = 0; ns.Len() > 0; j++ {
		ns.Pop()
	}

	if j != i {
		t.Logf("pushed %d items, popped %d items", i, j)
	}

	return
}

// benchmark the operations
func BenchmarkPush(b *testing.B) {
	ns := New()
	for i := 0; i < b.N; i++ {
		ns.Push(i)
	}
}

func BenchmarkPop(b *testing.B) {
	ns := New()

	for i := 0; i < b.N; i++ {
		ns.Push(i)
	}
	b.ResetTimer()

	for !ns.Empty() {
		ns.Pop()
	}
}
