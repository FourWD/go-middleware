package orm

type SyncMaintenace struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	GormModel

	Status int    `db:"status"  json:"status" gorm:"default:null;type:int(1)"`
	Source string `db:"source"  json:"source" gorm:"index;default:null;type:varchar(6)"`
	ApiUrl string `db:"api_url"  json:"api_url" gorm:"type:varchar(255)"`
}
