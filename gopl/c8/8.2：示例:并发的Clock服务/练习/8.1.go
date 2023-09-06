package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

// var timezonePortMap = map[string]int{
// 	"NewYork": 8001,
// 	"London":  8002,
// 	"Tokyo":   8003,
// }

// 练习 8.1： 修改clock2来支持传入参数作为端口号，然后写一个clockwall的程序，
// 这个程序可以同时与多个clock服务器通信，从多个服务器中读取时间，并且在一个表格中一次显示所有服务器传回的结果，
// 类似于你在某些办公室里看到的时钟墙。

func main() {
	clockWall()
}

func clockWall() {
	timezoneAndPortsMap := make(map[string]int)
	timezoneAndPorts := os.Args[1:]

	for _, timezoneAndPort := range timezoneAndPorts {
		parts := strings.Split(timezoneAndPort, "=")
		port, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil {
			log.Fatal(err)
		}

		timezoneAndPortsMap[strings.TrimSpace(parts[0])] = port
	}

	fmt.Println("timezoneAndPortsMap", timezoneAndPortsMap)

	var wg sync.WaitGroup

	for zone, port := range timezoneAndPortsMap {
		wg.Add(1)
		// 为了传递循环变量的值给 listen 函数，以避免循环变量的副作用对并发执行的影响。
		go func(zone string, port int) {
			defer wg.Done()
			listen(zone, port)
		}(zone, port)
	}

	wg.Wait()
}

func listen(zone string, port int) {
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(zone + " is listening on port " + strconv.Itoa(port))

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}

		handleConn(conn, zone, port)

		time.Sleep(100 * time.Millisecond)
	}
}

func handleConn(c net.Conn, zone string, port int) {
	defer c.Close()

	for {
		_, err := io.WriteString(c, fmt.Sprintf("%-10s | %s\n", zone, generateCurrentTimeByTimezone(zone)))
		if err != nil {
			return // e.g. client disconnected
		}

		time.Sleep(1 * time.Second)
	}
}

func generateCurrentTimeByTimezone(zone string) string {
	location, err := time.LoadLocation(zone)
	if err != nil {
		log.Fatal(err)
	}

	return time.Now().In(location).Format("15:04:05\n")
}
