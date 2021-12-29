package gorm

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/richsoap/RecipeCalculator/storage/impl/test_sample"
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
		test_sample.RunItemStorageTest(t, db)
	})
}

func TestSqliteRecipeStorage(t *testing.T) {
	Convey("CUDR recipe with SQLite", t, func() {
		file, err := ioutil.TempFile(".", "sqlite")
		So(err, ShouldBeNil)
		file.Close()
		defer os.Remove(file.Name())
		db, err := NewSqliteRecipeStorage(file.Name())
		So(err, ShouldBeNil)
		test_sample.RunRecipeStorage(t, db)
	})
}
