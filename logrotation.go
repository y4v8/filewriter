package logrotation

import (
	"bytes"
	"io"
	"os"
	"sync"
)

type logRotation struct {
	name string
	flag int
	perm os.FileMode

	setOutput func(writer io.Writer)

	f   *os.File
	buf *bytes.Buffer
	m   sync.Mutex
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func New(name string, flag int, perm os.FileMode, setOutput func(writer io.Writer)) (*logRotation, error) {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}

	data := make([]byte, 0, 256)
	w := &logRotation{
		name:      name,
		flag:      flag,
		perm:      perm,
		setOutput: setOutput,
		f:         f,
		buf:       bytes.NewBuffer(data),
	}

	setOutput(f)

	return w, nil
}

func (w *logRotation) Write(p []byte) (int, error) {
	w.m.Lock()
	n, err := w.buf.Write(p)
	w.m.Unlock()
	return n, err
}

func (w *logRotation) Close() {
	w.f.Close()
}

func (w *logRotation) Rotate(rotFileName string) {
	w.buf.Reset()

	w.setOutput(w)

	err := w.f.Sync()
	panicIf(err)
	err = w.f.Close()
	panicIf(err)

	err = os.Rename(w.name, rotFileName)
	panicIf(err)

	w.f, err = os.OpenFile(w.name, w.flag, w.perm)
	panicIf(err)

	w.setOutput(w.f)

	rotFile, err := os.OpenFile(rotFileName, w.flag, w.perm)
	panicIf(err)

	_, err = w.buf.WriteTo(rotFile)
	panicIf(err)

	err = rotFile.Sync()
	panicIf(err)
	err = rotFile.Close()
	panicIf(err)
}
