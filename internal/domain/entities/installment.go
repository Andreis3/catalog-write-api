package entities

type Installment struct {
	ID      int64
	OrderID int64
	Count   int64
	Price   float64
}
