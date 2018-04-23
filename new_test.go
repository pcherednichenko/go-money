package money

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMoney(t *testing.T) {
	unregisterAllMoney()

	err := RegisterNewMoney("btc", 8)
	if err != nil {
		t.Fail()
	}
	m, err := NewFromInt(2, "btc")
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, big.NewInt(200000000), m.value)
	assert.Equal(t, "BTC", m.currency)
}

func TestNewBigMoney(t *testing.T) {
	unregisterAllMoney()

	err := RegisterNewMoney("ltc", 10)
	if err != nil {
		t.Fail()
	}
	m, err := NewFromInt(3424, "ltc")
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, big.NewInt(34240000000000), m.value)
	assert.Equal(t, "LTC", m.currency)
}

func TestNewFromString(t *testing.T) {
	unregisterAllMoney()

	err := RegisterNewMoney("btc", 8)
	if err != nil {
		t.Fail()
	}
	m, err := NewFromString("32.111122223333", "btc")
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, big.NewInt(3211112222), m.value)
	assert.Equal(t, "BTC", m.currency)
}

func TestNewFromRoundString(t *testing.T) {
	unregisterAllMoney()

	err := RegisterNewMoney("EUR", 2)
	if err != nil {
		t.Fail()
	}
	m, err := NewFromString("932", "eur")
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, big.NewInt(93200), m.value)
	assert.Equal(t, "EUR", m.currency)
}

func TestNewFromBigString(t *testing.T) {
	unregisterAllMoney()

	err := RegisterNewMoney("btc", 8)
	if err != nil {
		t.Fail()
	}
	m, err := NewFromString("99999932.111122223333", "btc")
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, big.NewInt(9999993211112222), m.value)
	assert.Equal(t, "BTC", m.currency)
}

func TestNewFromFloat(t *testing.T) {
	unregisterAllMoney()

	err := RegisterNewMoney("eUr", 2)
	if err != nil {
		t.Fail()
	}
	m, err := NewFromFloat(234.424, "Eur")
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, big.NewInt(23442), m.value)
	assert.Equal(t, "EUR", m.currency)
}

func TestNewFromLongFloat(t *testing.T) {
	unregisterAllMoney()

	err := RegisterNewMoney("eUr", 6)
	if err != nil {
		t.Fail()
	}
	m, err := NewFromFloat(111.222, "Eur")
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, big.NewInt(111222000), m.value)
	assert.Equal(t, "EUR", m.currency)
}
