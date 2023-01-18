package orm

type Payment struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	GormModel

	Name     string `db:"name"  json:"name" gorm:"dafault:null;type:varchar(45)"`
	NameEn   string `db:"name_en"  json:"name_en" gorm:"dafault:null;type:varchar(45)"`
	RowOrder int    ` db:"row_order" json:"row_order" gorm:"type:int(11);" `
}
