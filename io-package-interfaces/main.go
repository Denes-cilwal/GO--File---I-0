package iopackageinterfaces

import (
	"fmt"
	"strings"
)

// Packages io provides basic interfaces to I/o(input/output) primitives

/*
	- Writer, Reader & seeker interfaces
	- Functions working alongside with these interfaces

 type Writer interface{
	 Write(p []byte)(n int,err error)

 }

*/

/*
The io.Reader interface represents an entity from which you can read a stream of bytes.

type Reader interface {
        Read(buf []byte) (n int, err error)
}
Read reads up to len(buf) bytes into buf and returns the number of bytes read – it returns an io.EOF error when the stream ends.

*/

func main() {

	// Read directly from a byte stream

	r := strings.NewReader("abcde")

	buf := make([]byte, 4)
	n, err := r.Read(buf)
	fmt.Println(n, err)
}
