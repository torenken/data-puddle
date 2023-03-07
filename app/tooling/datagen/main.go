package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/torenken/data-puddle/app/tooling/datagen/data"
)

func main() {
	now := time.Now()
	num := 10

	var billingAccounts []data.BillingAccount
	accountChannel := make(chan data.BillingAccount)

	account := data.BillingAccount{}
	for i := 0; i < num; i++ {
		go func() {
			err := faker.FakeData(&account)
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

	/*	for i := 0; i < 1; i++ {
		err := faker.FakeData(&account)
		if err != nil {
			fmt.Println(err)
		}
		billingAccounts = append(billingAccounts, account)
	}*/

	indent, _ := json.MarshalIndent(&billingAccounts, "", "\t")
	fmt.Println(string(indent))

	since := time.Since(now)
	fmt.Printf("its tooks: %v\n", since)
}
