package orm

type GearType struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	GormModel

	Name string ` db:"name" json:"name" gorm:"type:varchar(50);index " `
}
