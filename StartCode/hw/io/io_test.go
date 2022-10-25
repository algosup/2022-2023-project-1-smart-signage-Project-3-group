package io

import (
	"errors"
	"io"
	"testing"
	"time"
)

func TestTimeoutWriter_Write(t *testing.T) {
	timeout := time.Second
	w := NewTimeoutWriter(timeout)
	data := "Hello, World"

	start := time.Now()
	n, err := w.Write([]byte(data))
	end := time.Now()
	if err != nil {
		t.Fatal("the writer should be slow, but shouldn't fail")
	}

	if n != len(data) {
		t.Fatal("the writer should write the whole data")
	}

	if end.Sub(start) < timeout {
		t.Fatal("the writer should write in more than", timeout)
	}

}

func TestErrorWriter_Write(t *testing.T) {
	expErr := errors.New("my fancy error")

	w := NewErrorWriter(expErr)
	data := "Hello, World"

	_, err := w.Write([]byte(data))
	if err != expErr {
		t.Fatalf("didn't write the expected error: %v, got: %v", expErr, err)
	}
}

func TestSlowWriter_Write(t *testing.T) {
	speed := 2

	w := NewSlowWriter(speed)
	data := "Hello, World"

	for {
		n, err := w.Write([]byte(data))
		if err != nil {
			if err != io.EOF {
				break
			}
			t.Fatal("the writer should be slow, but shouldn't fail")
		}

		if n != speed {
			t.Fatal("the writer doesn't write at the specified speed")
		}
		data = data[:n]
		if string(w.currentData) == data {
			break
		}
	}

}

func TestGarbageWriter_Write(t *testing.T) {
	w := NewGarbageWriter()

	data := "Hello, World"

	// we don't care about n, because it writes garbage, it might not be the right size even
	_, err := w.Write([]byte(data))
	if err != nil {
		t.Fatal("the writer should write garbage, but not fail")
	}

	t.Log(string(w.writtenData))
	if data == string(w.writtenData) {
		t.Fatal("the data written shouldn't be the same as the original one")
	}

}
