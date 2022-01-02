package memory

import (
	"testing"

	"github.com/richsoap/RecipeCalculator/storage/impl/test_sample"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMemoryRecipeTest(t *testing.T) {
	Convey("run test for memory recipe storage", t, func() {
		s := NewRecipeStorage()
		test_sample.RunRecipeStorage(t, s)
	})

}
