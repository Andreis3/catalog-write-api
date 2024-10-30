package entities

import (
	"slices"

	"github.com/andreis3/catalog-write-api/internal/domain/errors"
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
	errors.EntityErrors
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

func (m *Media) Validate() *errors.EntityErrors {

	if m.url == "" {
		m.Add("url: cannot be empty")
	}

	if m.mediaType == "" {
		m.Add("media_type: cannot be empty")
	} else if !slices.Contains(MediaType[:], m.mediaType) {
		m.Add("media_type: media type is invalid, valid values are image or video")
	}

	if m.description == "" {
		m.Add("description: cannot be empty")
	}

	if m.index <= 0 {
		m.Add("index: cannot be less than 0")
	}

	return &m.EntityErrors
}
