package orm

type Config struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36); primaryKey"`
	GormModel

	DomainCode    string ` db:"domain_code" json:"domain_code" gorm:"type:varchar(100);" `
	ConfigGroupID string ` db:"config_group_id" json:"config_group_id" gorm:"type:varchar(36);index" `
	ConfigValue   string ` db:"config_value" json:"config_value" gorm:"type:varchar(100);" `
	// ConfigCode    string    ` db:"config_code" json:"config_code" gorm:"type:varchar(100);" `
	// SecretID      string    ` db:"secret_id" json:"secret_id" gorm:"type:varchar(100);" `
	// SecretKey     string    ` db:"secret_key" json:"secret_key" gorm:"type:varchar(100);" `
}
