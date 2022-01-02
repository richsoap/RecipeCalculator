package recipe

type RecipeStorage interface {
	GetRecipes(...SearchOption) (Recipes, error)
	AddRecipe(Recipe) (uint64, error)
	UpdateRecipe(Recipe) error
	DeleteRecipe(uint64) error
}

type SearchOption func(*OptionStruct)

type OptionStruct struct {
	FilterIDs   []uint64
	FilterItems []uint64
}

func ParseOptions(options ...SearchOption) *OptionStruct {
	result := &OptionStruct{}
	return result.Parse(options...)
}

func (o *OptionStruct) Parse(options ...SearchOption) *OptionStruct {
	for _, option := range options {
		option(o)
	}
	return o
}

func FilterByIDs(ids ...uint64) SearchOption {
	return func(option *OptionStruct) {
		option.FilterIDs = ids
	}
}

func FilterByItems(items ...uint64) SearchOption {
	return func(option *OptionStruct) {
		option.FilterItems = items
	}
}
