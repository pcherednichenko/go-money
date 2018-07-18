# Go-money

Package go-money you need to trade currency with the appropriate precision

## Example of use:
```go
package money_test

import (
	"fmt"
	
	"github.com/pcherednichenko/go-money"
)

func ExampleMoney() {
	err := money.RegisterNewMoney("btc", 8)
	if err != nil {
		panic(err)
	}

	m1, err := money.NewFromFloat(2.00953, "btc")
	if err != nil {
		panic(err)
	}

	m2, err := money.NewFromString("5.23", "btc")
	if err != nil {
		panic(err)
	}

	result := m1.Add(m2)
	fmt.Println(result.String())
	fmt.Println(m2.GreaterThan(m1))
	// Output:
	// 7.23953
	// true
}

```
