package filewriter

import (
	"os"
	"sync"
)

type file struct {
	name string
	perm os.FileMode

	f *os.File
	m sync.Mutex
}

// New returns a pointer to the new file.
func New(name string, perm os.FileMode) *file {
	w := &file{
		name: name,
		perm: perm,
	}

	return w
}

// Implementation of the Writer interface.
func (w *file) Write(p []byte) (n int, err error) {
	w.m.Lock()
	n, err = w.f.Write(p)
	w.m.Unlock()

	return
}

// Create closes, if necessary, and creates a file with the above name.
func (w *file) Create() (err error) {
	w.m.Lock()
	defer w.m.Unlock()

	if w.f != nil {
		if err = w.f.Close(); err != nil {
			return
		}
	}

	w.f, err = createFile(w.name, w.perm)

	return
}

// Close file.
func (w *file) Close() (err error) {
	w.m.Lock()
	err = w.f.Close()
	w.f = nil
	w.m.Unlock()

	return
}
