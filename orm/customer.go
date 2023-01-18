package orm

type CUstomer struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36); primaryKey"`
	GormModel

	DomainCode string ` db:"domain_code" json:"domain_code" gorm:"type:varchar(50); index" `

	Alias       string ` db:"alias" json:"alias" gorm:"type:varchar(50);" `
	AvatarImage string ` db:"avatar_image" json:"avatar_image" gorm:"type:varchar(200);" `
	Firstname   string ` db:"firstname" json:"firstname" gorm:"type:varchar(100);" `
	Lastname    string ` db:"lastname" json:"lastname" gorm:"type:varchar(100);" `
	CitizenID   string ` db:"citizen_id" json:"citizen_id" gorm:"type:varchar(13);" `
	Email       string ` db:"email" json:"email" gorm:"type:varchar(100);" `
	LineID      string ` db:"line_id" json:"line_id" gorm:"type:varchar(50);" `
	FacebookID  string ` db:"facebook_id" json:"facebook_id" gorm:"type:varchar(50);" `
}
