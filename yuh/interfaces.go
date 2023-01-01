package main

import (
	"bytes"
	"fmt"
)

type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct {
}

func (cw *ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

// Incrementer implement khong su dung struct
type Incrementer interface {
	Increment() int
}

type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}

// Writer2 implement nhieu, interface nay chua interface kia
type Writer2 interface {
	Write2([]byte) (int, error)
}

type Closer interface {
	Close() error
}

type WriterCloser interface {
	Writer2
	Closer
}

type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

func (bwc *BufferedWriterCloser) Write2(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}
	v := make([]byte, 8)
	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}
		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}
	return n, nil
}

func (bwc *BufferedWriterCloser) Close() error {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

func NewBufferWriterCloser() *BufferedWriterCloser {
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}
func main() {
	var w Writer = &ConsoleWriter{}
	w.Write([]byte("Hello world!"))
	//
	counter := IntCounter(0)
	var inc Incrementer = &counter
	for i := 1; i <= 5; i++ {
		fmt.Println(inc.Increment())
	}
	// buffer
	var wc WriterCloser = NewBufferWriterCloser()
	wc.Write2([]byte("Happy new year 2023"))
	wc.Close()
	fmt.Println()

	// interface rong
	var obj interface{} = NewBufferWriterCloser()
	wc, err := obj.(WriterCloser)
	if err {
		fmt.Println("Da co loi xay ra khi chuyen kieu")
	}
	wc.Write2([]byte("chuc mung ngay dau nam"))
	wc.Close()
}
