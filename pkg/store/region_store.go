package store

import "univers/pkg/store/entities"

func (s *Store) GetRegions() []*entities.Region {
	regions := make([]*entities.Region, 0)
	if err := s.db.Find(&regions).Error; nil != err {
		panic(err)
	}

	return regions
}