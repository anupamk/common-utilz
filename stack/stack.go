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
