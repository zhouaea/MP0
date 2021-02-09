package main

import (
	"MP0v2/emails"
	"MP0v2/errorchecking"
	"bufio"
	"fmt"
	"net"
)

func main() {
	// Prompt user for email to create an email object.
	email := emails.EmailPrompt()

	// Establish a connection with tcp port on localhost server.
	c, err := net.Dial("tcp", "127.0.0.1:1234")
	errorchecking.CheckError(err, "A")

	// Write email into connection with host server.
	fmt.Fprintf(c, email.String())

	// Wait for acknowledgement of host server and then exit.
	for {
		message, _ := bufio.NewReader(c).ReadString('\n')
		// TODO: Crashes when I check for error but works fine otherwise.
		//errorchecking.CheckError(err, "A")

		if message == "roger that" {
			fmt.Println("Acknowledgement received from server. TCP client exiting...")
			return
		}
	}
}
