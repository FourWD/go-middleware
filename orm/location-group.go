package orm

type LocationGroup struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	GormModel

	DomainCode string ` db:"domain_code" json:"domain_code" gorm:"type:varchar(100);" `

	Name string ` db:"name" json:"name" gorm:"type:varchar(150);" `
}
