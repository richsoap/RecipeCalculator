package recipe

type Recipes []Recipe

type Recipe struct {
	ID      uint64 `json:"id" gorm:"column:id;primaryKey"`
	Item    uint64 `json:"item" gorm:"column:item;index"`
	Depends string `json:"depends" gorm:"column:depends"`
}
