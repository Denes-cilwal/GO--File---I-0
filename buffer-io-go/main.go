package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
- FILE I/O - Buffered I/O
   - Writing Data to File
   - Reading Data to File
   - Appending to to File
 - Working with series of data
*/

func main() {

	fmt.Println("Working with buffered file---")
	file, err := os.Create("./newfile.txt")

	CheckNilErr(err)

	// create a buffer for the os file
	buf := bufio.NewWriter(file)
	defer file.Close()

	// Write data to buffer
	content := "Writing Data to Buffer\n"
	buf.WriteString(content) //  write data to buffer but not move buffer data to file

	buf.Flush() // clears the buffer data and writes to file

	ReadDataFromFile("./newfile.txt")
}

func ReadDataFromFile(fileName string) {

	file, _ := os.Open(fileName)

	// load file data into buffer
	reader := bufio.NewReader(file)

	// no of bytes to read from buffer
	data, err := reader.Peek(5)
	CheckNilErr(err)

	fmt.Println(string(data), "readdata is--->")

}

func CheckNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
