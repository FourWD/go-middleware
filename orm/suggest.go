package orm

type Suggest struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	GormModel

	SuggestBy     string `db:"suggest_by"  json:"suggest_by" gorm:"index;default:null;type:varchar(100)"`
	SuggestTypeID string `db:"suggest_type_id"  json:"suggest_type_id" gorm:"type:varchar(36);index"`
	RowOrder      int    `db:"row_order"  json:"row_order" gorm:"type:int(11)"`
}
