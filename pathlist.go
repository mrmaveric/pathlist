/*
Package pathlist is a cli tool for listing your PATH environment variable
with one entry per line.

pathlist takes no arguements.
*/
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	for _, entry := range strings.Split(os.Getenv("PATH"), sep) {
		fmt.Println(entry)
	}
}
