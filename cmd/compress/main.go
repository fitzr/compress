package main

import (
	"flag"
	"github.com/fitzr/compress"
	"github.com/pkg/profile"
	"log"
	"os"
)

func main() {

	defer profile.Start().Stop()

	c := flag.Bool("c", false, "compress")
	d := flag.Bool("d", false, "decompress")
	i := flag.String("i", "", "input file path")
	o := flag.String("o", "", "output file path")
	flag.Parse()

	if *c == *d || *i == "" || *o == "" {
		invalidArguments("invalid argment")
	}

	in, err := os.Open(*i)
	if err != nil {
		invalidArguments(err.Error())
	}
	defer in.Close()

	out, err := os.Create(*o)
	if err != nil {
		invalidArguments(err.Error())
	}
	defer out.Close()

	if *c {
		compress.Compress(in, out)
	} else {
		compress.Decompress(in, out)
	}
}

func invalidArguments(err string) {
	log.Println(err)
	flag.PrintDefaults()
	os.Exit(1)
}
