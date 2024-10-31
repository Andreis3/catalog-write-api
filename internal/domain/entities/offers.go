package entities

import (
	"slices"
	"strings"

	"github.com/andreis3/catalog-write-api/internal/domain/errors"
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
	errors.EntityErrors
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

func (o *Offer) Validate() *Offer {
	if o.externalID == "" {
		o.Add("external_id: ExternalID is required")
	}
	if o.name == "" {
		o.Add("name: Name is required")
	}
	if strings.Count(o.name, "") > 100 {
		o.Add("name: Name must have a maximum of 100 characters")
	}
	if o.description == "" {
		o.Add("description: Description is required")
	}
	if strings.Count(o.description, "") > 255 {
		o.Add("description: Description must have a maximum of 255 characters")
	}
	if o.price < 0 {
		o.Add("price: Price is required")
	}
	if o.oldPrice < 0 {
		o.Add("old_price: Old price is required")
	}
	if o.stock < 0 {
		o.Add("stock: Stock must be greater than or equal to 0")
	}
	if o.status == "" {
		o.Add("status: Status is required")
	} else if !slices.Contains(OfferStatus[:], o.status) {
		o.Add("status: Invalid status")
	}

	return o
}
