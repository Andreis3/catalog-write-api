package entities

type Sku struct {
	ID          int64
	ExternalID  string
	ProductID   int64
	Name        string
	Description string
	GTIN        string
	Status      string
}
