package main

import (
	"MP0v2/emails"
	"MP0v2/errorchecking"
	"bufio"
	"fmt"
	"io"
	"net"
)

func main() {
	// Prompt user for email to create an email-formatted string.
	email := emails.EmailPrompt().String()

	// Attempt to connect to a TCP channel on a shared tcp port, unused by other processes, on the localhost IP address.
	channel, err := net.Dial("tcp", "127.0.0.1:1234")
	errorchecking.CheckError(err, "A")

	// Send email through TCP channel.
	fmt.Fprintf(channel, email)

	// Wait until a specific acknowledgement is sent from the host server through the TCP channel
	// and then end the process.
	for {
		// Receive an incoming message coming from the host server through the TCP channel.
		message, err := bufio.NewReader(channel).ReadString('\n')
		// Ensure that the reader stopped reading the string at the delimiter.
		if err != io.EOF {
			errorchecking.CheckError(err, "A")
		}


		if message == "roger that" {
			fmt.Println("Acknowledgement received from server. TCP client exiting...")
			return
		}
	}
}
