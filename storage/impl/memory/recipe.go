package memory

import (
	"github.com/richsoap/RecipeCalculator/errors"
	"github.com/richsoap/RecipeCalculator/storage/recipe"
	"github.com/sirupsen/logrus"
)

type recipeStorage struct {
	recipes []recipeDescription
}

type recipeDescription struct {
	Recipe *recipe.Recipe
	Remove bool
}

func NewRecipeStorage() recipe.RecipeStorage {
	return &recipeStorage{
		make([]recipeDescription, 0),
	}
}

func (s *recipeStorage) AddRecipe(recipe recipe.Recipe) (uint64, error) {
	recipe.ID = uint64(len(s.recipes))
	s.recipes = append(s.recipes, recipeDescription{
		Recipe: &recipe,
		Remove: false,
	})
	logrus.WithField("item", recipe.Item).WithField("depends", recipe.Depends).Info("new recipe")
	return recipe.ID, nil
}

func (s *recipeStorage) UpdateRecipe(recipe recipe.Recipe) error {
	id := int(recipe.ID)
	if len(s.recipes) > id && !s.recipes[id].Remove {
		s.recipes[id].Recipe = &recipe
		return nil
	}
	return errors.RECIPE_NOT_FOUND
}

func (s *recipeStorage) DeleteRecipe(ID uint64) error {
	id := int(ID)
	if len(s.recipes) > id {
		s.recipes[id].Remove = true
	}
	return nil
}

func (s *recipeStorage) GetRecipes(options ...recipe.SearchOption) (recipe.Recipes, error) {
	option := recipe.ParseOptions(options...)
	ids := make(map[uint64]int)
	items := make(map[uint64]int)
	result := make(recipe.Recipes, 0)
	for _, id := range option.FilterIDs {
		ids[id] = 0
	}
	for _, id := range option.FilterItems {
		items[id] = 0
	}
	for _, desc := range s.recipes {
		if desc.Remove {
			continue
		}
		if _, exist := ids[desc.Recipe.ID]; len(ids) > 0 && !exist {
			continue
		}
		if _, exist := items[desc.Recipe.Item]; len(items) > 0 && !exist {
			continue
		}
		result = append(result, *desc.Recipe)
	}
	return result, nil

}
