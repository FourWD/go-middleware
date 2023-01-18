package orm

type BannerGroup struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	GormModel

	DomainCode string `db:"domain_code"  json:"domain_code" gorm:"index;default:null;type:varchar(20)"`

	Code string `db:"code"  json:"code" gorm:"unique;index;default:null;type:varchar(45)"`
	Name string `db:"name" json:"name" gorm:"not null;type:varchar(50)"`
}
