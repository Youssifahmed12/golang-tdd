package bank

type Transaction struct {
	From string
	To   string
	Sum  float64
}

func BalanceFor(transactions []Transaction, name string) float64 {
	adjustBalance := func(currentBalance float64, t Transaction) float64 {
		if t.From == name {
			return currentBalance - t.Sum
		}
		if t.To == name {
			return currentBalance + t.Sum
		}
		return currentBalance
	}
	return Reduce(transactions, adjustBalance, 0.0)
}

func Reduce[A, B any](collection []A, f func(B, A) B, intialValue B) B {
	var result = intialValue
	for _, x := range collection {
		result = f(result, x)
	}
	return result
}
