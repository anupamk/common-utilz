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
