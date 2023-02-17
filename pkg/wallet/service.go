package wallet

import (
	"errors"

	"github.com/Jaborov-U/Day-13-14-Wallet/pkg/types"
)

var ErrAccountNotFound = errors.New("favorite not found")
var ErrPaymentNotFound = errors.New("payment not found")

type Service struct{
	nextAccountID int64
	accounts []*types.Account
	payments []*types.Payment
}

func (s *Service) FindAccountByID(accountID int64) (*types.Account, error) {

	for _, account := range s.accounts {
		if account.ID == accountID {
			return account, nil
		}
	}

	return nil, ErrAccountNotFound
}
