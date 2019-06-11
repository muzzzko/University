package converter

import (
	"univers/pkg/models"
	"univers/pkg/store/entities"
)

func ConvertRegionFromEntityToModel(regions []*entities.Region) []*models.Region {
	model := make([]*models.Region, len(regions))
	for i, region := range regions {
		model[i] = &models.Region{
			ID: region.ID,
			Name: region.RegionName,
		}
	}

	return model
}
