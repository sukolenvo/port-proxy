package main

import (
	"flag"
	"fmt"
	"net"
)

func main() {
	addr := flag.String("address", "0.0.0.0", "Interface to listen on")
	port := flag.String("port", "8080", "Interface to listen on")
	connectAddress := flag.String("connectAddress", "127.0.0.1", "Address to connect to")
	connectPort := flag.String("connectPort", "", "Port to connect to (default same as -port)")
	flag.Parse()
	if *connectPort == "" {
		connectPort = port
	}
	if *addr+":"+*port == *connectAddress+":"+*connectPort || (*connectAddress == "127.0.0.1" && *port == *connectPort) {
		fmt.Println("Listen address should not be same as remote address")
		return
	}
	l, err := net.Listen("tcp4", *addr+":"+*port)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c, *connectAddress+":"+*connectPort)
	}
}

func handleConnection(c net.Conn, remoteAddr string) {
	fmt.Printf("Received connection from %s\n", c.RemoteAddr().String())
	remote, err := net.Dial("tcp4", remoteAddr)
	if err != nil {
		fmt.Printf("Failed to connect to %s\n", remoteAddr)
		c.Close()
		return
	}
	go forwardByteStream(c, remote)
	go forwardByteStream(remote, c)
}

func forwardByteStream(from net.Conn, to net.Conn) {
	defer from.Close()
	defer to.Close()
	buffer := make([]byte, 4096)
	for {
		count, err := from.Read(buffer)
		if err != nil {
			return
		}
		write, err := to.Write(buffer[0:count])
		if err != nil {
			fmt.Println(err)
			return
		}
		for write != count {
			write, err = to.Write(buffer[0:count])
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}
