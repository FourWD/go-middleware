package orm

type Domain struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	GormModel

	Code string `db:"code"  json:"code" gorm:"unique;index;default:null;type:varchar(45)"`
	Name string `db:"name"  json:"name" gorm:"not null;type:varchar(50)"`
	Url  string `db:"url"  json:"url" gorm:"type:varchar(150)"`
}
