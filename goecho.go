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
		case arg == "-e":
			replaced := strings.ReplaceAll(arg, "\\n", "\n")
			result = append(result, replaced)
		default:
			result = append(result, arg)
		}
	}

	output := strings.Join(result, " ")

	if addNewline {
		fmt.Print(output, "\n")
	} else {
		fmt.Print(output)
	}
}
