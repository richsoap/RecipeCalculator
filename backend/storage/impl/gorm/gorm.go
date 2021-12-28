package gorm

import (
	"github.com/richsoap/RecipeCalculator/errors"
	"github.com/richsoap/RecipeCalculator/storage/item"
	"gorm.io/gorm"
)

type gormItemStorage struct {
	db *gorm.DB
}

func newGormItemStorage(db *gorm.DB) item.ItemStorage {
	return &gormItemStorage{db: db}
}

// TODO
func (s *gormItemStorage) GetItems(options ...item.SearchOption) (item.Items, error) {
	if s.db == nil {
		return nil, errors.NOT_INITIALIZED
	}
	option := item.ParseOptions(options...)
	conn := s.db.Model(&item.Item{})
	if len(option.FilterIDs) > 0 {
		conn = conn.Where("id IN ?", option.FilterIDs)
	}
	if len(option.FilterNames) > 0 {
		conn = conn.Where("name IN ?", option.FilterNames)
	}
	result := make(item.Items, 0)
	if sqlResut := conn.Scan(&result); sqlResut.Error != nil {
		return nil, sqlResut.Error
	}
	return result, nil
}

func (s *gormItemStorage) AddItem(item item.Item) (uint64, error) {
	if s.db == nil {
		return 0, errors.NOT_INITIALIZED
	}
	if result := s.db.Create(&item); result.Error != nil {
		return 0, result.Error
	} else {
		return item.ID, nil
	}
}

// TODO
func (s *gormItemStorage) UpdateItem(item item.Item) error {
	if s.db == nil {
		return errors.NOT_INITIALIZED
	}
	s.db.Model(&item).Update("Name", item.Name)
	return nil
}

func (s *gormItemStorage) DeleteItem(id uint64) error {
	if s.db == nil {
		return errors.NOT_INITIALIZED
	}
	s.db.Delete(&item.Item{ID: id})
	return nil
}
