package emails

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

var scanner = bufio.NewScanner(os.Stdin)

// Prompts a user for a To, From, Title, and Content strings through the command line to create an email object.
func EmailPrompt() Email {
	fmt.Println("Please input each email field on a separate line down below:")

	to := promptUser("To: ")
	from := promptUser("From: ")
	date := time.Now().String()
	title := promptUser("Title: ")
	content := promptUser("Content: ")

	email := Email{to, from, date, title, content}
	return email
}

func promptUser(prompt string) string {
	fmt.Print(prompt)
	scanner.Scan()
	return scanner.Text()
}

type Email struct {
	To string
	From string
	Date string
	Title string
	Content string
}

func (email Email) String() string {
	string := fmt.Sprintf("To: %s\nFrom: %s\nDate: %s\nTitle: %s\nContent: %s\t", email.To, email.From, email.Date, email.Title, email.Content)
	return string
}
