package orm

type ArticleGroup struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	GormModel

	DomainCode  string ` db:"domaincode" json:"domaincode" gorm:"type:varchar(100); index" `
	Code        string ` db:"code" json:"code" gorm:"type:varchar(100); index "`
	Description string ` db:"description" json:"description" gorm:"type:varchar(500);"`
}
