package graph

import (
	"fmt"

	"github.com/richsoap/RecipeCalculator/errors"
	"github.com/richsoap/RecipeCalculator/storage/item"
	"github.com/richsoap/RecipeCalculator/storage/recipe"
)

type GraphCaculator struct {
	itemStorage   item.ItemStorage
	recipeStorage recipe.RecipeStorage
}

func NewGraphCaculator(itemStorage item.ItemStorage, recipeStorage recipe.RecipeStorage) *GraphCaculator {
	return &GraphCaculator{itemStorage: itemStorage, recipeStorage: recipeStorage}
}

type ItemStack map[uint64]int

type RecipeSelect map[uint64]uint64

type CaculateResult struct {
	Nodes            OrderedNodes
	UndefinedRecipes map[uint64]recipe.Recipes
	Error            error
}

// S1: 分析出配方依赖图Alpha，以Item为Node，配方所需物品数量为Edge
// S2: 预设所需物品数量
// S3: 每次从alpha中取入度为0的Node，基于该Node所需数量-已有数量，增加原料节点数量。如果已有数量超过所需数量，则不增加原料节点数量
// S4: 将S3中所选Node放入输出栈，将节点从图Alpha中删除，并更新入度
// S5: 重复S3和S4，直到A. Alpha中有节点，但入度均不为0，说明有环，无解; B. Alpha为空，到下一步
// S6: 从输出栈中挨个弹出Node，即为资源收集顺序
func (r *GraphCaculator) Caculate(want ItemStack, have ItemStack, recipeSelect RecipeSelect) *CaculateResult {
	// S1
	graph, undefinedRecipes, err := r.buildRecipeGraph(want, recipeSelect)
	if err != nil {
		return &CaculateResult{Error: err}
	}
	// 如果有未决定的配方，则返回error
	if len(undefinedRecipes) > 0 {
		return &CaculateResult{Nodes: graph.ToOrderedNodes(), UndefinedRecipes: undefinedRecipes, Error: errors.RECIPE_NOT_PROVIDED}
	}
	// S2
	for id, amount := range want {
		graph.AdjustAmount(id, amount)
	}
	for id, haveAmount := range have {
		graph.AdjustAmount(id, -haveAmount)
	}
	// S3 + S4
	nodes := r.popNodes(graph)
	if !graph.IsEmpty() {
		return &CaculateResult{
			Nodes: graph.ToOrderedNodes(),
			Error: errors.CIRCLE_DEPENDENCY,
		}
	}
	for i := range nodes {
		if i >= len(nodes)-1-i {
			break
		}
		nodes[i], nodes[len(nodes)-1-i] = nodes[len(nodes)-1-i], nodes[i]
	}
	return &CaculateResult{
		Nodes: nodes,
		Error: nil,
	}
}

func (r *GraphCaculator) buildRecipeGraph(want ItemStack, recipeSelect RecipeSelect) (*Graph, map[uint64]recipe.Recipes, error) {
	idStacks := make([]uint64, 0)
	alphaGraph := NewGraph()
	// s1
	for id := range want {
		idStacks = append(idStacks, id)
	}
	undefinedRecipes := make(map[uint64]recipe.Recipes)
	for len(idStacks) > 0 {
		id := idStacks[len(idStacks)-1]
		idStacks = idStacks[:len(idStacks)-1]
		var currentItem *item.Item
		var currentRecipe *recipe.Recipe
		// s1.1 如果节点已在图中，则无需再次计算
		if alphaGraph.IsExist(id) {
			continue
		}
		// s1.2 不存在则试图获取配方
		items, err := r.itemStorage.GetItems(item.FilterByIDs(id))
		if err != nil {
			return nil, nil, err
		}
		if len(items) == 0 {
			return nil, nil, fmt.Errorf("%w: %d", errors.ITEM_NOT_FOUND, id)
		}
		currentItem = &items[0]
		// s1.2.1 已经指定了配方
		if recipeID, exist := recipeSelect[id]; exist {
			recipe, err := r.recipeStorage.GetRecipes(recipe.FilterByIDs(recipeID))
			if err != nil {
				return nil, nil, err
			}
			if len(recipe) == 0 {
				return nil, nil, fmt.Errorf("%w: %d", errors.RECIPE_NOT_FOUND, recipeID)
			}
			currentRecipe = &recipe[0]
		} else {
			// s1.2.2 没有指定配方
			recipes, err := r.recipeStorage.GetRecipes(recipe.FilterByItems(id))
			if err != nil {
				return nil, nil, err
			}
			// s.1.2.2.1 唯一配方
			if len(recipes) == 1 {
				currentRecipe = &recipes[0]
			} else if len(recipes) > 1 {
				//s1.2.2.2 有多配方，需要指定
				undefinedRecipes[id] = recipes
			}
		}
		node, err := NewNodeFromItemAndRecipe(currentItem, currentRecipe)
		if err != nil {
			return nil, nil, err
		}
		alphaGraph.AddNode(node)
	}
	return alphaGraph, undefinedRecipes, nil
}

func (r *GraphCaculator) popNodes(g *Graph) OrderedNodes {
	result := make(OrderedNodes, 0)
	for !g.IsEmpty() && g.GetZeroDegreeNode() != nil {
		node := g.GetZeroDegreeNode()
		if node.Amount <= 0 {
			node.Amount = 0
		} else {
			for id, amount := range node.Depends {
				g.AdjustAmount(id, amount*node.Amount)
			}
		}
		result = append(result, node)
		g.RemoveNode(node)
	}
	return result
}
