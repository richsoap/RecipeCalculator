package gorm

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/richsoap/RecipeCalculator/storage/item"
	"github.com/richsoap/RecipeCalculator/storage/recipe"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSqliteItemStorage(t *testing.T) {
	Convey("CUDR item with SQLite", t, func() {
		file, err := ioutil.TempFile(".", "sqlite")
		So(err, ShouldBeNil)
		file.Close()
		defer os.Remove(file.Name())
		db, err := NewSqliteItemStorage(file.Name())
		So(err, ShouldBeNil)
		// Create
		origins := item.Items{
			{Name: "a"},
			{Name: "b"},
			{Name: "c"},
		}
		for index := range origins {
			origins[index].ID, err = db.AddItem(origins[index])
			So(err, ShouldBeNil)
		}
		// Read all
		items, err := db.GetItems()
		So(err, ShouldBeNil)
		So(items, ShouldResemble, origins)
		// Read by id
		items, err = db.GetItems(item.FilterByIDs(origins[1].ID))
		So(err, ShouldBeNil)
		So(items[0], ShouldResemble, origins[1])
		// Update
		newItem := origins[1]
		newItem.Name = "a"
		So(db.UpdateItem(newItem), ShouldBeNil)
		// check if 2 "a" in db
		items, err = db.GetItems(item.FilterByNames("a"))
		So(err, ShouldBeNil)
		So(len(items), ShouldEqual, 2)
		// Delete All
		for _, item := range origins {
			So(db.DeleteItem(item.ID), ShouldBeNil)
		}
		// Check if all deleted
		items, err = db.GetItems()
		So(err, ShouldBeNil)
		So(len(items), ShouldEqual, 0)
	})
}

func TestSqliteRecipeStorage(t *testing.T) {
	Convey("CUDR recipe with SQLite", t, func() {
		file, err := ioutil.TempFile(".", "sqlite")
		So(err, ShouldBeNil)
		file.Close()
		defer os.Remove(file.Name())
		db, err := NewSqliteRecipeStorage(file.Name())
		So(err, ShouldBeNil)
		// Create
		// 5 = 2*2+1
		// 3 = 1*3
		// 2 = 1*2
		origins := recipe.Recipes{
			{Item: 5, Depends: "2:2;1:1"},
			{Item: 3, Depends: "1:3"},
			{Item: 2, Depends: "1:2"},
		}
		for index := range origins {
			origins[index].ID, err = db.AddRecipe(origins[index])
			So(err, ShouldBeNil)
		}
		// Read all
		recipes, err := db.GetRecipes()
		So(err, ShouldBeNil)
		So(recipes, ShouldResemble, origins)
		// Read by id
		recipes, err = db.GetRecipes(recipe.FilterByIDs(origins[1].ID))
		So(err, ShouldBeNil)
		So(recipes[0], ShouldResemble, origins[1])
		// Update
		newRecipe := origins[1]
		newRecipe.Item = 5
		newRecipe.Depends = "1:5"
		So(db.UpdateRecipe(newRecipe), ShouldBeNil)
		// check if 2 "5" in db
		recipes, err = db.GetRecipes(recipe.FilterByItems(newRecipe.Item))
		So(err, ShouldBeNil)
		So(len(recipes), ShouldEqual, 2)
		// Delete All
		for _, item := range origins {
			So(db.DeleteRecipe(item.ID), ShouldBeNil)
		}
		// Check if all deleted
		recipes, err = db.GetRecipes()
		So(err, ShouldBeNil)
		So(len(recipes), ShouldEqual, 0)
	})
}
