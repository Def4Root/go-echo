package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	var result []string
	addNewline := true

	for _, arg := range os.Args[1:] {
		switch {
		case strings.HasPrefix(arg, "$"):
			envVar := os.Getenv(strings.TrimPrefix(arg, "$"))
			result = append(result, envVar)
		case arg == "-n":
			addNewline = false
		default:
			result = append(result, arg)
		}
	}

	output := strings.Join(result, " ")
	output = strings.ReplaceAll(output, "\\n", "\n")

	if addNewline {
		fmt.Print(output, "\n")
	} else {
		fmt.Print(output)
	}
}
