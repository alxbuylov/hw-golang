package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	from, to      string
	limit, offset int64
)

func init() {
	flag.StringVar(&from, "from", "", "file to read from")
	flag.StringVar(&to, "to", "", "file to write to")
	flag.Int64Var(&offset, "offset", 0, "offset in input file")
	flag.Int64Var(&limit, "limit", 0, "limit of bytes to copy")
}

func main() {
	flag.Parse()

	if err := Copy(from, to, offset, limit); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Done")
}
