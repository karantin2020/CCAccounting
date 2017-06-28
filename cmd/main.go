package main

import (
	"fmt"
	cli "github.com/jawher/mow.cli"
	"github.com/karantin2020/csvparse"

	"encoding/csv"
	"github.com/karantin2020/CCAccounting/types"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"os"
)

var (
	encf  *string
	fname *string
	enc   encoding.Encoding

	fl     = types.CancBookV1List{}
	parser = csvparse.Parser{}
)

var (
	VERSION = "0.0.1"
)

func main() {
	app := cli.App("accountant", "Criminal case accounting")
	app.Spec = "[-t] [-a] [-u] [-p] [-d] [-n] [-c] [-f] [-e]"
	app.Version("v version", "accountant v"+VERSION)

	setOptions(app)

	fname = app.String(cli.StringOpt{
		Name:   "f filename",
		Value:  "",
		Desc:   "input file to parse at startup",
		EnvVar: "CCA_FILE FILE",
	})
	encf = app.String(cli.StringOpt{
		Name:   "e encoding",
		Value:  "cp1251",
		Desc:   "input file encoding, default: CP1251",
		EnvVar: "CCA_ENCODING ENCODING",
	})

	app.Action = func() {
		if *fname != "" && *db_type == "mongodb" {
			err := FirstInitMongo()
			if err != nil {
				fmt.Println(err)
				return
			}

			parseFile()
			insertRows()
		} else {
			fmt.Println("fname option must not be empty")
			app.PrintLongHelp()
			cli.Exit(1)
		}
	}

	if session != nil {
		defer session.Close()
	}
	app.Run(os.Args)

}

func parseFile() {
	fmt.Println("Start parseFile " + *fname)
	enc, _ = charset.Lookup(*encf)

	re, err := os.Open(*fname)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	r := transform.NewReader(re, enc.NewDecoder())
	rc := csv.NewReader(r)
	rc.Comma = ';'

	if err := parser.ReadAll(rc, &fl); err != nil {
		fmt.Printf("%#v\n", err)
		os.Exit(1)
	}

	for _, i := range parser.Errors {
		fmt.Printf("%#v\n", i)
		os.Exit(1)
	}
	fmt.Printf("Parsed %d rows of csv\n", len(fl))
}

func insertRows() {
	var reslst []interface{}
	for i := 0; i < len(fl); i++ {
		// fmt.Printf("%#v\n", i)
		reslst = append(reslst, fl[i])
	}

	c := session.DB(*dbname).C(*collection)
	err := c.Insert(reslst...)
	if err != nil {
		fmt.Println("Err in insert", err)
	}
	fmt.Printf("Wrote %d rows to db\n", len(reslst))
}
