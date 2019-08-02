package m_account

import (
	"math/big"
)

type Account struct {
	model.Base
	MemberId   uint
	CurrencyId uint
	Balance    big.decimal
	Locked     big.decimal
}
