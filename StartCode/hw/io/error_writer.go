package io

type ErrorWriter struct {
	err error
}

func NewErrorWriter(err error) *ErrorWriter {
	return &ErrorWriter{
		err: err,
	}
}

func (w *ErrorWriter) Write(b []byte) (int, error) {
	return 0, w.err
}
