package orm

type RequestCallback struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	GormModel

	DomainCode string ` db:"domain_code" json:"domain_code" gorm:"type:varchar(45);" `

	CustomerID string `db:"customer_id" json:"customer_id" gorm:"type:varchar(36)"`
	Firstname  string `db:"firstname" json:"firstname" gorm:"type:varchar(50)"`
	Lastname   string `db:"lastname" json:"lastname" gorm:"type:varchar(50)"`
	Phone      string `db:"phone" json:"phone" gorm:"type:varchar(50)"`
	Email      string `db:"email" json:"email" gorm:"type:varchar(50)"`
	IsCallback string `db:"is_callback" json:"is_callback" gorm:"type:boolean"`
	CallbackBy string `db:"callback_by" json:"callback_by" gorm:"type:datetime"`
	CallbackAt string `db:"callback_at" json:"callback_at" gorm:"type:datetime"`
}
