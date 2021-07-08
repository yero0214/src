package main

import (
	"bufio"
	"os"
	"strings"
)

func input() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}
