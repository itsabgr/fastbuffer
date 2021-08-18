package linkedbuffer

import (
	"github.com/itsabgr/fastbuffer/pkg/bufferq"
	"github.com/itsabgr/go-bytespool"
	"io"
	"sync"
)
type Buffer struct {
	readMutex   sync.Mutex
	bufferQ bufferq.Q
	offset  uint
}



func (r *Buffer) Write(src []byte) (n int,err error){
	for n < len(src){
		dst := bytespool.Pull(uint(len(src) - n))
		n += copy(dst,src[n:])
		err = r.Push(dst)
		if err!=nil{
			return n,err
		}
	}
	return
}
func (r *Buffer) Push(data []byte) error{
	r.bufferQ.Push(data)
	return nil
}
func (r *Buffer) Len() uint{
	return r.bufferQ.Len()
}
func (r *Buffer) Flush() {
	r.readMutex.Lock()
	defer r.readMutex.Unlock()
	r.bufferQ.Reset()
	r.offset = 0
}
func (r *Buffer) Read(dst []byte) (int,error){
	r.readMutex.Lock()
	defer r.readMutex.Unlock()
	src := r.bufferQ.Peek()
	if src == nil{
		return 0,io.EOF
	}
	src = src[r.offset:]
	n := copy(dst,src)
	if n == len(src){
		r.bufferQ.Pull()
	}else{
		r.offset += uint(n)
	}
	return n,nil
}

func (r *Buffer) Pull() ([]byte,error){
	r.readMutex.Lock()
	defer r.readMutex.Unlock()
	b := r.bufferQ.Pull()
	if b == nil{
		return nil,io.EOF
	}
	b = b[r.offset:]
	r.offset = 0
	return b,nil
}

func (r *Buffer) Peek(dst []byte) (int,error){
	r.readMutex.Lock()
	defer r.readMutex.Unlock()
	src := r.bufferQ.Peek()
	if src == nil{
		return 0,io.EOF
	}
	src = src[r.offset:]
	return copy(dst,src),nil
}

