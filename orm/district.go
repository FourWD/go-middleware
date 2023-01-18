package orm

type District struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	GormModel

	ZipCode  uint   ` db:"zip_code" json:"zip_code" gorm:"type:int; " `
	Name     string ` db:"name" json:"name" gorm:"type:varchar(150);" `
	NameEn   string ` db:"name_en" json:"name_en" gorm:"type:varchar(150);" `
	AmphorID string ` db:"amphor_id" json:"amphor_id" gorm:"type:varchar(36);" `
}
