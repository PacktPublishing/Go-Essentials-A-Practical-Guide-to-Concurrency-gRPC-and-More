package main

import (
	"fmt"
	"net"
	"sort"
)

func scan(ports, results chan int){
	for p := range ports {
		//address:port
		address := fmt.Sprintf("scanme.nmap.org:%d", p)

		fmt.Printf("Address is: %s\n", address)

		conn, err := net.Dial("tcp", address)

		if err != nil {
			results <- 0
			continue
		}

		conn.Close()
		results <- p
	}
}

func main() {
	ports := make(chan int, 1023)

	results := make(chan int)

	var openPorts []int

	for i := 0; i < cap(ports); i++ {
		go scan(ports, results)
	}

	go func(){
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()

	for i := 0; i < 1024; i++ {
		port := <-results
		if port != 0 {
			openPorts = append(openPorts, port)
		}
	}

	close(ports)
	close(results)

	sort.Ints(openPorts)

	for _, port := range openPorts {
		fmt.Printf("%d open\n", port)
	}

}
