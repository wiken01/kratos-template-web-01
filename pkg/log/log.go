package log

import (
	"io"

	"go.uber.org/multierr"
)

type WriteSyncer interface {
	io.Writer
}

// NewMultiWriteSyncer creates a WriteSyncer that duplicates its writes
// and sync calls, much like io.MultiWriter.
func NewMultiWriteSyncer(ws ...WriteSyncer) WriteSyncer {
	if len(ws) == 1 {
		return ws[0]
	}
	return multiWriteSyncer(ws)
}

type multiWriteSyncer []WriteSyncer

// When not all underlying syncers write the same number of bytes,
// the smallest number is returned even though Write() is called on
// all of them.
func (ws multiWriteSyncer) Write(p []byte) (int, error) {
	var writeErr error
	nWritten := 0
	for _, w := range ws {
		n, err := w.Write(p)
		writeErr = multierr.Append(writeErr, err)
		if nWritten == 0 && n != 0 {
			nWritten = n
		} else if n < nWritten {
			nWritten = n
		}
	}
	return nWritten, writeErr
}
