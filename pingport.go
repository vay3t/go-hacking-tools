package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	host := flag.String("host", "", "Target host")
	port := flag.Int("port", 0, "Port to ping")
	count := flag.Int("count", 0, "stop after <count> replies")

	flag.Parse()

	target := fmt.Sprintf("%s:%d", *host, *port)

	if *host != "" && *port > 0 {
		fmt.Printf("[+] %s\n", target)
		if *count > 0 {
			for i := 0; i < *count; i++ {
				pinger(target, *port)
			}
		} else if *count == 0 {
			for true {
				pinger(target, *port)
			}
		} else {
			usage()
		}

	} else {
		usage()
	}

}

func usage() {
	fmt.Println("Usage: " + os.Args[0] + " -host <HOST> -port <PORT>")
	flag.PrintDefaults()
}

func pinger(target string, portNum int) {
	_, err := net.Dial("tcp", target)
	if err == nil {
		log.Printf("port open: %d\n", portNum)
	} else {
		log.Printf("port not found: %d\n", portNum)
	}

	time.Sleep(1 * time.Second)

}
