package gorm

import (
	"github.com/richsoap/RecipeCalculator/errors"
	"github.com/richsoap/RecipeCalculator/storage/item"
	"github.com/richsoap/RecipeCalculator/storage/recipe"
	"gorm.io/gorm"
)

type gormStorage struct {
	db *gorm.DB
}

func newGormItemStorage(db *gorm.DB) item.ItemStorage {
	db.AutoMigrate(&item.Item{})
	return &gormStorage{db: db}
}

func newGormRecipeStorage(db *gorm.DB) recipe.RecipeStorage {
	db.AutoMigrate(&recipe.Recipe{})
	return &gormStorage{db: db}
}

func (s *gormStorage) GetItems(options ...item.SearchOption) (item.Items, error) {
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

func (s *gormStorage) AddItem(item item.Item) (uint64, error) {
	if s.db == nil {
		return 0, errors.NOT_INITIALIZED
	}
	if result := s.db.Create(&item); result.Error != nil {
		return 0, result.Error
	} else {
		return item.ID, nil
	}
}

func (s *gormStorage) UpdateItem(item item.Item) error {
	if s.db == nil {
		return errors.NOT_INITIALIZED
	}
	if result := s.db.Updates(&item); result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *gormStorage) DeleteItem(id uint64) error {
	if s.db == nil {
		return errors.NOT_INITIALIZED
	}
	if result := s.db.Delete(&item.Item{}, id); result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *gormStorage) GetRecipes(options ...recipe.SearchOption) (recipe.Recipes, error) {
	if s.db == nil {
		return nil, errors.NOT_INITIALIZED
	}
	option := recipe.ParseOptions(options...)
	conn := s.db.Model(&recipe.Recipe{})
	if len(option.FilterIDs) > 0 {
		conn = conn.Where("id IN ?", option.FilterIDs)
	}
	if len(option.FilterItems) > 0 {
		conn = conn.Where("item IN ?", option.FilterItems)
	}
	result := make(recipe.Recipes, 0)
	if sqlResut := conn.Scan(&result); sqlResut.Error != nil {
		return nil, sqlResut.Error
	}
	return result, nil
}

func (s *gormStorage) AddRecipe(recipe recipe.Recipe) (uint64, error) {
	if s.db == nil {
		return 0, errors.NOT_INITIALIZED
	}
	if result := s.db.Create(&recipe); result.Error != nil {
		return 0, result.Error
	} else {
		return recipe.ID, nil
	}
}

func (s *gormStorage) UpdateRecipe(recipe recipe.Recipe) error {
	if s.db == nil {
		return errors.NOT_INITIALIZED
	}
	if result := s.db.Updates(&recipe); result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *gormStorage) DeleteRecipe(id uint64) error {
	if s.db == nil {
		return errors.NOT_INITIALIZED
	}
	if result := s.db.Delete(&recipe.Recipe{ID: id}); result.Error != nil {
		return result.Error
	}
	return nil
}
