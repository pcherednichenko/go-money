package money

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddMoney(t *testing.T) {
	unregisterAllMoney()

	err := RegisterNewMoney("btc", 8)
	if err != nil {
		t.Fail()
	}
	m, err := NewFromFloat(1.222, "btc")
	if err != nil {
		t.Fail()
	}

	m2, err := NewFromFloat(0.003201, "bTc")
	if err != nil {
		t.Fail()
	}
	res := m.Add(m2)
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, big.NewInt(122520100), res.value)
	assert.Equal(t, "BTC", res.currency)
}

func TestSubMoney(t *testing.T) {
	unregisterAllMoney()

	err := RegisterNewMoney("btc", 8)
	if err != nil {
		t.Fail()
	}
	m, err := NewFromFloat(0.001, "btc")
	if err != nil {
		t.Fail()
	}

	m2, err := NewFromFloat(0.00001, "bTc")
	if err != nil {
		t.Fail()
	}
	res := m.Sub(m2)
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, big.NewInt(99000), res.value)
	assert.Equal(t, "BTC", res.currency)
}

func TestMulRoundMoney(t *testing.T) {
	unregisterAllMoney()

	err := RegisterNewMoney("btc", 2)
	if err != nil {
		t.Fail()
	}
	m, err := NewFromFloat(20, "btc")
	if err != nil {
		t.Fail()
	}

	m2, err := NewFromFloat(6, "bTc")
	if err != nil {
		t.Fail()
	}
	res := m.Mul(m2)
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, big.NewInt(12000), res.value)
	assert.Equal(t, "BTC", res.currency)

	s := res.String()
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, "120", s)
}

func TestMulDecimalMoney(t *testing.T) {
	unregisterAllMoney()

	err := RegisterNewMoney("btc", 8)
	if err != nil {
		t.Fail()
	}
	m, err := NewFromFloat(40.05, "btc")
	if err != nil {
		t.Fail()
	}

	m2, err := NewFromFloat(0.1, "bTc")
	if err != nil {
		t.Fail()
	}
	res := m.Mul(m2)
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, big.NewInt(400500000), res.value)
	assert.Equal(t, "BTC", res.currency)

	s := res.String()
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, "4.005", s)
}

func TestMoneyToFloat(t *testing.T) {
	unregisterAllMoney()

	err := RegisterNewMoney("btc", 8)
	if err != nil {
		t.Fail()
	}
	m, err := NewFromFloat(0.001, "btc")
	if err != nil {
		t.Fail()
	}

	f := m.Float64()
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, 0.001, f)
}

func TestDecimalMoneyToString(t *testing.T) {
	unregisterAllMoney()

	err := RegisterNewMoney("btc", 8)
	if err != nil {
		t.Fail()
	}
	m, err := NewFromFloat(0.001, "btc")
	if err != nil {
		t.Fail()
	}

	s := m.String()
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, "0.001", s)
}

func TestBigMoneyToString(t *testing.T) {
	unregisterAllMoney()

	err := RegisterNewMoney("btc", 8)
	if err != nil {
		t.Fail()
	}
	m, err := NewFromFloat(230.001, "btc")
	if err != nil {
		t.Fail()
	}

	s := m.String()
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, "230.001", s)
}

func TestRoundMoneyToString(t *testing.T) {
	unregisterAllMoney()

	err := RegisterNewMoney("btc", 8)
	if err != nil {
		t.Fail()
	}
	m, err := NewFromFloat(530, "btc")
	if err != nil {
		t.Fail()
	}

	s := m.String()
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, "530", s)
}

func TestNegativeMoneyToString(t *testing.T) {
	unregisterAllMoney()

	err := RegisterNewMoney("btc", 8)
	if err != nil {
		t.Fail()
	}
	m, err := NewFromFloat(-1.22, "btc")
	if err != nil {
		t.Fail()
	}

	s := m.String()
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, "-1.22", s)
}

func TestCompareTwoMoney(t *testing.T) {
	unregisterAllMoney()

	err := RegisterNewMoney("btc", 8)
	if err != nil {
		t.Fail()
	}

	m, err := NewFromFloat(12.434, "btc")
	if err != nil {
		t.Fail()
	}

	m2, err := NewFromFloat(14.12, "btc")
	if err != nil {
		t.Fail()
	}

	result := m.GreaterThanOrEqual(m2)
	assert.False(t, result)

	result = m.LessThan(m2)
	assert.True(t, result)
}