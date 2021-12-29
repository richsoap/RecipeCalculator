package memory

import (
	"testing"

	"github.com/richsoap/RecipeCalculator/storage/impl/test_sample"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMemoryItemStorage(t *testing.T) {
	Convey("run test for memory item storage", t, func() {
		s := NewItemStorage()
		test_sample.RunItemStorageTest(t, s)
	})
}
