package bufferq

import (
	"github.com/itsabgr/atomic2"
	"github.com/itsabgr/go-q"
)

type Q struct {
	list q.Q
	size atomic2.Uintptr
}

func (q *Q) Push(b []byte) {
	q.list.Push(b)
	q.size.Add(uintptr(len(b)))
}
func (q *Q) Len() uint {
	return uint(q.size.Get())
}
func (q *Q) Reset() {
	q.list.Reset()
	q.size.Set(0)
}
func (q *Q) Pull() []byte {
	ib, ok := q.list.Pull()
	if !ok {
		return nil
	}
	b := ib.([]byte)
	q.size.Sub(uintptr(len(b)))
	return b
}
func (q *Q) Peek() []byte {
	b, ok := q.list.Peek()
	if !ok {
		return nil
	}
	return b.([]byte)
}
