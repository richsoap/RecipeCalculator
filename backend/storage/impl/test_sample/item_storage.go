package test_sample

import (
	"testing"

	"github.com/richsoap/RecipeCalculator/storage/item"
	. "github.com/smartystreets/goconvey/convey"
)

func RunItemStorageTest(t *testing.T, s item.ItemStorage) {
	var err error
	// Create
	origins := item.Items{
		{Name: "a"},
		{Name: "b"},
		{Name: "c"},
	}
	for index := range origins {
		origins[index].ID, err = s.AddItem(origins[index])
		So(err, ShouldBeNil)
	}
	// Read all
	items, err := s.GetItems()
	So(err, ShouldBeNil)
	So(items, ShouldResemble, origins)
	// Read by id
	items, err = s.GetItems(item.FilterByIDs(origins[1].ID))
	So(err, ShouldBeNil)
	So(items[0], ShouldResemble, origins[1])
	// Update
	newItem := origins[1]
	newItem.Name = "a"
	So(s.UpdateItem(newItem), ShouldBeNil)
	// check if 2 "a" in db
	items, err = s.GetItems(item.FilterByNames("a"))
	So(err, ShouldBeNil)
	So(len(items), ShouldEqual, 2)
	// Delete All
	for _, item := range origins {
		So(s.DeleteItem(item.ID), ShouldBeNil)
	}
	// Check if all deleted
	items, err = s.GetItems()
	So(err, ShouldBeNil)
	So(len(items), ShouldEqual, 0)
}
