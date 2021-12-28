package gorm

import (
	"github.com/richsoap/RecipeCalculator/storage/item"
	"github.com/richsoap/RecipeCalculator/storage/recipe"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSqliteItemStorage(dbPath string) (item.ItemStorage, error) {
	if db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{}); err != nil {
		logrus.WithFields(logrus.Fields{
			"path":      db,
			"error_msg": err,
		}).Errorf("open db error")
		return nil, err
	} else {
		return newGormItemStorage(db), nil
	}
}

func NewSqliteRecipeStorage(dbPath string) (recipe.RecipeStorage, error) {
	if db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{}); err != nil {
		logrus.WithFields(logrus.Fields{
			"path":      db,
			"error_msg": err,
		}).Errorf("open db error")
		return nil, err
	} else {
		return newGormRecipeStorage(db), nil
	}
}
