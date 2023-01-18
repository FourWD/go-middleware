package orm

type VehicleType struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	GormModel

	Name   string `db:"name" json:"type_name"`
	NameEn string `db:"name_en" json:"type_name_en"`
}
