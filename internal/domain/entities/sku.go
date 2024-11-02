package entities

import "github.com/andreis3/catalog-write-api/internal/domain/errors"

const (
	ACTIVE_SKU   = "active"
	INACTIVE_SKU = "inactive"
)

var SkuStatus = [...]string{ACTIVE_SKU, INACTIVE_SKU}

type Sku struct {
	id          int64
	externalID  string
	productID   int64
	name        string
	description string
	gtin        string
	status      string
	errors.EntityErrors
	errors.ValidateFields
}

func SkuBuilder() *Sku {
	return &Sku{}
}

func (s *Sku) GetID() int64 {
	return s.id
}

func (s *Sku) GetExternalID() string {
	return s.externalID
}

func (s *Sku) GetProductID() int64 {
	return s.productID
}

func (s *Sku) GetName() string {
	return s.name
}

func (s *Sku) GetDescription() string {
	return s.description
}

func (s *Sku) GetGtin() string {
	return s.gtin
}

func (s *Sku) GetStatus() string {
	return s.status
}

func (s *Sku) SetID(id int64) *Sku {
	s.id = id
	return s
}

func (s *Sku) SetExternalID(externalID string) *Sku {
	s.externalID = externalID
	return s
}

func (s *Sku) SetProductID(productID int64) *Sku {
	s.productID = productID
	return s
}

func (s *Sku) SetName(name string) *Sku {
	s.name = name
	return s
}

func (s *Sku) SetDescription(description string) *Sku {
	s.description = description
	return s
}

func (s *Sku) SetGtin(gtin string) *Sku {
	s.gtin = gtin
	return s
}

func (s *Sku) SetStatus(status string) *Sku {
	s.status = status
	return s
}

func (s *Sku) SetStatusActive() *Sku {
	s.status = ACTIVE_SKU
	return s
}

func (s *Sku) SetStatusInactive() *Sku {
	s.status = INACTIVE_SKU
	return s
}

func (s *Sku) Build() *Sku {
	return s
}

func (s *Sku) Validate() *errors.EntityErrors {
	s.Add(s.CheckEmptyField(s.externalID, "external_id"))
	s.Add(s.CheckEmptyField(s.name, "name"))
	s.Add(s.CheckEmptyField(s.description, "description"))
	s.Add(s.CheckEmptyField(s.gtin, "gtin"))
	s.Add(s.CheckEmptyField(s.status, "status"))
	s.Add(s.CheckMaxCharacters(s.externalID, "external_id", 20))
	s.Add(s.CheckMaxCharacters(s.name, "name", 100))
	s.Add(s.CheckMaxCharacters(s.description, "description", 255))
	s.Add(s.CheckIsValidStatus(s.status, "status", SkuStatus[:]))
	return &s.EntityErrors
}
