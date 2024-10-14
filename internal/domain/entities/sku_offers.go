package entities

type SkuOffer struct {
	ID           int64
	ExternalID   string
	SkuID        int64
	Name         string
	Description  string
	Price        float64
	OldPrice     float64
	Stock        int64
	Status       string
	SalesChannel string
	Seller       string
}
