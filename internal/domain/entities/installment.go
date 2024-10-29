package entities

import "github.com/andreis3/catalog-write-api/internal/domain/errors"

type Installment struct {
	id      int64
	orderID int64
	count   int64
	price   float64
	errors.EntityErrors
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

func (i *Installment) GetCount() int64 {
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

func (i *Installment) SetCount(count int64) *Installment {
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

func (i *Installment) Validate() *errors.EntityErrors {
	if i.price <= -1 {
		i.Add("price: price cannot be a negative number")
	}

	if i.count <= -1 {
		i.Add("count: count cannot be a negative number")
	}

	if i.count > 12 {
		i.Add("count: count cannot exceed 12")
	}

	return &i.EntityErrors

}
