package bank

type Transaction struct {
	From string
	To   string
	Sum  float64
}

func BalanceFor(transactions []Transaction, name string) float64 {
	var balance float64
	for _, t := range transactions {
		if t.From == name {
			balance -= t.Sum
		}
		if t.To == name {
			balance += t.Sum
		}
	}
	return balance
}

func Reduce[A any](collection []A, f func(A, A) A, intialValue A) A {
	var result = intialValue
	for _, x := range collection {
		result = f(result, x)
	}
	return result
}
