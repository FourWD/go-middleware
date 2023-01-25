package orm

type VehicleBrand struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	GormModel

	DomainCode string ` db:"domain_code" json:"domain_code" gorm:"type:varchar(45);" `

	Name   string `db:"name" json:"type_name"`
	NameEn string `db:"name_en" json:"type_name_en"`
}
