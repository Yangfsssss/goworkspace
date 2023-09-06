package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"sync"
)

func main() {
	ports := os.Args[1:]
	var wg sync.WaitGroup

	for _, port := range ports {
		wg.Add(1)
		go dialAndPrint(port, &wg)
	}

	wg.Wait()
}

func dialAndPrint(port string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("dialoguing port: " + port)
	conn, err := net.Dial("tcp", "localhost:"+port)
	if err != nil {
		log.Println(err)
		return
	}

	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Println(err)
	}
}
