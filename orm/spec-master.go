package orm

type SpecMaster struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	GormModel

	SpecMasterGroupID string `db:"spec_master_group_id" json:"spec_master_group_id" gorm:"type:varchar(36);index;"`

	Name     string `db:"name" json:"name" gorm:"default:null;type:varchar(200)"`
	NameEn   string `db:"name_en"  json:"name_en" gorm:"default:null;type:varchar(200)"`
	RowOrder int64  `db:"row_order" json:"row_order" gorm:"default:null; type:int"`
	// select option

}
