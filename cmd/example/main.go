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
		decimal.NewFromFloat(01),
		decimal.NewFromFloat(.01),
		decimal.NewFromFloat(.11),
		decimal.NewFromFloat(1234567891231.21),
		decimal.NewFromFloat(111.55),
		decimal.NewFromFloat(111.550),
		decimal.NewFromFloat(111.555),
		decimal.NewFromFloat(111.5),
	}
	for _, input := range inputs {
		fmt.Println(input)
		// convert decimal to thai text (baht) and print the result here
		result, err := thaicurrency.Decimal(input)
		if err != nil {
			fmt.Printf("Error converting %v: %v\n", input, err)
			continue
		}
		fmt.Println(result)
	}
}
