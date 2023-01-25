package orm

import (
	"time"
)

type Maintenance struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	GormModel

	VehicleID         string    ` db:"vehicle_id" json:"vehicle_id" gorm:"type:varchar(36);index" `
	LocationID        string    ` db:"location_id" json:"location_id" gorm:"type:varchar(36);index" `
	Mileage           int       ` db:"mileage" json:"mileage" gorm:"type:int(11);" `
	MaintenanceTypeID string    ` db:"maintenance_type_id" json:"maintenance_type_id" gorm:"type:varchar(36);index;" `
	MaintenanceDate   time.Time ` db:"maintenance_date" json:"maintenance_date" gorm:"type:datetime;" `
	Detail            string    ` db:"detail" json:"detail" gorm:"type:varchar(500);" `
}
