package tables

type Name_inof struct {
	Id   int    `gorm:"id"`
	Name string `gorm:"name"`
}

func (i Name_inof) TableName() string {
	return "name_info"
}
