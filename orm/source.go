package orm

import (
	"time"
)

type Source struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	GormModel

	Name                 string    `db:"name"  json:"name" gorm:"index;default:null;type:varchar(200)"`
	ShortName            string    `db:"short_name"  json:"short_name" gorm:"index;default:null;type:varchar(45)"`
	NameInitials         string    `db:"name_initials"  json:"name_initials" gorm:"index;default:null;type:varchar(50)"`
	RegistrationNumber   string    `db:"registration_number"  json:"registration_number" gorm:"index;default:null;type:varchar(20)"`
	CompanyName          string    `db:"company_name"  json:"company_name" gorm:"index;default:null;type:varchar(200)"`
	Tel                  string    `db:"tel"  json:"tel" gorm:"index;default:null;type:varchar(10)"`
	Email                string    `db:"email"  json:"email" gorm:"index;default:null;type:varchar(45)"`
	Address              string    `db:"address"  json:"address" gorm:"index;default:null;type:varchar(250)"`
	ProvinceID           string    `db:"province_id"  json:"province_id" gorm:"type:varchar(36);index;"`
	District             int       `db:"district"  json:"district" gorm:"index;default:null;type:int(11)"`
	Subdistrict          int       `db:"subdistrict"  json:"subdistrict" gorm:"index;default:null;type:int(11)"`
	Postcode             int       `db:"postcode"  json:"postcode" gorm:"index;default:null;type:int(11)"`
	ContractSigningDate  time.Time `db:"contract_signing_date"  json:"contract_signing_date" gorm:"index;default:null;type:datetime"`
	ContractClostingDate time.Time `db:"contract_closting_date"  json:"contract_closting_date" gorm:"index;default:null;type:datetime"`
}
