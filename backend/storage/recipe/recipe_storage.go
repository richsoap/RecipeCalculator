package recipe

type RecipeStorage interface {
	GetRecipes(uint64) (Recipes, error)
	AddRecipe(Recipe) error
	UpdateRecipe(Recipe) error
	DeleteRecipe(uint64) error
}

type RecipeSearchOption func(*recipeSearchOption)

type recipeSearchOption struct {
}
