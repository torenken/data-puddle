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

	var billingAccounts []data.BillingAccount
	for i := 0; i < 1; i++ {
		account := data.BillingAccount{}
		err := faker.FakeData(&account)
		if err != nil {
			fmt.Println(err)
		}
		billingAccounts = append(billingAccounts, account)
	}

	indent, _ := json.MarshalIndent(&billingAccounts, "", "\t")
	fmt.Println(string(indent))

	since := time.Since(now)
	fmt.Printf("its tooks: %v\n", since)
}
