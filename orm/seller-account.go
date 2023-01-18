package orm

type SellerAccount struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	GormModel

	Username string `db:"username"  json:"username" gorm:"index;type:varchar(50)"`
	Password string `db:"password"  json:"password" gorm:"type:varchar(50)"`
}
