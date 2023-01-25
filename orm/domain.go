package orm

type Domain struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	GormModel

	Code   string `db:"code" json:"code" gorm:"unique;index;default:null;type:varchar(45)"`
	Name   string `db:"name" json:"name" gorm:"not null;type:varchar(50)"`
	Url    string `db:"url" json:"url" gorm:"type:varchar(150)"`
	Logo   string `db:"logo" json:"logo" gorm:"type:varchar(100)"`
	Color1 string `db:"color1" json:"color1" gorm:"type:varchar(7)"`
	Color2 string `db:"color2" json:"color2" gorm:"type:varchar(7)"`
	Color3 string `db:"color3" json:"color3" gorm:"type:varchar(7)"`
	Color4 string `db:"color4" json:"color4" gorm:"type:varchar(7)"`
	Color5 string `db:"color5" json:"color5" gorm:"type:varchar(7)"`
	Color6 string `db:"color6" json:"color6" gorm:"type:varchar(7)"`
}
