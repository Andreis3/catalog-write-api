package entities

import (
	"github.com/andreis3/catalog-write-api/internal/domain/commons"
)

type Coordinate struct {
	id        int64
	offersID  int64
	latitude  float64
	longitude float64
	commons.EntityErrors
	commons.ValidateFields
}

func CoordinateBuilder() *Coordinate {
	return &Coordinate{}
}

func (c *Coordinate) GetID() int64 {
	return c.id
}

func (c *Coordinate) GetOffersID() int64 {
	return c.offersID
}

func (c *Coordinate) GetLatitude() float64 {
	return c.latitude
}

func (c *Coordinate) GetLongitude() float64 {
	return c.longitude
}

func (c *Coordinate) SetID(id int64) *Coordinate {
	c.id = id
	return c
}

func (c *Coordinate) SetOffersID(offersID int64) *Coordinate {
	c.offersID = offersID
	return c
}

func (c *Coordinate) SetLatitude(latitude float64) *Coordinate {
	c.latitude = latitude
	return c
}

func (c *Coordinate) SetLongitude(longitude float64) *Coordinate {
	c.longitude = longitude
	return c
}

func (c *Coordinate) Build() *Coordinate {
	return c
}

func (c *Coordinate) Validate() *commons.EntityErrors {
	c.Add(c.CheckLatitudeRange(c.latitude))
	c.Add(c.CheckLongitudeRange(c.longitude))
	c.Add(c.CheckSetField(c.latitude, c.longitude, "latitude", "longitude"))
	c.Add(c.CheckSetField(c.longitude, c.latitude, "longitude", "latitude"))

	return &c.EntityErrors
}
