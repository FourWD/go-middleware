package orm

type Spec struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	GormModel

	VehicleID string `db:"vehicle_id" json:"vehicle_id" gorm:"type:varchar(36);index;"`

	// Features
	AirbagID            string `db:"airbag_id"  json:"airbag_id" gorm:"default:null;type:varchar(36); comment:'แอร์แบ็ค'"`
	BlindSpotSystemID   string `db:"blind_spot_system_id"  json:"blind_spot_system_id" gorm:"default:null;type:varchar(36) ; comment:'ระบบกล้อง'" `
	BluetoothID         string `db:"bluetooth_id"  json:"bluetooth_id" gorm:"default:null;type:varchar(36); comment:'บลูธูท'"`
	BootOperationID     string `db:"boot_operation_id"  json:"boot_operation_id" gorm:"default:null;type:varchar(36); comment:'กดปุ่มสตาร์ทรถ'"`
	SunroofID           string `db:"sunroof_id"  json:"sunroof_id" gorm:"default:null;type:varchar(36); comment:'หลังคารับแสง'"`
	ReverseCameraID     string `db:"reverse_camera_id"  json:"reverse_camera_id" gorm:"default:null;type:varchar(36); comment:'กล้องถอยหลัง'"`
	CruiseControlID     string `db:"cruise_control_id"  json:"cruise_control_id" gorm:"default:null;type:varchar(36); comment:'ระบบคุมความเร็วอัตโนมัติ'"`
	ParkingSensorRearID string `db:"parking_sensor_rear_id"  json:"parking_sensor_rear_id" gorm:"default:null;type:varchar(36); comment:'เซ็นเซอร์จอดด้านหลัง'"`
	PowerWindowID       string `db:"power_window_id"  json:"power_window_id" gorm:"default:null;type:varchar(36); comment:'หน้าต่างไฟฟ้า'"`
	ChildSafetyLockID   string `db:"child_safety_lock_id"  json:"child_safety_lock_id" gorm:"default:null;type:varchar(36); comment:'ติดตั้งคาร์ซีทสำหรับเด็ก'"`

	// Specs
	DriveTypeID        string `db:"drive_type_id"  json:"drive_type_id" gorm:"default:null; type:varchar(36) comment:'ลักษณะการขับเคลื่อน'"`
	GearNumberID       string `db:"gear_number_id"  json:"gear_number_id" gorm:"default:null; type:varchar(36) comment:'จำนวนเกียร์'"`
	ManufacturerID     string `db:"manufacturer"  json:"manufacturer" gorm:"default:'ไม่มีข้อมูล'; type:varchar(100) comment:'โรงงานที่ผลิต'"`
	Co2                string `db:"co2" json:"co2" gorm:"default:'ไม่มีข้อมูล'; type:varchar(100); comment:'คาร์บอนไดออกไซด์';"`
	EngineTech         string `db:"engine_tech,"  json:"engine_tech," gorm:"default:'ไม่มีข้อมูล'; type:varchar(100) comment:'เทคโนโลยีเครื่องยนต์'"`
	BoreStroke         string `db:"bore_stroke"  json:"bore_stroke" gorm:"default:'ไม่มีข้อมูล'; type:varchar(100) comment:'กระบอกสูบ-ระยะชัก'"`
	Torque             string `db:"torque"  json:"torque" gorm:"default:'ไม่มีข้อมูล'; type:varchar(100) comment:'แรงบิด'"`
	FuelCapacity       string `db:"fuel_capacity"  json:"fuel_capacity" gorm:"default:'ไม่มีข้อมูล'; type:varchar(100) comment:'ปริมาณเชื้อเพลิง'"`
	Height             string `db:"height"  json:"height" gorm:"default:'ไม่มีข้อมูล'; type:varchar(50) comment:'ความสูง'"`
	Width              string `db:"width"  json:"width" gorm:"default:'ไม่มีข้อมูล'; type:varchar(50) comment:'ความกว้าง'"`
	FuelTypeID         string `db:"fuel_type_id"  json:"fuel_type_id" gorm:"default:null ; type:varchar(36) comment:'ประเภทเชื้อเพลิง'"`
	SeatingCapacityID  string `db:"seating_capacity_id"  json:"seating_capacity_id" gorm:"default:null ; type:varchar(36) comment:'ความจุที่นั่ง'"`
	FrontTyreSize      string `db:"front_tyre_size"  json:"front_tyre_size" gorm:"default:'ไม่มีข้อมูล'; type:varchar(100) comment:'ขนาดล้อหน้า'"`
	RearTyreSize       string `db:"rear_tyre_size"  json:"rear_tyre_size" gorm:"default:'ไม่มีข้อมูล'; type:varchar(100) comment:'ขนาดล้อหลัง'"`
	FrontTyreType      string `db:"front_tyre_type"  json:"front_tyre_type" gorm:"default:'ไม่มีข้อมูล';type:varchar(100) comment:'ประเภทล้อหน้า'"`
	RearTyreType       string `db:"rear_tyre_type"  json:"rear_tyre_type" gorm:"default:'ไม่มีข้อมูล';type:varchar(100) comment:'ประเภทล้อหลัง'"`
	FrontBrakes        string `db:"front_brake"  json:"front_brake" gorm:"default:'ไม่มีข้อมูล';type:varchar(100) comment:'เบรกหน้า'"`
	RearBrakes         string `db:"rear_brake"  json:"rear_brake" gorm:"default:'ไม่มีข้อมูล';type:varchar(100) comment:'เบรกหลัง'"`
	FrontSuspension    string `db:"front_suspension"  json:"front_suspension" gorm:"default:'ไม่มีข้อมูล';type:varchar(100) comment:'ระบบกันสะเทือนช่วงล่างด้านหน้า'"`
	RearSuspension     string `db:"rear_suspension"  json:"rear_suspension" gorm:"default:'ไม่มีข้อมูล';type:varchar(100) comment:'ระบบกันสะเทือนช่วงล่างด้านหลัง'"`
	SteeringID         string `db:"steering_id"  json:"steering_id" gorm:"default:null; type:varchar(36) comment:'พวงมาลัย'"`
	TransmissionTypeID string `db:"transmission_type_id"  json:"transmission_type_id" gorm:"default:null; type:varchar(36) comment:'รูปแบบของเกียร์'"`
	TransmissionName   string `db:"transmission_name"  json:"transmission_name" gorm:"default:'ไม่มีข้อมูล'; type:varchar(100) comment:'ชื่อประเภทของเกียร์'"`
	RatedEconomy       string `db:"rated_economy"  json:"rated_economy" gorm:"default:'ไม่มีข้อมูล'; type:varchar(100) comment:'อัตราประหยัดน้ำมัน'"`
	DoorID             string `db:"door_id"  json:"door_id" gorm:"default:null; type:varchar(36) comment:'จำนวนประตู'"`
	EngineArrangement  string `db:"engine_arrangement"  json:"engine_arrangement" gorm:"default:'ไม่มีข้อมูล'; type:varchar(100) comment:'ประเภทของเครื่องยนต์'"`
	HorsePower         string `db:"horse_power"  json:"horse_power" gorm:"default:'ไม่มีข้อมูล';type:varchar(50) comment:'แรงม้า'"`
	CompressionRatio   string `db:"compression_ratio"  json:"compression_ratio" gorm:"default:'ไม่มีข้อมูล';type:varchar(50) comment:'อัตราส่วนกำลังอัด'"`
	WheelBase          string `db:"wheel_base"  json:"wheel_base" gorm:"default:'ไม่มีข้อมูล';type:varchar(50) comment:'ระยะล้อหน้าถีงหลัง'"`
	Length             string `db:"length"  json:"length" gorm:"default:'ไม่มีข้อมูล';type:varchar(50) comment:'ความยาว'"`
	Weight             string `db:"weight"  json:"weight" gorm:"default:'ไม่มีข้อมูล';type:varchar(50) comment:'น้ำหนัก'"`
	BootSpace          string `db:"boot_space"  json:"boot_space" gorm:"default:'ไม่มีข้อมูล';type:varchar(100) comment:'พื้นที่ความจุกระโปรงท้าย'"`
	SpareTyreID        string `db:"spare_tyre_id"  json:"spare_tyre_id" gorm:"default:null; type:varchar(36) comment:'ล้ออะไหล่'"`
	Assembly           string `db:"assembly"  json:"assembly" gorm:"default:'ไม่มีข้อมูล'; type:varchar(50) comment:'ประเทศที่ผลิต'"`
	Warranty           string `db:"warranty"  json:"warranty" gorm:"default:'ไม่มีข้อมูล'; type:varchar(50) comment:'การรับประกัน'"`
	ColorID            string `db:"color_id"  json:"color_id" gorm:"default:'ไม่มีข้อมูล'; type:varchar(50) comment:'สี'"`
	EngineCapacityID   string `db:"engine_capacity_id"  json:"engine_capacity_id" gorm:"default:'ไม่มีข้อมูล'; type:varchar(50) comment:'การรับประกัน'"`
}
