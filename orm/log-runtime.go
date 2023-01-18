package orm

type LogRuntime struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	GormModel

	ChassisNumber string `db:"chassis_number"  json:"chassis_number" gorm:"type:varchar(50)"`
}
