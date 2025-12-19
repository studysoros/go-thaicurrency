package main

import (
	"fmt"

	"github.com/shopspring/decimal"
	"github.com/studysoros/go-thaicurrency"
)

func main() {
	inputs := []decimal.Decimal{
		decimal.NewFromFloat(1234),
		decimal.NewFromFloat(33333.75),
	}
	for _, input := range inputs {
		fmt.Println(input)
		// convert decimal to thai text (baht) and print the result here
		result, err := thaicurrency.Decimal(input)
		if err != nil {
			panic(err)
		}
		fmt.Println(result)
	}
}
