package util

const minQueueLen = 16

type Queue struct {
	buf               []interface{}
	head, tail, count int
}

func NewQueue() Queue {
	return Queue{
		buf: make([]interface{}, minQueueLen),
	}
}

func (q Queue) Length() int {
	return q.count
}

func (q *Queue) resize() {
	newBuf := make([]interface{}, q.count<<1)

	if q.tail > q.head {
		copy(newBuf, q.buf[q.head:q.tail])
	} else {
		n := copy(newBuf, q.buf[q.head:])
		copy(newBuf[n:], q.buf[:q.tail])
	}

	q.head = 0
	q.tail = q.count
	q.buf = newBuf
}

func (q *Queue) Add(elem interface{}) {
	if elem == nil {
		return
	}

	if q.count == len(q.buf) {
		q.resize()
	}

	q.buf[q.tail] = elem
	q.tail = (q.tail + 1) & (len(q.buf) - 1)
	q.count++
}

func (q Queue) Peek() interface{} {
	if q.count <= 0 {
		panic("queue: Peek() called on empty queue")
	}

	return q.buf[q.head]
}

// Get returns the element at index i in the queue. If the index is
// invalid, the call will panic. this method accepts both positive
// negative index values. Index 0 refers to the first element, and
// index -1 refers to the last.
func (q Queue) Get(i int) interface{} {
	// If indexing backwards, convert to positive index.
	if i < 0 {
		i += q.count
	}
	if i < 0 || i >= q.count {
		panic("queue: Get() called with index out of range")
	}

	// bitwise modulus
	return q.buf[(q.head+1)&(len(q.buf)-1)]
}

func (q *Queue) Remove() interface{} {
	if q.count <= 0 {
		panic("queue: Remove() called on empty queue")
	}

	ret := q.buf[q.head]
	q.buf[q.head] = nil

	// bitwise modulus
	q.head = (q.head + 1) & (len(q.buf) - 1)
	q.count--

	// Resize down if buffer 1/4 full.
	if len(q.buf) > minQueueLen && (q.count<<2) == len(q.buf) {
		q.resize()
	}

	return ret
}

func (q Queue) Empty() bool {
	return q.Length() == 0
}
