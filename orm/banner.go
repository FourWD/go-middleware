package orm

type Banner struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	GormModel

	BannerGroupID string `db:"banner_grounp_id"  json:"banner_grounp_id" gorm:"type:varchar(36);index"`
	ImagePath     string `db:"image_path"  json:"image_path" gorm:"dafault:null;type:varchar(100)"`
	Url           string `db:"url"  json:"url" gorm:"dafault:null;type:varchar(100)"`
	YoutubeIframe string `db:"youtube_iframe" json:"youtube_iframe" gorm:"dafault:null;type:varchar(100)"`
	Description   string `db:"description"  json:"description" gorm:"dafault:null;type:varchar(100)"`
}
