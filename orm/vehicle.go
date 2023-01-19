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
	DriveSystem       string    `db:"drive_system" json:"drive_system" gorm:"type:varchar(50) ; dafault:null "`
	GearTypeID        string    `db:"gear_type_id" json:"gear_type_id" gorm:"type:varchar(36);  "`
	ColorID           string    `db:"color_id" json:"color_id" gorm:"type:varchar(36);"`
	EngineCapacity    string    `db:"engine_capacity" json:"engine_capacity" gorm:"type:varchar(50) ; dafault:null "`
	EngineOilID       string    `db:"engine_oil_id" json:"engine_oil_id" gorm:"type:varchar(36)"`
	EngineSize        string    `db:"engine_size" json:"engine_size" gorm:"type:varchar(50) ; dafault:null "`
	Seat              int       `db:"seat" json:"seat" gorm:"type:int(11) ; dafault:null "`
	Mileage           int       `db:"mileage" json:"mileage" gorm:"type:int(11) "`
	LicensePlate      string    `db:"license_plate" json:"license_plate" gorm:"type:varchar(50) ; dafault:null "`
	LicenseProvinceID string    `db:"license_province_id" json:"license_province_id" gorm:"type:varchar(36) "`
	ImgStrFront       string    `db:"img_str_front" json:"img_str_front" gorm:"type:varchar(50) ; dafault:null "`
	ImgStrBack        string    `db:"img_str_back" json:"img_str_back" gorm:"type:varchar(50) ; dafault:null "`
	ImgStrRight       string    `db:"img_str_right" json:"img_str_right" gorm:"type:varchar(50) ; dafault:null "`
	ImgStrLeft        string    `db:"img_str_left" json:"img_str_left" gorm:"type:varchar(50) ; dafault:null "`
	ImgFrontLeft_45   string    `db:"img_front_left_45" json:"img_front_left_45" gorm:"type:varchar(50) ; dafault:null "`
	ImgFrontRight_45  string    `db:"img_front_right_45" json:"img_front_right_45" gorm:"type:varchar(50) ; dafault:null "`
	ImgBackLeft_45    string    `db:"img_back_left_45" json:"img_back_left_45" gorm:"type:varchar(50) ; dafault:null "`
	ImgBackRight_45   string    `db:"img_back_right_45" json:"img_back_right_45" gorm:"type:varchar(50) ; dafault:null "`
	ImgInBront        string    `db:"img_in_front" json:"img_in_front" gorm:"type:varchar(50) ; dafault:null "`
	ImgInBack         string    `db:"img_in_back" json:"img_in_back" gorm:"type:varchar(50) ; dafault:null "`
	ImgConsole        string    `db:"img_console" json:"img_console" gorm:"type:varchar(50) ; dafault:null "`
	ImgMileage        string    `db:"img_mileage" json:"img_mileage" gorm:"type:varchar(50) ; dafault:null "`
	ImgVehTools       string    `db:"img_veh_tools" json:"img_veh_tools" gorm:"type:varchar(50) ; dafault:null "`
	ImgSpareWheel     string    `db:"img_spare_wheel" json:"img_spare_wheel" gorm:"type:varchar(50) ; dafault:null "`
	ImgEngineRoom     string    `db:"img_engine_room" json:"img_engine_room" gorm:"type:varchar(50) ; dafault:null "`
	ImgOut360         string    `db:"img_out_360" json:"img_out_360" gorm:"type:varchar(50) ; dafault:null "`
	Price             int       `db:"price" json:"price" gorm:"type:int(11);  "`
	Url               string    `db:"url"  json:"url" gorm:"dafault:null;type:varchar(100)"`
	OnShelf           bool      `db:"on_shelf"  json:"on_shelf" gorm:"type:tinyint(1)"`
	OnShelfDate       time.Time `db:"on_shelf_date"  json:"on_shelf_date" gorm:"type:datetime"`
}
