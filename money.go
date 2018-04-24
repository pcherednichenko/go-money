package money

import (
	"fmt"
	"math/big"
	"strings"
)

var (
	tenInt = big.NewInt(10)
)

type Money struct {
	value    *big.Int
	currency string
}

func (m Money) Add(m2 Money) Money {
	if m.currency != m2.currency {
		panic(fmt.Sprintf(
			"try to sum currency '%s' with different currency '%s'",
			m.currency, m2.currency),
		)
	}
	res := new(big.Int).Add(m.value, m2.value)
	return Money{
		value:    res,
		currency: m.currency,
	}
}

func (m Money) Sub(m2 Money) Money {
	if m.currency != m2.currency {
		panic(fmt.Sprintf(
			"try to sub currency '%s' with different currency '%s'",
			m.currency, m2.currency),
		)
	}
	res := new(big.Int).Sub(m.value, m2.value)
	return Money{
		value:    res,
		currency: m.currency,
	}
}

func (m Money) Mul(m2 Money) Money {
	if m.currency != m2.currency {
		panic(fmt.Sprintf(
			"try to mul currency '%s' with different currency '%s'",
			m.currency, m2.currency),
		)
	}
	p, err := moneyPrecision(m)
	if err != nil {
		panic(err)
	}
	d3Full := new(big.Int).Mul(m.value, m2.value)
	mul := new(big.Int).Exp(tenInt, big.NewInt(int64(p)), nil)
	d3Result := new(big.Int).Quo(d3Full, mul)
	return Money{
		value:    d3Result,
		currency: m.currency,
	}
}

func (m Money) MulFloat(f float64) Money {
	if !currencyRegistered(m.currency) {
		panic(fmt.Sprintf("currency with name %s not registered, "+errNeedToCreateBefore, m.currency))
	}
	m2, err := NewFromFloat(f, m.currency)
	if err != nil {
		panic(err)
	}
	res := m.Mul(m2)
	return res
}

func (m Money) Div(m2 Money) Money {
	if m.currency != m2.currency {
		panic(fmt.Sprintf(
			"try to div currency '%s' with different currency '%s'",
			m.currency, m2.currency),
		)
	}
	p, err := moneyPrecision(m)
	if err != nil {
		panic(err)
	}
	if m2.value.Sign() == 0 {
		panic("money division by 0")
	}

	mul := new(big.Int).Exp(tenInt, big.NewInt(int64(p)), nil)
	bigM := new(big.Int).Mul(m.value, mul)
	newVal := new(big.Int).Quo(bigM, m2.value)

	return Money{
		value:    newVal,
		currency: m.currency,
	}
}

func (m Money) DivFloat(f float64) Money {
	if !currencyRegistered(m.currency) {
		panic(fmt.Sprintf("currency with name %s not registered, "+errNeedToCreateBefore, m.currency))
	}
	m2, err := NewFromFloat(f, m.currency)
	if err != nil {
		panic(err)
	}
	res := m.Div(m2)
	return res
}

func (m Money) Float64() float64 {
	r, err := m.Rat()
	if err != nil {
		panic(err)
	}
	result, _ := r.Float64()
	return result
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

func (m Money) String() string {
	p, err := moneyPrecision(m)
	if err != nil {
		panic(err)
	}
	abs := new(big.Int).Abs(m.value)
	str := abs.String()
	var result string
	if len(str) < p {
		result = "0."
		result += strings.Repeat("0", p-len(str))
		result += strings.TrimRight(str, "0")
	} else {
		before := str[:len(str)-p]
		if len(before) != 0 {
			result = before
		} else {
			result = "0"
		}
		afterPoint := strings.TrimRight(str[len(str)-p:], "0")
		if len(afterPoint) > 0 {
			result += "." + afterPoint
		}
	}

	if m.value.Sign() < 0 {
		return "-" + result
	}
	return result
}

// Equal returns whether the numbers represented by m and m2 are equal
func (m Money) Equal(m2 Money) bool {
	return m.Cmp(m2) == 0
}

// Equals is deprecated, please use Equal method instead
func (m Money) Equals(m2 Money) bool {
	return m.Equal(m2)
}

// GreaterThan (GT) returns true when m is greater than m2
func (m Money) GreaterThan(m2 Money) bool {
	return m.Cmp(m2) == 1
}

// GreaterThanOrEqual (GTE) returns true when m is greater than or equal to m2
func (m Money) GreaterThanOrEqual(m2 Money) bool {
	cmp := m.Cmp(m2)
	return cmp == 1 || cmp == 0
}

// LessThan (LT) returns true when m is less than m2
func (m Money) LessThan(m2 Money) bool {
	return m.Cmp(m2) == -1
}

// LessThanOrEqual (LTE) returns true when m is less than or equal to m2
func (m Money) LessThanOrEqual(m2 Money) bool {
	cmp := m.Cmp(m2)
	return cmp == -1 || cmp == 0
}

// Cmp compares the numbers represented by m and m2 and returns:
//
//     -1 if m <  d2
//      0 if m == d2
//     +1 if m >  d2
//
func (m Money) Cmp(m2 Money) int {
	if m.currency != m2.currency {
		panic(fmt.Sprintf("try to compare moneys with different currencies: %s with %s", m.currency, m2.currency))
	}
	return m.value.Cmp(m2.value)
}
