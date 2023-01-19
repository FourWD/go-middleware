package orm

type User struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	GormModel

	Username string `db:"username" json:"username" gorm:"type:varchar(100)"`
	Password string `db:"password" json:"password" gorm:"type:varchar(100)"`
	Fullname string `db:"fullname" json:"fullname" gorm:"type:varchar(100)"`
	Email    string `db:"email" json:"email" gorm:"type:varchar(100)"`
}
