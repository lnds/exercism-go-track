package paasio

import (
	"io"
	"sync"
)

type ReadCounterImpl struct {
	reader    io.Reader
	bytesRead int64
	calls     int
	sync.RWMutex
}

func NewReadCounter(reader io.Reader) ReadCounter {
	return &ReadCounterImpl{reader: reader}
}

func (reader *ReadCounterImpl) Read(buf []byte) (n int, err error) {
	n, err = reader.reader.Read(buf)
	reader.Lock()
	defer reader.Unlock()
	reader.calls++
	reader.bytesRead += int64(n)
	return
}

func (reader *ReadCounterImpl) ReadCount() (int64, int) {
	reader.RLock()
	defer reader.RUnlock()
	return reader.bytesRead, reader.calls
}

type WriteCounterImpl struct {
	writer       io.Writer
	bytesWritten int64
	calls        int
	sync.RWMutex
}

func NewWriteCounter(writer io.Writer) WriteCounter {
	return &WriteCounterImpl{writer: writer}
}

func (writer *WriteCounterImpl) Write(p []byte) (n int, err error) {
	n, err = writer.writer.Write(p)
	writer.Lock()
	defer writer.Unlock()
	writer.calls++
	writer.bytesWritten += int64(n)
	return
}

func (writer *WriteCounterImpl) WriteCount() (n int64, nops int) {
	writer.RLock()
	defer writer.RUnlock()
	return writer.bytesWritten, writer.calls
}

type ReadWriteCounterImpl struct {
	ReadCounterImpl
	WriteCounterImpl
}

func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {
	return &ReadWriteCounterImpl{
		ReadCounterImpl{reader: rw},
		WriteCounterImpl{writer: rw},
	}
}
