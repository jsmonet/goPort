package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

var (
	host    = flag.String("h", "", "Hostname or IP address (string)")
	port    = flag.Int("p", 22, "TCP port to test (int)")
	timeout = flag.Int("t", 5, "Timeout for the request in seconds")
)

func main() {
	flag.Parse()
	checkResult, err := checkPort(*host, *port, *timeout)
	if err != nil {
		fmt.Println(err)
	}

	os.Exit(checkResult)
}

func checkPort(address string, port int, timeout int) (result int, checkErr error) {
	result = 0                                    // explicitly zeroing
	target := fmt.Sprintf("%v:%v", address, port) // net.DialTimeout requires address:22 formatting, where 22 is just the int port number
	timeOutSeconds := time.Duration(timeout) * time.Second
	_, checkErr = net.DialTimeout("tcp", target, timeOutSeconds)
	if checkErr != nil {
		result = 1
	}
	return result, checkErr
}
