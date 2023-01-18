package orm

type Menu struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	GormModel

	MenuGroupID string `db:"menu_group_id" json:"menu_group_id" gorm:"type:varchar(36);"`
	Icon        string ` db:"icon" json:"icon" gorm:"type:varchar(45);" `
	Name        string ` db:"name" json:"name" gorm:"type:varchar(45);" `
	UrlPath     string ` db:"path" json:"path" gorm:"type:varchar(255);" `
	MenuID      string ` db:"menu_id" json:"menu_id" gorm:"type:varchar(36);" `
	RowOrder    int    ` db:"row_order" json:"row_order" gorm:"type:int;" `
}
