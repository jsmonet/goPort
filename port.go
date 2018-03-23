package main

import (
	"flag"
	"fmt"
	"net"
	"time"
)

var (
	host    = flag.String("h", "", "Hostname or IP address (string)")
	port    = flag.Int("p", 22, "TCP port to test (int)")
	timeout = flag.Int("t", 5, "Timeout for the request in seconds")
)

func main() {
	flag.Parse()
	checkResult := checkPort(*host, *port, *timeout)
	if checkResult != 0 {
		fmt.Println("no good") // oh my god i hate first pass results
	} else {
		fmt.Println("omg yay") // delete this, destroy this, stop doing this
	}
}

func checkPort(address string, port int, timeout int) (result int) {
	result = 0                                    // explicitly zeroing
	target := fmt.Sprintf("%v:%v", address, port) // net.DialTimeout requires address:22 formatting, where 22 is just the int port number
	timeOutSeconds := time.Duration(timeout) * time.Second
	_, err := net.DialTimeout("tcp", target, timeOutSeconds)
	if err != nil {
		result = 2
	}

	return result

}
