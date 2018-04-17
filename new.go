package gomoney

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
)

func New(value int64, currency string) (Money, error) {
	if !currencyRegistered(currency) {
		return Money{}, fmt.Errorf("currency with name %s not registered, "+errNeedToCreateBefore, currency)
	}

	return Money{
		value:    big.NewInt(value),
		currency: strings.ToUpper(currency),
	}, nil
}

func NewFromInt(value int64, currency string) (Money, error) {
	p, err := currencyPrecision(currency)
	if err != nil {
		return Money{}, err
	}

	return Money{
		value:    big.NewInt(value * int64(powInt(10, p))),
		currency: strings.ToUpper(currency),
	}, nil
}

func NewFromString(value string, currency string) (Money, error) {
	p, err := currencyPrecision(currency)
	if err != nil {
		return Money{}, err
	}

	var intString string
	parts := strings.Split(value, ".")
	if len(parts) == 1 {
		// There is no decimal point, we can just parse the original string as
		// an int
		intString = value + strings.Repeat("0", p)
	} else if len(parts) == 2 {
		decimalPart := parts[1]
		count := len(decimalPart)
		if count < p {
			decimalPart += strings.Repeat("0", p-count)
		}
		intString = parts[0] + decimalPart[:p]
	} else {
		return Money{}, fmt.Errorf("can't convert %s to decimal: too many .s", value)
	}

	dValue := new(big.Int)
	_, ok := dValue.SetString(intString, 10)
	if !ok {
		return Money{}, fmt.Errorf("can't convert %s to decimal", value)
	}

	return Money{
		value:    dValue,
		currency: strings.ToUpper(currency),
	}, nil
}

func NewFromFloat(value float64, currency string) (Money, error) {
	p, err := currencyPrecision(currency)
	if err != nil {
		return Money{}, err
	}
	floor := math.Floor(value)

	// fast path, where float is an int
	if floor == value && value <= math.MaxInt64 && value >= math.MinInt64 {
		return NewFromInt(int64(value), currency)
	}

	// slow path: float is a decimal
	str := strconv.FormatFloat(value, 'f', p, 64)
	mon, err := NewFromString(str, currency)
	if err != nil {
		return Money{}, err
	}
	return mon, err
}

func powInt(number, power int) int {
	result := number
	for i := 1; i < power; i++ {
		result = result * number
	}
	return result
}
