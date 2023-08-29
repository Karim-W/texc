package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Replace(line, "\\n", "\n", -1)
		// replace all \t with a tab
		line = strings.Replace(line, "\\t", "\t", -1)
		// replace all \r with a carriage return
		line = strings.Replace(line, "\\r", "\r", -1)
		// replace all \f with a form feed
		line = strings.Replace(line, "\\f", "\f", -1)
		// replace all \b with a backspace
		line = strings.Replace(line, "\\b", "\b", -1)
		// replace all \a with a bell
		line = strings.Replace(line, "\\a", "\a", -1)
		// replace all \v with a vertical tab
		line = strings.Replace(line, "\\v", "\v", -1)
		// replace all \" with a double quote"
		line = strings.Replace(line, "\\\"", "\"", -1)
		fmt.Println(line)
	}
}
