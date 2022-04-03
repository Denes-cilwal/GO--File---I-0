# GO-File-I-0

✅ [Reading data and writing file using Os package](https://golangdocs.com/io-package-in-golang) - ✔

✅ [Reading data and writing using buffered I/O.](https://www.educative.io/edpresso/how-to-read-and-write-with-golang-bufio) - ✔

✅ [Buffered reading and Scanning, Reader | Writer Interface](https://www.educative.io/edpresso/how-to-read-and-write-with-golang-bufio) - ✔

✅ [Buffers and streams](https://www.youtube.com/watch?v=-O8wpYYKstE&t=655s) - ✔

✅ [Buffers and streams](https://www.youtube.com/watch?v=GlybFFMXXmQ&t=138s) - ✔


- The io package provides an interface for basic I/O primitives and wraps them into shared public interfaces that abstracts the functionality.

- The ioutil package provides utility functions for I/O operations.
  Note that as of Go 1.16 the same functionalities can be accessed through io or os package directly.

- The bufio provides an interface for buffered I/O operation with the file. Buffer is actually a temporary space in memory where data is stored and I/O       operations are performed from this temporary space. It also means that if we are not using bufio we are basically having unbuffered I/O         operations.Typically all I/O operations are unbuffered unless specified. The key advantage of having a buffer is that it minimizes system calls as well as disk I/O and is particularly suitable for block transfer of data. This is not suitable for single character-oriented I/O operations.
