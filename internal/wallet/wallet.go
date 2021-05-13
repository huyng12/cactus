package wallet

type Wallet struct {
	balance      int
	transactions []Tx
}

type Walleter interface {
	GetBalance() int
	GetTx() []Tx
	DoTx(tx Tx)
}

func New() Walleter {
	return &Wallet{balance: 0}
}

func (w *Wallet) GetBalance() int {
	return w.balance
}

func (w *Wallet) GetTx() []Tx {
	return w.transactions
}

func (w *Wallet) DoTx(tx Tx) {
	w.balance += tx.Amount
	w.transactions = append(w.transactions, tx)
}
