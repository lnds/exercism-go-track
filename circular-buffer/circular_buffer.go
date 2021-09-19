package circular

import "fmt"

type Buffer struct {
	ring  []byte
	read  int
	write int
	count int
}

func NewBuffer(size int) *Buffer {
	return &Buffer{
		ring:  make([]byte, size),
		read:  0,
		write: 0,
		count: 0,
	}
}

func (buf *Buffer) ReadByte() (byte, error) {
	if buf.count == 0 {
		return 0, fmt.Errorf("buffer empty")
	}
	val := buf.ring[buf.read]
	buf.read = (buf.read + 1) % len(buf.ring)
	buf.count--
	return val, nil
}

func (buf *Buffer) WriteByte(c byte) error {
	if buf.count == len(buf.ring) {
		return fmt.Errorf("buffer full")
	}
	buf.ring[buf.write] = c
	buf.write = (buf.write + 1) % len(buf.ring)
	buf.count++
	return nil
}

func (buf *Buffer) Overwrite(c byte) {
	buf.ring[buf.write] = c
	buf.write = (buf.write + 1) % len(buf.ring)
	if buf.count == len(buf.ring) {
		buf.read = buf.write
	} else {
		buf.count++
	}
}

func (buf *Buffer) Reset() {
	buf.write = 0
	buf.read = 0
	buf.count = 0
}
