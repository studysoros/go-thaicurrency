# Go Thai Currency

A Go library to convert decimal numbers into Thai currency text (Baht and Satang).

## Features

- Converts `decimal.Decimal` to Thai Baht text.
- Handles millions and multiple grouping levels.
- Supports Satang (decimal) parts correctly.
- Proper handling of special cases like "เอ็ด" and "ยี่".

## Installation

```bash
go get github.com/studysoros/go-thaicurrency
```

## Usage

```go
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
			fmt.Printf("Error converting %v: %v\n", input, err)
			continue
		}
		fmt.Println(result)
	}
}
```

## How to Run Example

You can run the provided example in the `cmd` directory:

```bash
go run .\cmd\example\main.go
```
