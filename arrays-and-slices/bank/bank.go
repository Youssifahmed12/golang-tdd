package bank

type Transaction struct {
	From string
	To   string
	Sum  float64
}

func NewTransaction(from, to Account, sum float64) Transaction {
	return Transaction{From: from.Name, To: to.Name, Sum: sum}
}

type Account struct {
	Name    string
	Balance float64
}

func NewBalanceFor(account Account, transactions []Transaction) Account {
	return Reduce(transactions, applyTransaction, account)
}

func applyTransaction(a Account, transaction Transaction) Account {
	if transaction.From == a.Name {
		a.Balance -= transaction.Sum
	}
	if transaction.To == a.Name {
		a.Balance += transaction.Sum
	}
	return a
}

func Reduce[A, B any](collection []A, f func(B, A) B, intialValue B) B {
	var result = intialValue
	for _, x := range collection {
		result = f(result, x)
	}
	return result
}

func Find[I any](collection []I, f func(item I) bool) (foundItem I, found bool) {
	for _, x := range collection {
		if f(x) {
			return x, true
		}
	}
	return
}
