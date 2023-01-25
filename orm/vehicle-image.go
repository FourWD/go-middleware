package orm

type VehicleImage struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	GormModel

	Name                string `db:"name"  json:"name" gorm:"not null;type:varchar(50)"`
	Image               string `db:"Image"  json:"Image" gorm:"type:varchar(100)"`
	RowOrder            string `db:"row_order"  json:"row_order" gorm:"type:varchar(100)"`
	VehicleImageGroupID string `db:"vehicle_image_group_id"  json:"vehicle_image_group_id" gorm:"type:varchar(36)"`
}
