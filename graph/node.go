package graph

import (
	"github.com/richsoap/RecipeCalculator/storage/item"
	"github.com/richsoap/RecipeCalculator/storage/recipe"
	"github.com/richsoap/RecipeCalculator/utils"
)

type Node struct {
	ID      uint64         `json:"id"`
	Name    string         `json:"name"`
	Amount  int            `json:"amount"`
	Depends map[uint64]int `json:"depends"`
}

type OrderedNodes []*Node

func NewNodeFromItemAndRecipe(item *item.Item, recipe *recipe.Recipe) (*Node, error) {
	node := &Node{
		ID:      item.ID,
		Name:    item.Name,
		Amount:  0,
		Depends: make(map[uint64]int),
	}
	if recipe != nil {
		depends, err := utils.StringDependsToMap(recipe.Depends)
		if err != nil {
			return nil, err
		}
		node.Depends = depends
	}
	return node, nil
}
