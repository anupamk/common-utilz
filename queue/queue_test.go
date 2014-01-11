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
