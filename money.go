package money

import (
	"encoding/binary"
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

func (m Money) Currency() string {
	return m.currency
}

func (m Money) Add(m2 Money) Money {
	if m.currency != m2.currency {
		panic(fmt.Sprintf(
			"try to sum currency '%s' with different currency '%s'",
			m.currency, m2.currency),
		)
	}
	return Money{
		value:    new(big.Int).Add(m.value, m2.value),
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
	return Money{
		value:    new(big.Int).Sub(m.value, m2.value),
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
	return Money{
		value:    new(big.Int).Quo(d3Full, mul),
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
	return m.Mul(m2)
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

	return Money{
		value:    new(big.Int).Quo(bigM, m2.value),
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
	return m.Div(m2)
}

// Neg returns -d.
func (m Money) Neg() Money {
	val := new(big.Int).Neg(m.value)
	return Money{
		value:    val,
		currency: m.currency,
	}
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

func (m Money) IsZero() bool {
	return m.value.Sign() == 0
}

func (m Money) GreaterThanZero() bool {
	return m.value.Cmp(big.NewInt(0)) == 1
}

func (m Money) GreaterThanOrEqualsZero() bool {
	cmp := m.value.Cmp(big.NewInt(0))
	return cmp == 1 || cmp == 0
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

// MarshalJSON implements the json.Marshaler interface.
func (m Money) MarshalJSON() ([]byte, error) {
	return []byte("\"" + m.String() + "\""), nil
}

// MarshalText implements the encoding.TextMarshaler interface for XML
// serialization.
func (m Money) MarshalText() (text []byte, err error) {
	return []byte(m.String()), nil
}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
func (m Money) MarshalBinary() (data []byte, err error) {
	p, err := moneyPrecision(m)
	if err != nil {
		return nil, err
	}
	// Write the exponent first since it's a fixed size
	v1 := make([]byte, 4)
	binary.BigEndian.PutUint32(v1, uint32(p))

	// Add the value
	var v2 []byte
	if v2, err = m.value.GobEncode(); err != nil {
		return
	}

	// Return the byte array
	data = append(v1, v2...)
	return
}

// GobEncode implements the gob.GobEncoder interface for gob serialization.
func (m Money) GobEncode() ([]byte, error) {
	return m.MarshalBinary()
}
