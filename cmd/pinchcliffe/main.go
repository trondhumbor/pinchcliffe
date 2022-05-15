package main

import (
	"flag"

	"github.com/trondhumbor/pinchcliffe/internal/pinchcliffe"
)

func main() {
	bin := flag.String("bin", "", "Path to binary asset file")
	out := flag.String("out", "", "Path to output folder")

	flag.Parse()

	if *bin == "" || *out == "" {
		flag.Usage()
		return
	}

	pinchcliffe.ExtractArchive(*bin, *out)
}
