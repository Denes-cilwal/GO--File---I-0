package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	fmt.Println("Welcome to text File read / write  in golang")
	content := "This need to go in new file created-"

	// create file
	file, err := os.Create("./info.txt")
	CheckNilErr(err)

	// write to file
	// WriteString writes the contents of the string content to file,
	length, err := io.WriteString(file, content)
	CheckNilErr(err)

	fmt.Println("length is:", length)

	defer file.Close()

	ReadFile("./info.txt")
}

func ReadFile(fileName string) {
	// ioutil package

	// make([]byte, 30) - data read will be the size of the byte array
	bytedata, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Println(bytedata, "bytedate inside file is--")

}

func CheckNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
