package recipe

type Recipes []Recipe

type Recipe struct {
	Id      uint64            `json:"id" gorm:"column:id"`
	Item    uint64            `json:"item" gorm:"column:item"`
	Depends map[uint64]uint64 `json:"depends" gorm:"column:depends"`
}
