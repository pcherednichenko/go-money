package gomoney

import (
	"fmt"
	"strings"
)

var money = map[string]int{}

// RegisterNewMoney before you can use it
func RegisterNewMoney(currency string, precision int) error {
	upperCur := strings.ToUpper(currency)
	if _, exist := money[upperCur]; exist {
		return fmt.Errorf("money with currency %s already exist", upperCur)
	}
	money[upperCur] = precision
	return nil
}

func currencyPrecision(currency string) (int, error) {
	p, ok := money[strings.ToUpper(currency)]
	if !ok {
		return 0, fmt.Errorf("currency with name '%s' not registered, "+errNeedToCreateBefore, currency)
	}
	return p, nil
}

func currencyRegistered(currency string) bool {
	if _, ok := money[strings.ToUpper(currency)]; ok {
		return true
	}
	return false
}

// unregisterAllMoney for tests
func unregisterAllMoney() {
	money = map[string]int{}
}
