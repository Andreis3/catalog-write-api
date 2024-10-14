package entities

type Product struct {
	ID          int64
	ExternalID  string
	ApikeyID    string
	Name        string
	Description string
	Status      string
	Brand       string
	ReleaseDate string
	Publish     bool
}

func NewProductBuilder() *Product {
	return &Product{}
}

func (p *Product) SetID(id int64) *Product {
	p.ID = id
	return p
}

func (p *Product) SetExternalID(externalID string) *Product {
	p.ExternalID = externalID
	return p
}

func (p *Product) SetApikeyID(apikeyID string) *Product {
	p.ApikeyID = apikeyID
	return p
}

func (p *Product) SetName(name string) *Product {
	p.Name = name
	return p
}

func (p *Product) SetDescription(description string) *Product {
	p.Description = description
	return p
}

func (p *Product) SetStatus(status string) *Product {
	p.Status = status
	return p
}

func (p *Product) SetBrand(brand string) *Product {
	p.Brand = brand
	return p
}

func (p *Product) SetReleaseDate(releaseDate string) *Product {
	p.ReleaseDate = releaseDate
	return p
}

func (p *Product) SetPublish(publish bool) *Product {
	p.Publish = publish
	return p
}

func (p *Product) Build() *Product {
	return p
}

func (p *Product) Facade(input *Product) *Product {
	return NewProductBuilder().
		SetID(input.ID).
		SetExternalID(input.ExternalID).
		SetApikeyID(input.ApikeyID).
		SetName(input.Name).
		SetDescription(input.Description).
		SetStatus(input.Status).
		SetBrand(input.Brand).
		SetReleaseDate(input.ReleaseDate).
		SetPublish(input.Publish).
		Build()
}
