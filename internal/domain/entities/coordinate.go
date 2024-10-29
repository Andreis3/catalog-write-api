package entities

import (
	"github.com/andreis3/catalog-write-api/internal/domain/errors"
)

type Coordinate struct {
	id        int64
	offersID  int64
	latitude  float64
	longitude float64
	errors.EntityErrors
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

func (c *Coordinate) Validate() *errors.EntityErrors {
	if c.latitude != 0 && c.latitude < -90 || c.latitude > 90 {
		c.Add("latitude: must be between -90 and +90 degrees")
	}

	if c.longitude != 0 && c.longitude < -180 || c.longitude > 180 {
		c.Add("longitude: must be between -180 and +180 degrees")
	}

	if c.latitude == 0 && c.longitude != 0 {
		c.Add("latitude: must be set if longitude is set")
	}

	if c.latitude != 0 && c.longitude == 0 {
		c.Add("longitude: must be set if latitude is set")
	}

	return &c.EntityErrors
}
