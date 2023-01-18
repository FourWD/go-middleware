package orm

type ConfigGroup struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36); primaryKey"`
	GormModel

	Code        string ` db:"code" json:"code" gorm:"type:varchar(50); unique_index" `
	MasterTable string ` db:"master_table" json:"master_table" gorm:"type:varchar(50);index" `
	Description string ` db:"description" json:"description" gorm:"type:varchar(100)" `
}
