package storage

import (
	"github.com/richsoap/RecipeCalculator/storage/item"
	"github.com/richsoap/RecipeCalculator/storage/recipe"
)

type Storage interface {
	item.ItemStorage
	recipe.RecipeStorage
}
