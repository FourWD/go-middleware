package orm

type LogUserVisit struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	GormModel

	DomainCode string ` db:"domain_code" json:"domain_code" gorm:"type:varchar(45);" `

	UserID    string `db:"user_id" json:"user_id" gorm:"type:varchar(36)"`
	KeyID     string `db:"key_id" json:"key_id" gorm:"type:varchar(36)"`
	VisitType string `db:"visit_type" json:"visit_type" gorm:"type:varchar(10)"`
}
