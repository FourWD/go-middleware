package orm

type Branch struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	GormModel

	DomainCode string `db:"domain_code"  json:"domain_code" gorm:"index;default:null;type:varchar(20)"`

	Name     string  ` db:"name" json:"name" gorm:"type:varchar(100);index" `
	Address  string  ` db:"address" json:"address" gorm:"type:varchar(250);" `
	Phone    string  ` db:"phone" json:"phone" gorm:"type:varchar(100);" `
	WorkHour string  ` db:"work_hour" json:"work_hour" gorm:"type:varchar(200);" `
	Email    string  ` db:"email" json:"email" gorm:"type:varchar(100);" `
	Image    string  ` db:"image" json:"image" gorm:"type:text" `
	Lat      float64 ` db:"lat" json:"lat" gorm:"type:decimal(11,7)" `
	Long     float64 ` db:"long" json:"long" gorm:"type:decimal(11,7)" `
	RowOrder int     ` db:"row_order" json:"row_order" gorm:"type:int(11);index" `
}
