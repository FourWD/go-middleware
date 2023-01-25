package orm

type Article struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	GormModel

	GroupID     string ` db:"group_id" json:"group_id" gorm:"type:varchar(36);index" `
	Subject     string ` db:"subject" json:"subject" gorm:"type:varchar(500); " `
	Description string ` db:"description" json:"description" gorm:"type:varchar(500); " `
	Detail      string ` db:"detail" json:"detail" gorm:"type:varchar(500);" ` // gorm to text
	Image       string ` db:"image" json:"image" gorm:"type:varchar(200);" `
	Video       string ` db:"video" json:"video" gorm:"type:varchar(200);" `
	YoutubeUrl  string ` db:"youtube_url" json:"youtube_url" gorm:"type:varchar(200);" `
	Url         string ` db:"url" json:"url" gorm:"type:varchar(100);" `
}
