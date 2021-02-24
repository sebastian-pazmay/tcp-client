package main

import (
	"fmt"
	"math/rand"
	"net"
)

func CheckError(err error) {
	if err != nil {
		// TODO:
		// Fix error catching
		//fmt.Println("Error: " , err)
	}
}

// TODO:
// include read addr as argument
// use localhost by default
//const addr = "192.168.2.160:9090"
const addr = "localhost:9090"

func main() {
	fmt.Println("Starting TCP Client")
	// connect to this socket
	fmt.Printf("using given address: %v\n", addr)
	conn, _ := net.Dial("tcp", addr)
	defer conn.Close()
	/*fmt.Printf("Current Unix Time: %v\n", time.Now().Unix())

	time.Sleep(100 * time.Millisecond)
	//time.Sleep(1 * time.Second)

	fmt.Printf("Current Unix Time: %v\n", time.Now().Unix())*/
	//for {
	// read in input from stdin
	//reader := bufio.NewReader(os.Stdin)
	//fmt.Print("Text to send: ")
	//text := ("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	//1Kb
	text := randBytes(1 * 1024)
	//10Kb
	//text := randBytes(20 * 1024)
	//100Kb
	//text := randBytes(500 * 1024)
	//1MB
	//text := randBytes(1024 * 1024)
	//2MB
	//text := randBytes(1 * 1024 * 1024)
	//100MB
	//text := randBytes(100 * 1024 * 1024)
	// send to socket
	//fmt.Fprintf(conn, text + "\n")
	//fmt.Fprintf(conn, text)
	//fmt.Println(text)
	length, err := conn.Write(text)
	CheckError(err)
	// log the bytes written
	fmt.Printf("WROTE %d bytes\n", length)
	//Waiting for the reply
	//buf_rec := make([]byte, 104857600)
	//n,err := conn.Read(buf_rec)
	//fmt.Println(addr , string(buf_rec[0:n]))
	//CheckError(err)
	//}
	//conn.Close()
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randBytes(n int) []byte {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return b
}
