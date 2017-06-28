package main

import (
	"encoding/csv"
	"fmt"
	"github.com/karantin2020/cli"
	"github.com/karantin2020/csvparse"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"os"
	// "strings"

	"github.com/karantin2020/CCAccounting/types"
)

var (
	fname string
	encf  string
	enc   encoding.Encoding

	start, end int
)

func main() {
	flags := cli.New("Test ccbook struct parsing", "0.0.1")
	flags.StringVarP(&fname, "fname", "f", "", "input file to parse")
	flags.StringVarP(&encf, "encf", "e", "cp1251", "input file encoding, default: CP1251")

	flags.IntVarP(&start, "start", "a", 0, "start row to print after parsing, default: 0")
	flags.IntVarP(&end, "end", "z", 10, "end row to print after parsing, default: 10")

	flags.Parse()

	if start >= end {
		fmt.Println("Start needs to be less than end")
		os.Exit(1)
	}
	enc, _ = charset.Lookup(encf)

	re, err := os.Open(fname)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	r := transform.NewReader(re, enc.NewDecoder())
	rc := csv.NewReader(r)
	rc.Comma = ';'
	fl := types.CancBookV1List{}

	fmt.Println("Start reading", fname)

	parser := csvparse.Parser{}

	if err := parser.ReadAll(rc, &fl); err != nil {
		fmt.Printf("%#v\n", err)
		os.Exit(1)
	}

	if end > len(fl) {
		end = len(fl)
	}

	for i := start; i < end; i++ {
		fmt.Printf("%+v\n", fl[i])
	}
	for _, i := range parser.Errors {
		fmt.Printf("%#v\n", i)
	}
}
