package gomoney

import (
	"fmt"
	"math/big"
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
	res := m.value.Add(m.value, m2.value)
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
	res := m.value.Sub(m.value, m2.value)
	return Money{
		value:    res,
		currency: m.currency,
	}, nil
}
