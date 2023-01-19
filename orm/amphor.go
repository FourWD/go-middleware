package orm

type Amphor struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	GormModel

	Name       string ` db:"name" json:"name" gorm:"type:varchar(150);" `
	NameEn     string ` db:"name_en" json:"name_en" gorm:"type:varchar(100);" `
	ProvinceID string ` db:"province_id" json:"province_id" gorm:"type:varchar(36);index;" `
	BTest      string ` db:"test" json:"test" gorm:"type:varchar(36);" `
}
