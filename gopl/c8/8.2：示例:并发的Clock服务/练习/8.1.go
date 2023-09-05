package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
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

	for zone, port := range timezoneAndPortsMap {
		listen(zone, port)
	}
}

func listen(zone string, port int) {
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}

		go handleConn(conn, zone, port)
	}
}

func handleConn(c net.Conn, zone string, port int) {
	defer c.Close()

	for {
		_, err := io.WriteString(c, generateCurrentTimeByTimezone(zone))
		if err != nil {
			return // e.g. client disconnected
		}
	}
}

func generateCurrentTimeByTimezone(zone string) string {
	location, err := time.LoadLocation(zone)
	if err != nil {
		log.Fatal(err)
	}

	return time.Now().In(location).Format("15:04:05\n")
}
