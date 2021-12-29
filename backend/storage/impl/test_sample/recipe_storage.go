package test_sample

import (
	"testing"

	"github.com/richsoap/RecipeCalculator/storage/recipe"
	. "github.com/smartystreets/goconvey/convey"
)

func RunRecipeStorage(t *testing.T, s recipe.RecipeStorage) {
	// Create
	// 5 = 2*2+1
	// 3 = 1*3
	// 2 = 1*2
	var err error
	origins := recipe.Recipes{
		{Item: 5, Depends: "2:2;1:1"},
		{Item: 3, Depends: "1:3"},
		{Item: 2, Depends: "1:2"},
	}
	for index := range origins {
		origins[index].ID, err = s.AddRecipe(origins[index])
		So(err, ShouldBeNil)
	}
	// Read all
	recipes, err := s.GetRecipes()
	So(err, ShouldBeNil)
	So(recipes, ShouldResemble, origins)
	// Read by id
	recipes, err = s.GetRecipes(recipe.FilterByIDs(origins[1].ID))
	So(err, ShouldBeNil)
	So(recipes[0], ShouldResemble, origins[1])
	// Update
	newRecipe := origins[1]
	newRecipe.Item = 5
	newRecipe.Depends = "1:5"
	So(s.UpdateRecipe(newRecipe), ShouldBeNil)
	// check if 2 "5" in db
	recipes, err = s.GetRecipes(recipe.FilterByItems(newRecipe.Item))
	So(err, ShouldBeNil)
	So(len(recipes), ShouldEqual, 2)
	// Delete All
	for _, item := range origins {
		So(s.DeleteRecipe(item.ID), ShouldBeNil)
	}
	// Check if all deleted
	recipes, err = s.GetRecipes()
	So(err, ShouldBeNil)
	So(len(recipes), ShouldEqual, 0)
}
