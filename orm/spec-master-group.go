package orm

type SpecMasterGroup struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	GormModel

	Name                  string `db:"name"  json:"name" gorm:"default:null;type:varchar(200)"`
	NameEn                string `db:"name_en"  json:"name_en" gorm:"default:null;type:varchar(200)"`
	SpecMasterGroupTypeID string `db:"spec_master_group_type_id" json:"spec_master_group_type_id" gorm:"type:varchar(36);index;"`
	// airbag group feature
}
