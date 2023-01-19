package orm

type Article struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	GormModel

	DomainCode string ` db:"domaincode" json:"domaincode" gorm:"type:varchar(40); " `
	GroupID    string ` db:"group_id" json:"group_id" gorm:"type:varchar(36);index" `
	Subject    string ` db:"subject" json:"subject" gorm:"type:varchar(400); " `
	Details    string ` db:"details" json:"details" gorm:"type:varchar(400);" `
	Image      string ` db:"image" json:"image" gorm:"type:varchar(50);" `
}
