package graph

import (
	"fmt"
	"testing"

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
--------------
C1: want:a have:nil=>a*1,b*2,c*5,d*12
*/
func TestCaculatorNormalUse(t *testing.T) {
	Convey("run caculator under simple situation", t, func() {
		itemStorage := memory.NewItemStorage()
		recipeStorage := memory.NewRecipeStorage()
		ids := make(map[string]uint64)
		names := []string{"a", "b", "c", "d"}
		for _, name := range names {
			ids[name], _ = itemStorage.AddItem(item.Item{Name: name})
		}
		recipes := map[uint64]map[uint64]int64{
			ids["a"]: {ids["b"]: 2, ids["c"]: 1},
			ids["b"]: {ids["c"]: 2, ids["d"]: 1},
			ids["c"]: {ids["d"]: 2},
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
			So(result.Error, ShouldBeNil)
			for index, node := range result.Nodes {
				So(node.Amount, ShouldEqual, amounts[node.Name])
				SoMsg(fmt.Sprintf("index=%d", index), node.Name, ShouldEqual, order[index])
			}
		})
	})
}
