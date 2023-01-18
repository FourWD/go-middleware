package orm

import (
	"time"

	"gorm.io/gorm"
)

type Promotion struct {
	ID        uint64         `db:"id" json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `db:"created_at" json:"created_at" gorm:"<-:create"`
	UpdatedAt time.Time      `db:"updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `db:"deleted_at" json:"deleted_at" gorm:"index"`
	Code      string         `db:"code" json:"code" gorm:"not null;index;unique;type:varchar(2)"`

	Name          string    `db:"name"  json:"name" gorm:"default:null;type:varchar(255);index"`
	Image         string    `db:"image"  json:"image" gorm:"default:null;type:varchar(255)"`
	ImageMobile   string    `db:"image_mobile"  json:"image_mobile" gorm:"default:null;type:varchar(255)"`
	IsActive      bool      `db:"_isActive"  json:"_isActive" gorm:"default:null;type:tinyint(1);index"`
	ShowHome      bool      `db:"show_home"  json:"show_home" gorm:"default:null;type:tinyint(1);index"`
	ShowPromotion bool      `db:"show_promotion"  json:"show_promotion" gorm:"default:null;type:tinyint(1);index"`
	FinID         int       `db:"fin_id"  json:"fin_id" gorm:"default:null;type:int(11);index"`
	PromotionSeq  int       `db:"promotion_seq"  json:"promotion_seq" gorm:"default:null;type:int(11)"`
	StartDatetime time.Time `db:"start_datetime"  json:"start_datetime" gorm:"default:null;type:datetime;index"`
	EndDatetime   time.Time `db:"end_datetime"  json:"end_datetime" gorm:"default:null;type:datetime;index"`
}

//string `gorm:"type:varchar(36);primary_key;

//ID        int            `db:"id" json:"id" gorm:"primaryKey"`
