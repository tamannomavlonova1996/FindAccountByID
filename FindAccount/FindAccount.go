package FindAccount

import (
	"FindAccountById/types"
	"errors"
)

func FindAccountByID(accountID int64) (*types.Account, error) {
	acc := []types.Account{
		{
			ID: 1, Phone: "12345", Balance: 12,
		},
		{
			ID: 2, Phone: "654321", Balance: 21,
		},
	}

	for _, v := range acc {
		if v.ID == accountID {
			return &v, nil
		}
	}

	return nil, errors.New("dsds")
}
