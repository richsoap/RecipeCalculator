package memory

import (
	"github.com/richsoap/RecipeCalculator/errors"
	"github.com/richsoap/RecipeCalculator/storage/item"
	"github.com/sirupsen/logrus"
)

// only for test
type memoryItemStorage struct {
	items []memoryItemDescription
}

type memoryItemDescription struct {
	Item   *item.Item
	Remove bool
}

func NewItemStorage() item.ItemStorage {
	logrus.WithField("storage", "memory").Warn("memory item storage only for test")
	return &memoryItemStorage{
		items: make([]memoryItemDescription, 0),
	}
}

func (s *memoryItemStorage) AddItem(item item.Item) (uint64, error) {
	item.ID = uint64(len(s.items))
	s.items = append(s.items, memoryItemDescription{
		Item:   &item,
		Remove: false,
	})
	return item.ID, nil
}

func (s *memoryItemStorage) DeleteItem(ID uint64) error {
	id := int(ID)
	if len(s.items) > id {
		s.items[id].Remove = true
	}
	return nil
}

func (s *memoryItemStorage) UpdateItem(item item.Item) error {
	id := int(item.ID)
	if len(s.items) > id && !s.items[id].Remove {
		s.items[id].Item = &item
		return nil
	}
	return errors.ITEM_NOT_FOUND
}

func (s *memoryItemStorage) GetItems(options ...item.SearchOption) (item.Items, error) {
	option := item.ParseOptions(options...)
	ids := make(map[uint64]int)
	names := make(map[string]int)
	result := make(item.Items, 0)
	for _, id := range option.FilterIDs {
		ids[id] = 0
	}
	for _, name := range option.FilterNames {
		names[name] = 0
	}
	for _, desc := range s.items {
		if desc.Remove {
			continue
		}
		if _, exist := ids[desc.Item.ID]; len(ids) > 0 && !exist {
			continue
		}
		if _, exist := names[desc.Item.Name]; len(names) > 0 && !exist {
			continue
		}
		result = append(result, *desc.Item)
	}
	return result, nil
}
