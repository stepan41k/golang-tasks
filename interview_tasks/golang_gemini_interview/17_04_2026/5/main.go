package main

import (
	"fmt"
)

package main

import (
	"fmt"
	"sync"
)

type BufferPool struct {
	size int
	buf sync.Pool
}

func NewBufferPool(size int) *BufferPool {
	return &BufferPool{
		size: size,
		buf: sync.Pool{
			New: func() any {
				return make([]byte, size)
			},
		},
	}
}

func (bp *BufferPool) GetBuffer() []byte {
	buf := bp.buf.Get().([]byte)
	return buf
}

func (bp *BufferPool) PutBuffer(buf []byte) {
	if cap(buf) < bp.size {
		return
	}
	buf = buf[:bp.size]
	
	clear(buf)
	
	bp.buf.Put(buf)
}

func main() {
	bp := NewBufferPool(1024)
	buf := bp.GetBuffer()
	
	for i := 0; i < 1000; i++ {
		buf = append(buf, byte(i))
	}

	bp.PutBuffer(buf)
}
