package orm

type SuggestType struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	GormModel

	Name string `db:"name"  json:"name" gorm:"index;default:null;type:varchar(200)"`
}
