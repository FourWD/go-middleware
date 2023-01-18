package orm

type Finance struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	GormModel

	Name             string ` db:"name" json:"name" gorm:"type:varchar(50); unique_index " `
	ContactName      string ` db:"contact_name" json:"contact_name" gorm:"type:varchar(50);" `
	ContactTel       string ` db:"contact_tel" json:"contact_tel" gorm:"type:varchar(50);" `
	ImgSize1         string ` db:"img_size1" json:"img_size1" gorm:"type:varchar(100);" `
	ImgSize2         string ` db:"img_size2" json:"img_size2" gorm:"type:varchar(100);" `
	NameEn           string ` db:"name_en" json:"name_en" gorm:"type:varchar(150);" `
	PrefinanceStatus uint   ` db:"prefinance_status" json:"prefinance_status" gorm:"type:int;" `
}
