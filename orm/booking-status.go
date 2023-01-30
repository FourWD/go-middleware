package orm

type BookingStatus struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	GormModel

	Name string ` db:"name" json:"name" gorm:"type:varchar(100);index" `
}
