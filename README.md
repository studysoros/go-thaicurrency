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

## Limitations

### Float64 Precision with Large Numbers

When using `decimal.NewFromFloat()` with very large numbers, you may encounter precision loss due to the inherent limitations of `float64` (which has ~15-17 significant decimal digits of precision).

**Example of precision loss:**

```go
// ❌ The .01 will be lost due to float64 precision limits
decimal.NewFromFloat(9223372036854775000.01)
// Actually becomes: 9223372036854776000 (no decimal part!)

// ✅ Use RequireFromString for exact precision
decimal.RequireFromString("9223372036854775000.01")
// Correctly preserves: 9223372036854775000.01
```

**Recommendation:** For amounts exceeding 15 significant digits or when precision is critical, always use `decimal.RequireFromString()` instead of `decimal.NewFromFloat()`.
