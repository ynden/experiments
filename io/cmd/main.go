package main

import (
	"fmt"
	"os"

	"github.com/ynden/experiments/io/stream"
)

func main() {
	path := os.Args[1]
	fmt.Println(stream.Read(4, path))
}
