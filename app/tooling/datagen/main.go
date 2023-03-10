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

var usage = `Usage: datagen [options...] <choose account|...>

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
	case "account":
		handleAccount(num)
	default:
		usageAndExit()
	}
}

func handleAccount(num int) {
	var billingAccounts []data.BillingAccount
	accountChannel := make(chan data.BillingAccount)

	account := data.BillingAccount{}
	for i := 0; i < num; i++ {
		go func() {
			err := faker.FakeData(&account, options.WithRandomMapAndSliceMinSize(1), options.WithRandomMapAndSliceMaxSize(2))
			if err != nil {
				fmt.Println(err)
			}
			accountChannel <- account
		}()
	}

	for i := 0; i < num; i++ {
		r := <-accountChannel
		billingAccounts = append(billingAccounts, r)
	}

	if err := json.NewEncoder(os.Stdout).Encode(&billingAccounts); err != nil {
		log.Fatal(err)
	}
}

func usageAndExit() {
	flag.Usage()
	os.Exit(1)
}
