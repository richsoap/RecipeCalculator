package item

type ItemStorage interface {
	GetItems(...SearchOption) (Items, error)
	AddItem(Item) (uint64, error)
	UpdateItem(Item) error
	DeleteItem(uint64) error
}

type SearchOption func(*OptionStruct)

type OptionStruct struct {
	FilterNames []string
	FilterIDs   []uint64
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

func FilterByNames(names ...string) SearchOption {
	return func(option *OptionStruct) {
		option.FilterNames = names
	}
}
