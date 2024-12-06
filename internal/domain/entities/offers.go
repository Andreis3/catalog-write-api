package entities

import (
	"github.com/andreis3/catalog-write-api/internal/domain/commons"
)

const (
	AVAILABLE   = "available"
	UNAVAILABLE = "unavailable"
	REMOVED     = "removed"
)

var OfferStatus = [...]string{AVAILABLE, UNAVAILABLE, REMOVED}

type Offer struct {
	id           int64
	externalID   string
	skuID        int64
	name         string
	description  string
	price        float64
	oldPrice     float64
	stock        int64
	status       string
	salesChannel string
	seller       string
	commons.EntityErrors
	commons.ValidateFields
}

func OfferBuilder() *Offer {
	return &Offer{}
}

func (o *Offer) GetID() int64 {
	return o.id
}

func (o *Offer) GetExternalID() string {
	return o.externalID
}

func (o *Offer) GetSkuID() int64 {
	return o.skuID
}

func (o *Offer) GetName() string {
	return o.name
}

func (o *Offer) GetDescription() string {
	return o.description
}

func (o *Offer) GetPrice() float64 {
	return o.price
}

func (o *Offer) GetOldPrice() float64 {
	return o.oldPrice
}

func (o *Offer) GetStock() int64 {
	return o.stock
}

func (o *Offer) GetStatus() string {
	return o.status
}

func (o *Offer) GetSalesChannel() string {
	return o.salesChannel
}

func (o *Offer) GetSeller() string {
	return o.seller
}

func (o *Offer) SetID(id int64) *Offer {
	o.id = id
	return o
}

func (o *Offer) SetExternalID(externalID string) *Offer {
	o.externalID = externalID
	return o
}

func (o *Offer) SetSkuID(skuID int64) *Offer {
	o.skuID = skuID
	return o
}

func (o *Offer) SetName(name string) *Offer {
	o.name = name
	return o
}

func (o *Offer) SetDescription(description string) *Offer {
	o.description = description
	return o
}

func (o *Offer) SetPrice(price float64) *Offer {
	o.price = price
	return o
}

func (o *Offer) SetOldPrice(oldPrice float64) *Offer {
	o.oldPrice = oldPrice
	return o
}

func (o *Offer) SetStock(stock int64) *Offer {
	o.stock = stock
	return o
}

func (o *Offer) SetStatus(status string) *Offer {
	o.status = status
	return o
}

func (o *Offer) SetStatusAvailable() *Offer {
	o.status = AVAILABLE
	return o
}

func (o *Offer) SetStatusUnavailable() *Offer {
	o.status = UNAVAILABLE
	return o
}

func (o *Offer) SetStatusRemoved() *Offer {
	o.status = REMOVED
	return o
}

func (o *Offer) SetSalesChannel(salesChannel string) *Offer {
	o.salesChannel = salesChannel
	return o
}

func (o *Offer) SetSeller(seller string) *Offer {
	o.seller = seller
	return o
}

func (o *Offer) Build() *Offer {
	return o
}

func (o *Offer) Validate() *commons.EntityErrors {
	o.Add(o.CheckEmptyField(o.externalID, "external_id"))
	o.Add(o.CheckEmptyField(o.name, "name"))
	o.Add(o.CheckEmptyField(o.description, "description"))
	o.Add(o.CheckMaxCharacters(o.name, "name", 100))
	o.Add(o.CheckMaxCharacters(o.description, "description", 255))
	o.Add(o.CheckNegativeField(o.price, "price"))
	o.Add(o.CheckNegativeField(o.oldPrice, "old_price"))
	o.Add(o.CheckNegativeField(o.stock, "stock"))
	o.Add(o.CheckEmptyField(o.status, "status"))
	o.Add(o.CheckIsValidStatus(o.status, "status", OfferStatus[:]))

	return &o.EntityErrors
}
