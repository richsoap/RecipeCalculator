package gorm

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/richsoap/RecipeCalculator/storage/item"
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
