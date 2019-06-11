package store

import (
	"univers/pkg/store/entities"
	"univers/pkg/utils"
)

func (s *Store) GetUniversity(id int64) (*entities.University, bool) {
	university := &entities.University{}

	if err := s.db.Preload("Faculties").
		Preload("Faculties.Conditions").
		Preload("Region").
		Preload("Rector").
		Find(university, "id = ?", id).Error; utils.NotFound(err) {
		return nil, false
	}

	return university, true
}

func (s *Store) GetUniversities(offset, limit, regionID *int64, status *string) ([]*entities.University, bool) {
	universities := make([]*entities.University, 0)

	db := s.db.Offset(*offset).Limit(*limit)
	if nil != regionID {
		db = db.Where("region_id = ?", *regionID)
	}
	if nil != status {
		db = db.Where("status = ?", *status)
	}
	if err := db.Find(&universities).Error; utils.NotFound(err) {
		return universities, false
	}

	return universities, true
}

func (s *Store) GetCountUniversity(regionID *int64, status *string) int64 {
	var count int64
	db := s.db.Model(&entities.University{})
	if nil != regionID {
		db = db.Where("region_id = ?", *regionID)
	}
	if nil != status {
		db = db.Where("status = ?", *status)
	}
	if err := db.Count(&count).Error; nil != err {
		panic(err)
	}

	return count
}

func (s *Store) CreateUniversity(university *entities.University) {
	if err := s.db.Create(university).Error; nil != err {
		panic(err)
	}
}

func (s *Store) UpdateUniversity(university *entities.University) {
	tx := s.db.Begin()
	if err := tx.Model(university).Association("Faculties").Replace(university.Faculties).Error; nil != err {
		panic(err)
	}

	if err := tx.Model(university).Updates(university).Error; nil != err {
		tx.Rollback()
		panic(err)
	}

	university.Faculties = nil
	if err := tx.Preload("Faculties").
		Preload("Faculties.Conditions").
		Preload("Region").
		Preload("Rector").
		Find(university, "id = ?", university.ID).Error; nil != err {
		tx.Rollback()
		panic(err)
	}

	tx.Commit()
}