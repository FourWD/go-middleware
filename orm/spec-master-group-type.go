package orm

type SpecMasterGroupType struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	GormModel

	Name string `db:"name"  json:"name" gorm:"default:null;type:varchar(200)"`
	// feature & spec
}
