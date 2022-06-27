package locationService

import (
	"scan/app/service/location/model"
)

func (c *Config) All() ([]model.Location, error) {
	locations, err := c.locationRepository.All()
	if err != nil {
		return nil, err
	}

	return locations, nil
}
