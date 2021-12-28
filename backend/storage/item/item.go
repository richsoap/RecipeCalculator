package item

type Items []Item

type Item struct {
	ID   uint64 `json:"id" gorm:"column:id;primaryKey"`
	Name string `json:"name" gorm:"column:name;index"`
}
