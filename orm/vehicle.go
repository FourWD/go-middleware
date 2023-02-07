package orm

import (
	"time"
)

type Vehicle struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	GormModel

	DomainCode string `db:"domain_code"  json:"domain_code" gorm:"type:varchar(50); index"`

	Code              string    `db:"code"  json:"code" gorm:"type:varchar(50) ; dafault:null ; index"`
	Source            string    `db:"source"  json:"source" gorm:"type:varchar(50) ; dafault:null "`
	Grade             string    `db:"grade" json:"grade" gorm:"type:varchar(50) ; dafault:null "`
	YearBu            string    `db:"year_bu" json:"year_bu" gorm:"type:varchar(4) ; dafault:null "`
	YearReg           string    `db:"year_reg" json:"year_reg" gorm:"type:varchar(4) ; dafault:null "`
	VehicleTypeID     string    `db:"vehicle_type_id" json:"vehicle_type_id" gorm:"type:varchar(36);"`
	VehicleBrandID    string    `db:"vehicle_brand_id" json:"vehicle_brand_id" gorm:"type:varchar(36) "`
	VehicleModelID    string    `db:"vehicle_model_id" json:"vehicle_model_id" gorm:"type:varchar(36);"`
	VehicleSubmodelID string    `db:"vehicle_submodel_id" json:"vehicle_submodel_id" gorm:"type:varchar(36); "`
	BranchID          string    `db:"branch_id" json:"branch_id" gorm:"type:varchar(36); "`
	Mileage           int       `db:"mileage" json:"mileage" gorm:"type:int(11) "`
	LicensePlate      string    `db:"license_plate" json:"license_plate" gorm:"type:varchar(50) ; dafault:null "`
	LicenseProvinceID string    `db:"license_province_id" json:"license_province_id" gorm:"type:varchar(36) "`
	Price             int       `db:"price" json:"price" gorm:"type:int(11);  "`
	Url               string    `db:"url"  json:"url" gorm:"dafault:null;type:varchar(100)"`
	OnShelf           int       `db:"on_shelf"  json:"on_shelf" gorm:"type:tinyint(1)"`
	OnShelfDate       time.Time `db:"on_shelf_date"  json:"on_shelf_date" gorm:"type:datetime"`
	Iview             int       `db:"iview" json:"iview" gorm:"type:int(11);  "`
}
