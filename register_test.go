package gomoney

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnregisteredMoney(t *testing.T) {
	unregisterAllMoney()

	_, err := New(2, "btc")
	assert.Error(t, err)
}

func TestRegisterAlreadyRegisteredMoney(t *testing.T) {
	unregisterAllMoney()

	err := RegisterNewMoney("usd", 8)
	assert.NoError(t, err)

	err = RegisterNewMoney("Usd", 8)
	assert.Error(t, err)
}

// unregisterAllMoney for tests
func unregisterAllMoney() {
	money = map[string]int{}
}
