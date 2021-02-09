/*Sample input:
	Bob
	Sally
	Hello
	World*/

package emails

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// Prompts a user for a To, From, Title, and Content strings through the command line to create an email object.
func EmailPrompt() Email {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please input each email field on a separate line down below:")

	fmt.Print("To: ")
	scanner.Scan()
	to := scanner.Text()

	fmt.Print("From: ")
	scanner.Scan()
	from := scanner.Text()

	/*fmt.Print("Date: ")
	scanner.Scan()
	date := scanner.Text()*/
	date := time.Now().String()

	fmt.Print("Title: ")
	scanner.Scan()
	title := scanner.Text()

	fmt.Print("Content: ")
	scanner.Scan()
	content := scanner.Text()

	email := Email{to, from, date, title, content}
	return email
}
// Generic email type as specified in assignment.
type Email struct {
	To string
	From string
	Date string
	Title string
	Content string
}

func (email *Email) String() string {
	string := fmt.Sprintf("To: %s\nFrom: %s\nDate: %s\nTitle: %s\nContent: %s\t", email.To, email.From, email.Date, email.Title, email.Content)
	return string
}