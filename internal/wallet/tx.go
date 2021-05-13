package wallet

import "time"

type Tx struct {
	Amount    int
	Purpose   string
	CreatedAt time.Time
}

func MakeTx(amount int, purpose string) *Tx {
	return &Tx{
		Amount:    amount,
		Purpose:   purpose,
		CreatedAt: time.Now(),
	}
}
