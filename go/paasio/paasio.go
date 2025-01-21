package paasio

import (
	"io"
	"sync"
)

// Define readCounter and writeCounter types here.

// For the return of the function NewReadWriteCounter, you must also define a type that satisfies the ReadWriteCounter interface.

type Counter struct {
	bytesCount int64
	opsCount   int
	mu         *sync.Mutex
}

func (c *Counter) AddBytes(bytes int64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.opsCount++
	c.bytesCount += bytes
}

func (c *Counter) count() (int64, int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.bytesCount, c.opsCount
}

type readCounter struct {
	reader io.Reader
	Counter
}

type writeCounter struct {
	writer io.Writer
	Counter
}

type readWriteCounter struct {
	ReadCounter
	WriteCounter
}

func NewWriteCounter(writer io.Writer) WriteCounter {
	return &writeCounter{
		writer,
		Counter{
			bytesCount: 0,
			opsCount:   0,
			mu:         &sync.Mutex{},
		},
	}
}

func NewReadCounter(reader io.Reader) ReadCounter {
	return &readCounter{
		reader,
		Counter{
			bytesCount: 0,
			opsCount:   0,
			mu:         &sync.Mutex{},
		},
	}
}

func NewReadWriteCounter(readwriter io.ReadWriter) ReadWriteCounter {
	return &readWriteCounter{
		NewReadCounter(readwriter),
		NewWriteCounter(readwriter),
	}
}

func (rc *readCounter) Read(p []byte) (int, error) {
	readBytes, err := rc.reader.Read(p)
	// if err != nil {
	// 	return 0, err
	// }
	rc.AddBytes(int64(readBytes))
	return readBytes, err
}

func (rc *readCounter) ReadCount() (int64, int) {
	return rc.count()
}

func (wc *writeCounter) Write(p []byte) (int, error) {
	writtenBytes, err := wc.writer.Write(p)
	// if err != nil {
	// 	return 0, err
	// }
	wc.AddBytes(int64(writtenBytes))
	return writtenBytes, err
}

func (wc *writeCounter) WriteCount() (int64, int) {
	return wc.count()
}
