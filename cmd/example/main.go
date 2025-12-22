package main

import (
	"fmt"
	"math"

	"github.com/shopspring/decimal"
	"github.com/studysoros/go-thaicurrency"
)

func main() {
	inputs := []decimal.Decimal{
		decimal.NewFromFloat(1234),
		decimal.NewFromFloat(33333.75),
		decimal.NewFromFloat(01),
		decimal.NewFromFloat(.01),
		decimal.NewFromFloat(.2),
		decimal.NewFromFloat(.02),
		decimal.NewFromFloat(.11),
		decimal.NewFromFloat(1234567891231.21),
		decimal.NewFromFloat(111.55),
		decimal.NewFromFloat(111.550),
		decimal.NewFromFloat(111.555),
		decimal.NewFromFloat(111.5),
		decimal.NewFromFloat(5.461),
		decimal.NewFromFloat(5.465),
		decimal.NewFromFloat(0),
		decimal.NewFromInt(math.MaxInt64), // 9223372036854775807

		decimal.RequireFromString("9223372036854775808"),
		decimal.RequireFromString("9223372036854775807.01"),
		decimal.NewFromInt(math.MaxInt64).Add(decimal.NewFromInt(1)),
		decimal.NewFromFloat(9223372036854775808),
	}
	for _, input := range inputs {
		fmt.Println(input)
		// convert decimal to thai text (baht) and print the result here
		result, err := thaicurrency.Decimal(input)
		if err != nil {
			fmt.Printf("Error converting %v: %v\n\n", input, err)
			continue
		}
		fmt.Printf("%v\n\n", result)
	}
}
