package entities

import (
	"github.com/andreis3/catalog-write-api/internal/domain/commons"
)

const (
	IMAGE = "image"
	VIDEO = "video"
)

var MediaType = [...]string{IMAGE, VIDEO}

type Media struct {
	id          int64
	skuID       int64
	url         string
	mediaType   string
	description string
	index       int
	commons.EntityErrors
	commons.ValidateFields
}

func MediaBuilder() *Media {
	return &Media{}
}

func (m *Media) GetID() int64 {
	return m.id
}

func (m *Media) GetSkuID() int64 {
	return m.skuID
}

func (m *Media) GetURL() string {
	return m.url
}

func (m *Media) GetMediaType() string {
	return m.mediaType
}

func (m *Media) GetDescription() string {
	return m.description
}

func (m *Media) GetIndex() int {
	return m.index
}

func (m *Media) SetID(id int64) *Media {
	m.id = id
	return m
}

func (m *Media) SetSkuID(skuID int64) *Media {
	m.skuID = skuID
	return m
}

func (m *Media) SetURL(url string) *Media {
	m.url = url
	return m
}

func (m *Media) SetMediaType(mediaType string) *Media {
	m.mediaType = mediaType
	return m
}

func (m *Media) SetDescription(description string) *Media {
	m.description = description
	return m
}

func (m *Media) SetIndex(index int) *Media {
	m.index = index
	return m
}

func (m *Media) SetMediaTypeImage() *Media {
	m.mediaType = IMAGE
	return m
}

func (m *Media) SetMediaTypeVideo() *Media {
	m.mediaType = VIDEO
	return m
}

func (m *Media) Build() *Media {
	return m
}

func (m *Media) Validate() *commons.EntityErrors {
	m.Add(m.CheckEmptyField(m.url, "url"))
	m.Add(m.CheckEmptyField(m.mediaType, "media_type"))
	m.Add(m.CheckIsValidStatus(m.mediaType, "media_type", MediaType[:]))
	m.Add(m.CheckEmptyField(m.description, "description"))
	m.Add(m.CheckNegativeField(m.index, "index"))
	m.Add(m.CheckFieldEqualZero(m.index, "index"))
	return &m.EntityErrors
}
