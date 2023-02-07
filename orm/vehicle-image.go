package orm

type VehicleImage struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	GormModel

	VehicleImageGroupID string `db:"vehicle_image_group_id"  json:"vehicle_image_group_id" gorm:"type:varchar(36)"`
	VehicleID           string `db:"vehicle_id"  json:"vehicle_id" gorm:"not null; type:varchar(50)"`
	ImagePath           string `db:"image_path"  json:"image_path" gorm:"dafault:null; type:varchar(100)"`
	Description         string `db:"description"  json:"description" gorm:"dafault:null; type:varchar(100)"`
	IsMain              bool   `db:"is_main"  json:"is_main" gorm:"type:tinyint(1)"`
	RowOrder            string `db:"row_order"  json:"row_order" gorm:"type:varchar(100)"`
}
