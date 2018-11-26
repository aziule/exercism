package paasio

import (
    "io"
    "sync"
)

type counter struct {
    nBytes int64
    nOps int
    mutex *sync.Mutex
}

func (c *counter) addBytes(n int) {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    c.nOps++
    c.nBytes += int64(n)
}

func (c *counter) count() (int64, int) {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    return c.nBytes, c.nOps
}

func newCounter() *counter {
    return &counter{
        0,
        0,
        new(sync.Mutex),
    }
}

type readCounter struct {
    reader io.Reader
    rCounter *counter
}

type writeCounter struct {
    writer io.Writer
    wCounter *counter
}

type readWriteCounter struct {
    ReadCounter
    WriteCounter
}

func NewReadCounter(r io.Reader) ReadCounter {
    return &readCounter{r, newCounter()}
}

func NewWriteCounter(w io.Writer) WriteCounter {
    return &writeCounter{w, newCounter()}
}

func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {
    return &readWriteCounter{
        NewReadCounter(rw),
        NewWriteCounter(rw),
    }
}

func (r *readCounter) ReadCount() (int64, int) {
    return r.rCounter.count()
}

func (r *readCounter) Read(p []byte) (int, error) {
    bytes, err := r.reader.Read(p)
    r.rCounter.addBytes(bytes)
    return bytes, err
}

func (w *writeCounter) WriteCount() (int64, int) {
    return w.wCounter.count()
}

func (w *writeCounter) Write(p []byte) (int, error) {
    bytes, err := w.writer.Write(p)
    w.wCounter.addBytes(bytes)
    return bytes, err
}
