package gomoney

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
	res, err := m.Add(m2)
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, big.NewInt(122520100), res.value)
	assert.Equal(t, "BTC", m.currency)
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
	res, err := m.Sub(m2)
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, big.NewInt(99000), res.value)
	assert.Equal(t, "BTC", m.currency)
}
