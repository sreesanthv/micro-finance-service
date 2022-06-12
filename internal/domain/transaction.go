package domain

type Transaction struct {
	ID        int
	AccountId int
	Debit     float64
	Credit    float64
	Narration string
}
