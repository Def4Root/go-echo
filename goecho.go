package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("")
		return
	}

	args := []string{}
	for i := 1; i < len(os.Args); i++ {
		if strings.Contains(os.Args[i], "$") {
			varArg := strings.Replace(os.Args[i], "$", "", -1)
			args = append(args, os.Getenv(varArg))
		} else {
			args = append(args, os.Args[i])
		}
	}

	for _, arg := range args {
		fmt.Print(arg, " ")
	}

	fmt.Println("")
}