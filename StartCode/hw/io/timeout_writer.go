package io

import "time"

type TimeoutWriter struct {
	duration time.Duration
}

func NewTimeoutWriter(duration time.Duration) *TimeoutWriter {
	return &TimeoutWriter{
		duration: duration,
	}
}

func (w *TimeoutWriter) Write(b []byte) (int, error) {
	time.Sleep(w.duration)
	return len(b), nil
}
