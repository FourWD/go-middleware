package orm

type Spec struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	GormModel

	VehicleID string `db:"vehicle_id" json:"vehicle_id" gorm:"type:varchar(36);index;"`

	AirbagID          string `db:"airbag_id"  json:"airbag_id" gorm:"default:null;type:varchar(36); comment:'แอร์แบ็ค'"`
	BlindSpotSystemID string `db:"blind_spot_system_id"  json:"blind_spot_system_id" gorm:"default:null;type:varchar(36) ; comment:'ระบบกล้อง'" `

	DriveTypeID string `db:"drive_type_id"  json:"drive_type_id" gorm:"default:null;type:varchar(36)"`
}
