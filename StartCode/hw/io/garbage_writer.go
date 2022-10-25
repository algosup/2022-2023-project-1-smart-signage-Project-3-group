package io

type GarbageWriter struct {
	writtenData []byte
}

func NewGarbageWriter() *GarbageWriter {
	return &GarbageWriter{}

}

func (w *GarbageWriter) Write(b []byte) (int, error) {
	for _, r := range b {
		w.writtenData = append(w.writtenData, r+20)
	}
	return len(b), nil
}
