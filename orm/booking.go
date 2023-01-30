package orm

type Booking struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	GormModel

	DomainCode string `db:"domain_code"  json:"domain_code" gorm:"index;default:null;type:varchar(20)"`

	BookingNo       string `db:"booking_no" json:"booking_no" gorm:"type:varchar(12);index" ` //YYYYMMDD-4digit
	UserID          string `db:"user_id" json:"user_id" gorm:"type:varchar(36);"`
	VehicleID       string `db:"vehicle_id" json:"vehicle_id" gorm:"type:varchar(36);"`
	BookingStatusID string `db:"booking_status_id" json:"booking_status_id" gorm:"type:varchar(36);"`
	SaleID          string `db:"sale_id" json:"sale_id" gorm:"type:varchar(36);"`
}
