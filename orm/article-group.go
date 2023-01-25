package orm

type ArticleGroup struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	GormModel

	DomainCode string ` db:"domaincode" json:"domaincode" gorm:"type:varchar(40); " `

	Name string ` db:"name" json:"name" gorm:"type:varchar(200);" `
}
