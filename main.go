package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var (
	netAddress    string
	netPort       string
	bytesTransfer int
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func randBytes(n int) []byte {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return b
}

func main() {
	flag.StringVar(&netAddress, "a", "localhost", "Specify IP address. Default is localhost")
	flag.StringVar(&netPort, "p", "9090", "Specify network port. Default is 9090")
	flag.IntVar(&bytesTransfer, "b", 10, "Specify kilo-bytes to send. Default is 10KB")
	flag.Parse()
	netSocket := netAddress + ":" + netPort
	fmt.Println("Starting TCP Client")
	fmt.Printf("Connecting to: %v\n", netSocket)
	tcpConnection, err := net.Dial("tcp", netSocket)
	CheckError(err)
	defer tcpConnection.Close()
	bytesToSend := randBytes(bytesTransfer * 1024)
	bytesLength, err := tcpConnection.Write(bytesToSend)
	CheckError(err)
	fmt.Printf("Sent %v KB!\n", float64(bytesLength/1024.0))
}
