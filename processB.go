package main

import (
	"MP0v2/errorchecking"
	"bufio"
	"fmt"
	"net"
)

func main() {
	// Listen to an unused TCP port on localhost.
	listener, err := net.Listen("tcp", ":1234")
	errorchecking.CheckError(err, "B")
	fmt.Println("Listening to tcp port was successful!")
	defer listener.Close()

	// Wait for a connection from a client to our TCP port and then set up a TCP channel with them.
	c, err := listener.Accept()
	errorchecking.CheckError(err, "B")
	fmt.Println("Connection to client was successful!")

	// Read and print email sent by client through TCP channel
	email, err := bufio.NewReader(c).ReadString('\t')
	errorchecking.CheckError(err, "B")
	fmt.Println("Email received!")
	fmt.Println("---------------")
	fmt.Println(email)
	fmt.Println("---------------")

	// Send acknowledgement to client.
	fmt.Fprintf(c, "roger that")
	fmt.Println("Sent acknowledgement to client!")

	fmt.Println("Exiting program...")
	return
}