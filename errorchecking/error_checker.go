package errorchecking

import (
	"fmt"
	"os"
)

func CheckError(error error, process string) {
	if error != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s in process %s\n", error.Error(), process)
		os.Exit(1)
	}
}