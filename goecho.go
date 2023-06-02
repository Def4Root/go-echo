package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	addNewline := flag.Bool("n", false, "出力文字の最後の改行を削除します。")
	escapeFlag := flag.Bool("e", false, "エスケープ文字を有効にします。")
	flag.Parse()

	var result []string

	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "$") {
			envVar := os.Getenv(strings.TrimPrefix(arg, "$"))
			result = append(result, envVar)
		} else {
			result = append(result, arg)
		}
	}

	if *escapeFlag {
		for i := range result {
			result[i] = strings.Replace(result[i], "\\n", "\n", -1)
		}
	}

	output := strings.Join(result, " ")
	output = strings.Replace(output, "-n", "", -1)
	output = strings.Replace(output, "-e", "", -1)

	if !*addNewline {
		fmt.Print(output)
	} else {
		fmt.Println(output)
	}
}
