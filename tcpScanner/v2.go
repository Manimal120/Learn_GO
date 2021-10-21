package main

import (
	"fmt"
	"net"
	"sort"
)

func worker(ports chan int, res chan int) {
	for p := range ports {
		address := fmt.Sprintf("20.194.168.28:%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			res <- 0
			continue
		}
		conn.Close()
		res <- p
	}
}

func main() {
	ports := make(chan int, 100)
	res := make(chan int)
	var openPorts []int
	var closePorts []int

	for i := 0; i < cap(ports); i++ {
		go worker(ports, res)
	}

	for i := 1; i < 1024; i++ {
		ports <- i
	}

	go func() {
		for i := 1; i < 1024; i++ {
			ports <- i
		}
	}()

	for i := 1; i < 1024; i++ {
		port := <-res
		if port != 0 {
			openPorts = append(openPorts, port)
		} else {
			closePorts = append(closePorts, port)
		}
	}

	close(ports)
	close(res)

	sort.Ints(openPorts)
	sort.Ints(closePorts)

	//for _, port := range closePorts {
	//	fmt.Printf("%d closed \n", port)
	//}

	for _, port := range openPorts {
		fmt.Printf("%d open \n", port)
	}
}
