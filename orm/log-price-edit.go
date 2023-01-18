package orm

type LogPriceEdit struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	GormModel

	VehicleID string `db:"vehicle_id"  json:"vehicle_id" gorm:"type:varchar(36)"`
	OldPrice  int    `db:"old_price"  json:"old_price" gorm:"type:int(11)"`
	NewPrice  int    `db:"new_price"  json:"new_price" gorm:"type:int(11)"`
}
