package main

import (
	"MP0v2/errorchecking"
	"bufio"
	"fmt"
	"net"
)

func main() {
	// Listen to TCP port on localhost.
	listener, err := net.Listen("tcp", ":1234")
	errorchecking.CheckError(err, "B")
	fmt.Println("Listening to tcp port was successful!")
	defer listener.Close()

	// Wait for connection from a client.
	c, err := listener.Accept()
	errorchecking.CheckError(err, "B")
	fmt.Println("Connection to client was successful!")

	// Serve client as long as process is activated.
	for {
		// Read and print email sent by client.
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
}