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
// this package provides an implementation of fifo-queue. each element
// in the queue is a chunk of MAX_ITEMS_PER_CHUNK elements. currently
// only fifo-esque operations i.e. add-to-tail and remove-from-head
// are supported...

package queue

const MAX_ITEMS_PER_CHUNK = 4096

type Data interface{}

type chunk struct {
	items [MAX_ITEMS_PER_CHUNK]Data
	first uint16
	last  uint16
	next  *chunk
}

// add an item to the chunk
func (C *chunk) push(item Data) {
	C.items[C.last] = item
	C.last++
}

// remove an item from the chunk
func (C *chunk) pop() (item Data) {
	item = C.items[C.first]
	C.first++

	return
}

//
// a queue made out of list of chunk elements
//
type ChunkedQueue struct {
	front *chunk
	back  *chunk
	size  int32
}

// create a new queue
func New() *ChunkedQueue {
	ch := new(chunk)

	return &ChunkedQueue{
		front: ch,
		back:  ch,
		size:  0,
	}
}

//
// this function is called to add an item to the end of the
// queue. grow the queue by a chunk if required.
//
func (CQ *ChunkedQueue) Push(item Data) {
	tail := CQ.back

	// if the tail chunk is full, create a new one
	if tail.last == MAX_ITEMS_PER_CHUNK {
		new_chunk := new(chunk)
		tail.next = new_chunk
		CQ.back = new_chunk

		tail = new_chunk
	}

	tail.push(item)
	CQ.size++
}

//
// this function is called to remove an item from the head of the
// queue. when the number of items contained in a chunk drops to 0, it
// is removed from the queue
//
func (CQ *ChunkedQueue) Pop() (item Data) {

	head := CQ.front
	item = head.pop()
	CQ.size--

	if head.first >= head.last {
		if CQ.size == 0 {
			head.first = 0
			head.last = 0
			head.next = nil
		} else {
			CQ.front = head.next
		}
	}

	return
}

func (CQ *ChunkedQueue) Len() int32  { return CQ.size }
func (CQ *ChunkedQueue) Empty() bool { return CQ.size == 0 }
