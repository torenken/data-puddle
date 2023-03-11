package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/go-faker/faker/v4"
	"github.com/go-faker/faker/v4/pkg/options"
	"github.com/torenken/data-puddle/app/tooling/datagen/data"
)

var (
	n = flag.Int("n", 3, "")
)

var usage = `Usage: datagen [options...] <choose customer|account|agreement|...>

Options:
  -n  Number of entries to generate the records. Default is 3.
`

func main() {
	flag.Usage = func() { _, _ = fmt.Fprintf(os.Stderr, "%s\n", usage) }

	flag.Parse()
	if flag.NArg() < 1 {
		usageAndExit()
	}
	num := *n
	t := flag.Args()[0]
	switch t {
	case "customer":
		handle(num, data.Customer{})
	case "account":
		handle(num, data.BillingAccount{})
	case "agreement":
		handle(num, data.Agreement{})
	default:
		usageAndExit()
	}
}

type DataValue interface{}

func handle(num int, content DataValue) {
	var valueCol []interface{}
	valueCh := make(chan interface{})

	for i := 0; i < num; i++ {
		go func() {
			err := faker.FakeData(&content, options.WithRandomMapAndSliceMinSize(1), options.WithRandomMapAndSliceMaxSize(2))
			if err != nil {
				fmt.Println(err)
			}
			valueCh <- content
		}()
	}

	for i := 0; i < num; i++ {
		r := <-valueCh
		valueCol = append(valueCol, r)
	}

	if err := json.NewEncoder(os.Stdout).Encode(&valueCol); err != nil {
		log.Fatal(err)
	}
}

func usageAndExit() {
	flag.Usage()
	os.Exit(1)
}
