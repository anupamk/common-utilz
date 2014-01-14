//
// this package provides an implementation of a 'chunked' stack. each
// element in the stack is a chunk of 'MAX_ITEMS_PER_CHUNK' elements.
//
package stack

const MAX_ITEMS_PER_CHUNK = 4096

type Data interface{}

type chunk struct {
	items    [MAX_ITEMS_PER_CHUNK]Data
	push_idx uint16
	pop_idx  uint16
	next     *chunk
}

func (sc *chunk) push(item Data) {
	sc.pop_idx = sc.push_idx
	sc.items[sc.push_idx] = item
	sc.push_idx++

	return
}

func (sc *chunk) pop() (item Data) {
	sc.push_idx = sc.pop_idx
	item = sc.items[sc.pop_idx]
	sc.pop_idx--

	return
}

//
// a stack made out of list of chunk elements
//
type ChunkedStack struct {
	top  *chunk
	size int32
}

// create a new stack
func New() *ChunkedStack {
	init_chunk := new(chunk)

	return &ChunkedStack{
		top:  init_chunk,
		size: 0,
	}
}

//
// this function is called to add an item to the top of the
// stack. grow the stack by a chunk if required.
//
func (cs *ChunkedStack) Push(item Data) {
	head := cs.top

	// if this chunk is full, create a new one
	if head.push_idx == MAX_ITEMS_PER_CHUNK {
		new_chunk := new(chunk)
		new_chunk.next = head
		cs.top = new_chunk

		head = new_chunk
	}

	head.push(item)
	cs.size++
}

//
// this function is called to remove an item from the stack. when the
// number of items in a chunk drops to 0, it is removed from the
// stack.
//
func (cs *ChunkedStack) Pop() (item Data) {
	head := cs.top
	item = head.pop()
	cs.size--

	// shrink if required
	if head.push_idx == 0 {
		if cs.size == 0 {
			cs.top = nil
			head.next = nil
		} else {
			cs.top = head.next
		}
	}

	return
}

func (cs *ChunkedStack) Len() int32  { return cs.size }
func (cs *ChunkedStack) Empty() bool { return !(cs.size > 0) }
