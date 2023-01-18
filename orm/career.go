package orm

type Career struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	GormModel

	Name     string ` db:"name" json:"name" gorm:"type:varchar(100);" `
	RowOrder int    ` db:"row_order" json:"row_order" gorm:"type:int(11);" `
}
