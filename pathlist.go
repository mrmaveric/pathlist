package main

import (
	"fmt"
	"os"
	"strings"
	"runtime"
)

func main() {
	var sep string;

	if runtime.GOOS == "windows" {
		sep = ";"
	} else {
		sep = ":"
	}

	for _, entry := range strings.Split(os.Getenv("PATH"), sep) {
		fmt.Println(entry)
	}
}
