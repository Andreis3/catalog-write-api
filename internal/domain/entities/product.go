package entities

import (
	"github.com/andreis3/catalog-write-api/internal/domain/commons"
)

const (
	ENABLED  = "enabled"
	DISABLED = "disabled"
)

var ProductStatus = [...]string{ENABLED, DISABLED}

type Product struct {
	id          int64
	apikeyID    int64
	externalID  string
	apikey      string
	name        string
	description string
	status      string
	brand       string
	releaseDate string
	commons.EntityErrors
	commons.ValidateFields
}

func ProductBuilder() *Product {
	return &Product{}
}

func (p *Product) GetID() int64 {
	return p.id
}

func (p *Product) GetExternalID() string {
	return p.externalID
}

func (p *Product) GetApikeyID() int64 {
	return p.apikeyID
}

func (p *Product) GetApikey() string {
	return p.apikey
}

func (p *Product) GetName() string {
	return p.name
}

func (p *Product) GetDescription() string {
	return p.description
}

func (p *Product) GetStatus() string {
	return p.status
}

func (p *Product) GetBrand() string {
	return p.brand
}

func (p *Product) GetReleaseDate() string {
	return p.releaseDate
}

func (p *Product) SetID(id int64) *Product {
	p.id = id
	return p
}

func (p *Product) SetExternalID(externalID string) *Product {
	p.externalID = externalID
	return p
}

func (p *Product) SetApikeyID(apikeyID int64) *Product {
	p.apikeyID = apikeyID
	return p
}

func (p *Product) SetApikey(apikey string) *Product {
	p.apikey = apikey
	return p
}

func (p *Product) SetName(name string) *Product {
	p.name = name
	return p
}

func (p *Product) SetDescription(description string) *Product {
	p.description = description
	return p
}

func (p *Product) SetStatus(status string) *Product {
	p.status = status
	return p
}

func (p *Product) SetBrand(brand string) *Product {
	p.brand = brand
	return p
}

func (p *Product) SetReleaseDate(releaseDate string) *Product {
	p.releaseDate = releaseDate
	return p
}

func (p *Product) Build() *Product {
	return p
}

func (p *Product) Validate() *commons.EntityErrors {
	p.Add(p.CheckEmptyField(p.externalID, "external_id"))
	p.Add(p.CheckEmptyField(p.apikey, "apikey"))
	p.Add(p.CheckEmptyField(p.name, "name"))
	p.Add(p.CheckEmptyField(p.description, "description"))
	p.Add(p.CheckEmptyField(p.brand, "brand"))
	p.Add(p.CheckEmptyField(p.status, "status"))
	p.Add(p.CheckIsValidStatus(p.status, "status", ProductStatus[:]))
	p.Add(p.CheckMaxCharacters(p.externalID, "external_id", 50))
	p.Add(p.CheckMaxCharacters(p.name, "name", 50))
	p.Add(p.CheckMaxCharacters(p.description, "description", 255))
	p.Add(p.CheckMaxCharacters(p.brand, "brand", 100))
	return &p.EntityErrors
}
