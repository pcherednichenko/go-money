package gomoney

import (
	"fmt"
	"math/big"
	"strings"
)

var (
	zeroInt   = big.NewInt(0)
	oneInt    = big.NewInt(1)
	twoInt    = big.NewInt(2)
	fourInt   = big.NewInt(4)
	fiveInt   = big.NewInt(5)
	tenInt    = big.NewInt(10)
	twentyInt = big.NewInt(20)
)

type Money struct {
	value    *big.Int
	currency string
}

func (m Money) Add(m2 Money) (Money, error) {
	if m.currency != m2.currency {
		return Money{}, fmt.Errorf(
			"try to sum currency '%s' with different currency '%s'",
			m.currency, m2.currency)
	}
	res := new(big.Int).Add(m.value, m2.value)
	return Money{
		value:    res,
		currency: m.currency,
	}, nil
}

func (m Money) Sub(m2 Money) (Money, error) {
	if m.currency != m2.currency {
		return Money{}, fmt.Errorf(
			"try to sub currency '%s' with different currency '%s'",
			m.currency, m2.currency)
	}
	res := new(big.Int).Sub(m.value, m2.value)
	return Money{
		value:    res,
		currency: m.currency,
	}, nil
}

func (m Money) Float64() (float64, error) {
	r, err := m.Rat()
	if err != nil {
		return 0.0, err
	}
	result, _ := r.Float64()
	return result, nil
}

// Rat returns a rational number representation of money
func (m Money) Rat() (*big.Rat, error) {
	p, err := moneyPrecision(m)
	if err != nil {
		return nil, err
	}
	mul := new(big.Int).Exp(tenInt, big.NewInt(int64(p)), nil)
	return new(big.Rat).SetFrac(m.value, mul), nil
}

func (m Money) String() (string, error) {
	p, err := moneyPrecision(m)
	if err != nil {
		return "", err
	}
	abs := new(big.Int).Abs(m.value)
	str := abs.String()
	var result string
	if len(str) < p {
		result = "0."
		result += strings.Repeat("0", p-len(str))
		result += strings.TrimRight(str, "0")
	} else {
		result = str[:len(str)-p]
		afterPoint := strings.TrimRight(str[len(str)-p:], "0")
		if len(afterPoint) > 0 {
			result += "." + afterPoint
		}
	}

	if m.value.Sign() < 0 {
		return "-" + result, nil
	}
	return result, nil
}

func (m Money) Scan() (string, error) {
	p, err := moneyPrecision(m)
	if err != nil {
		return "", err
	}
	abs := new(big.Int).Abs(m.value)
	str := abs.String()
	var result string
	if len(str) < p {
		result = "0."
		result += strings.Repeat("0", p-len(str))
		result += strings.TrimRight(str, "0")
	} else {
		result = str[:len(str)-p]
		afterPoint := strings.TrimRight(str[len(str)-p:], "0")
		if len(afterPoint) > 0 {
			result += "." + afterPoint
		}
	}
	return result, nil
}
