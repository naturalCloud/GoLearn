package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {

	dial, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		panic(err)
	}

	writer := bufio.NewWriterSize(dial, 512)
	reader := bufio.NewReaderSize(dial, 512)

	for {
		str, _ := reader.ReadString('\n')
		fmt.Println("read", str)

		writer.Write([]byte("hello word"))
		writer.Flush()
	}

}
