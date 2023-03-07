package main

import (
	"encoding/json"
	"fmt"

	"github.com/go-faker/faker/v4"
	"github.com/torenken/data-puddle/app/tooling/datagen/data"
)

func main() {
	a := data.BillingAccount{}
	err := faker.FakeData(&a)
	if err != nil {
		fmt.Println(err)
	}
	indent, _ := json.MarshalIndent(&a, "", "\t")
	fmt.Printf(string(indent))
}
