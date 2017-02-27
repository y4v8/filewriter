package logrotation

import (
	"io"
	"os"
)

type logRotation struct {
	name string
	perm os.FileMode

	setOutput func(writer io.Writer)

	f   *os.File
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func New(name string, perm os.FileMode, setOutput func(writer io.Writer)) (*logRotation, error) {
	f, err := createFile(name, perm)
	if err != nil {
		return nil, err
	}

	w := &logRotation{
		name:      name,
		perm:      perm,
		setOutput: setOutput,
		f:         f,
	}

	setOutput(f)

	return w, nil
}

func (w *logRotation) Close() {
	w.f.Close()
}

func (w *logRotation) FileName() string {
	return w.name
}

func (w *logRotation) Rotate() {
	f, err := createFile(w.name, w.perm)
	panicIf(err)

	w.setOutput(f)

	err = w.f.Sync()
	panicIf(err)
	err = w.f.Close()
	panicIf(err)

	w.f = f
}
