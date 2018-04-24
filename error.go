package money

import (
	"errors"
)

var (
	errNeedToCreateBefore = "you need to call money.RegisterNewMoney() before using a new currency"
	errEmptyCurrency      = errors.New("money has empty currency")
)
