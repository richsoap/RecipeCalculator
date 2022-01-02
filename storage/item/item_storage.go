package item

import "github.com/richsoap/RecipeCalculator/errors"

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
	Prefix      string
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

func (o *OptionStruct) IsValid() error {
	modeSelected := false
	// FilterMode for accuracy query
	if len(o.FilterNames) > 0 || len(o.FilterIDs) > 0 {
		modeSelected = true
	}
	if len(o.Prefix) > 0 {
		if modeSelected {
			return errors.CONFLICT_OPTIONS
		}
		modeSelected = true
	}
	return nil
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

func HavePrefix(prefix string) SearchOption {
	return func(option *OptionStruct) {
		option.Prefix = prefix
	}
}
