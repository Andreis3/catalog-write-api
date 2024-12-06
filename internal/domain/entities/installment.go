package entities

import "github.com/andreis3/catalog-write-api/internal/domain/commons"

type Installment struct {
	id      int64
	orderID int64
	count   int
	price   float64
	commons.EntityErrors
	commons.ValidateFields
}

func InstallmentBuilder() *Installment {
	return &Installment{}
}

func (i *Installment) GetID() int64 {
	return i.id
}

func (i *Installment) GetOrderID() int64 {
	return i.orderID
}

func (i *Installment) GetCount() int {
	return i.count
}

func (i *Installment) GetPrice() float64 {
	return i.price
}

func (i *Installment) SetID(id int64) *Installment {
	i.id = id
	return i
}

func (i *Installment) SetOrderID(orderID int64) *Installment {
	i.orderID = orderID
	return i
}

func (i *Installment) SetCount(count int) *Installment {
	i.count = count
	return i
}

func (i *Installment) SetPrice(price float64) *Installment {
	i.price = price
	return i
}

func (i *Installment) Build() *Installment {
	return i
}

func (i *Installment) Validate() *commons.EntityErrors {
	i.Add(i.CheckNegativeField(i.price, "price"))
	i.Add(i.CheckNegativeField(i.count, "count"))
	i.Add(i.CheckExceedField(i.count, "count", 12))
	return &i.EntityErrors

}
