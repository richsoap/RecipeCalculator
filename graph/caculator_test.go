package graph

import (
	"errors"
	"testing"

	myError "github.com/richsoap/RecipeCalculator/errors"
	"github.com/richsoap/RecipeCalculator/storage/impl/memory"
	"github.com/richsoap/RecipeCalculator/storage/item"
	"github.com/richsoap/RecipeCalculator/storage/recipe"
	"github.com/richsoap/RecipeCalculator/utils"
	. "github.com/smartystreets/goconvey/convey"
)

/*
a=b*2+c
b=c*2+d
c=d*2
e=1*a
e=20*d
--------------
C1: want:a have:nil=>a*1,b*2,c*5,d*12
*/
func TestCaculatorNormalUse(t *testing.T) {
	Convey("run caculator under simple situation", t, func() {
		itemStorage := memory.NewItemStorage()
		recipeStorage := memory.NewRecipeStorage()
		ids := make(map[string]uint64)
		names := []string{"a", "b", "c", "d", "e"}
		for _, name := range names {
			ids[name], _ = itemStorage.AddItem(item.Item{Name: name})
		}
		recipes := map[uint64]map[uint64]int64{
			ids["a"]: {ids["b"]: 2, ids["c"]: 1},
			ids["b"]: {ids["c"]: 2, ids["d"]: 1},
			ids["c"]: {ids["d"]: 2},
			ids["e"]: {ids["a"]: 1},
		}
		for key, designedRecipe := range recipes {
			recipeStorage.AddRecipe(recipe.Recipe{
				Item:    key,
				Depends: utils.MapDependsToString(designedRecipe),
			})
		}
		caculator := NewGraphCaculator(itemStorage, recipeStorage)
		Convey("query one a, with nothing in pocket", func() {
			result := caculator.Caculate(ItemStack{ids["a"]: 1}, nil, nil)
			amounts := map[string]int{"a": 1, "b": 2, "c": 5, "d": 12}
			order := []string{"d", "c", "b", "a"}
			checkResult(result, amounts, order)
		})
		Convey("add another recipe for e", func() {
			recipeID, err := recipeStorage.AddRecipe(recipe.Recipe{
				Item:    ids["e"],
				Depends: utils.MapDependsToString(map[uint64]int64{ids["d"]: 20}),
			})
			So(err, ShouldBeNil)
			Convey("not provided recipe select", func() {
				result := caculator.Caculate(ItemStack{ids["e"]: 1}, nil, nil)
				So(errors.Is(result.Error, myError.RECIPE_NOT_PROVIDED), ShouldBeTrue)
				So(len(result.UndefinedRecipes[ids["e"]]), ShouldEqual, 2)
			})
			Convey("provided recipe select", func() {
				result := caculator.Caculate(ItemStack{ids["e"]: 1}, nil, RecipeSelect{ids["e"]: recipeID})
				amounts := map[string]int{"e": 1, "d": 20}
				order := []string{"d", "e"}
				checkResult(result, amounts, order)
			})
		})
	})
}

func checkResult(result *CaculateResult, amounts map[string]int, order []string) {
	So(result.Error, ShouldBeNil)
	if amounts != nil {
		checkResultAmount(result.Nodes, amounts)
	}
	if order != nil {
		checkResultOrder(result.Nodes, order)
	}
}

func checkResultAmount(nodes OrderedNodes, amounts map[string]int) {
	for _, node := range nodes {
		So(node.Amount, ShouldEqual, amounts[node.Name])
	}
}

func checkResultOrder(nodes OrderedNodes, orders []string) {
	for index, node := range nodes {
		So(node.Name, ShouldEqual, orders[index])
	}
}
