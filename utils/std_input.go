package utils

import (
	"bufio"
	"os"
	"runtime"
	"strings"
)

func ReadFromStdIn() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	// convert CRLF to LF
	if runtime.GOOS != "windows" {
		text = strings.Replace(text, "\n", "", -1)
	} else {
		text = strings.Replace(text, "\r\n", "", -1)
	}
	return text
}
