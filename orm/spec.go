package orm

type Spec struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	GormModel

	VehicleID string `db:"vehicle_id" json:"vehicle_id" gorm:"type:varchar(36);index;"`

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

	DriveTypeID         string `db:"drive_type_id"  json:"drive_type_id" gorm:"default:null; type:varchar(36); comment:'ลักษณะการขับเคลื่อน'"`
	GearNumberID        string `db:"gear_number_id"  json:"gear_number_id" gorm:"default:null; type:varchar(36); comment:'จำนวนเกียร์'"`
	ManufacturerID      string `db:"manufacturer_id"  json:"manufacturer_id" gorm:"default:'ไม่มีข้อมูล'; type:varchar(36); comment:'โรงงานที่ผลิต'"`
	Co2ID               string `db:"co2_id" json:"co2_id" gorm:"default:'ไม่มีข้อมูล'; type:varchar(36); comment:'คาร์บอนไดออกไซด์';"`
	EngineTechID        string `db:"engine_tech_id,"  json:"engine_tech_id," gorm:"default:'ไม่มีข้อมูล'; type:varchar(36); comment:'เทคโนโลยีเครื่องยนต์'"`
	BoreStrokeID        string `db:"bore_stroke_id"  json:"bore_stroke_id" gorm:"default:'ไม่มีข้อมูล'; type:varchar(36); comment:'กระบอกสูบ-ระยะชัก'"`
	TorqueID            string `db:"torque_id"  json:"torque_id" gorm:"default:'ไม่มีข้อมูล'; type:varchar(36); comment:'แรงบิด'"`
	FuelCapacityID      string `db:"fuel_capacity_id"  json:"fuel_capacity_id" gorm:"default:'ไม่มีข้อมูล'; type:varchar(36); comment:'ปริมาณเชื้อเพลิง'"`
	Height              string `db:"height"  json:"height" gorm:"default:'ไม่มีข้อมูล'; type:varchar(50); comment:'ความสูง'"`
	Width               string `db:"width"  json:"width" gorm:"default:'ไม่มีข้อมูล'; type:varchar(50); comment:'ความกว้าง'"`
	FuelTypeID          string `db:"fuel_type_id"  json:"fuel_type_id" gorm:"default:null ; type:varchar(36); comment:'ประเภทเชื้อเพลิง'"`
	SeatingCapacityID   string `db:"seating_capacity_id"  json:"seating_capacity_id" gorm:"default:null ; type:varchar(36); comment:'ความจุที่นั่ง'"`
	FrontTyreSizeID     string `db:"front_tyre_size_id"  json:"front_tyre_size_id" gorm:"default:'ไม่มีข้อมูล'; type:varchar(36); comment:'ขนาดล้อหน้า'"`
	RearTyreSizeID      string `db:"rear_tyre_size_id"  json:"rear_tyre_size_id" gorm:"default:'ไม่มีข้อมูล'; type:varchar(36); comment:'ขนาดล้อหลัง'"`
	FrontTyreTypeID     string `db:"front_tyre_type_id"  json:"front_tyre_type_id" gorm:"default:'ไม่มีข้อมูล';type:varchar(36); comment:'ประเภทล้อหน้า'"`
	RearTyreTypeID      string `db:"rear_tyre_type_id"  json:"rear_tyre_type_id" gorm:"default:'ไม่มีข้อมูล';type:varchar(36); comment:'ประเภทล้อหลัง'"`
	FrontBrakesID       string `db:"front_brake_id"  json:"front_brake_id" gorm:"default:'ไม่มีข้อมูล';type:varchar(36); comment:'เบรกหน้า'"`
	RearBrakesID        string `db:"rear_brake_id"  json:"rear_brake_id" gorm:"default:'ไม่มีข้อมูล';type:varchar(36); comment:'เบรกหลัง'"`
	FrontSuspensionID   string `db:"front_suspension_id"  json:"front_suspension_id" gorm:"default:'ไม่มีข้อมูล';type:varchar(36); comment:'ระบบกันสะเทือนช่วงล่างด้านหน้า'"`
	RearSuspensionID    string `db:"rear_suspension_id"  json:"rear_suspension_id" gorm:"default:'ไม่มีข้อมูล';type:varchar(36); comment:'ระบบกันสะเทือนช่วงล่างด้านหลัง'"`
	SteeringID          string `db:"steering_id"  json:"steering_id" gorm:"default:null; type:varchar(36); comment:'พวงมาลัย'"`
	TransmissionTypeID  string `db:"transmission_type_id"  json:"transmission_type_id" gorm:"default:null; type:varchar(36); comment:'รูปแบบของเกียร์'"`
	TransmissionNameID  string `db:"transmission_name_id"  json:"transmission_name_id" gorm:"default:'ไม่มีข้อมูล'; type:varchar(36); comment:'ชื่อประเภทของเกียร์'"`
	RatedEconomyID      string `db:"rated_economy_id"  json:"rated_economy_id" gorm:"default:'ไม่มีข้อมูล'; type:varchar(36); comment:'อัตราประหยัดน้ำมัน'"`
	DoorID              string `db:"door_id"  json:"door_id" gorm:"default:null; type:varchar(36); comment:'จำนวนประตู'"`
	EngineArrangementID string `db:"engine_arrangement_id"  json:"engine_arrangement_id" gorm:"default:'ไม่มีข้อมูล'; type:varchar(36); comment:'ประเภทของเครื่องยนต์'"`
	HorsePowerID        string `db:"horse_power_id"  json:"horse_powe_idr" gorm:"default:'ไม่มีข้อมูล';type:varchar(36); comment:'แรงม้า'"`
	CompressionRatioID  string `db:"compression_ratio_id"  json:"compression_ratio_id" gorm:"default:'ไม่มีข้อมูล';type:varchar(36); comment:'อัตราส่วนกำลังอัด'"`
	WheelBaseID         string `db:"wheel_base_id"  json:"wheel_base_id" gorm:"default:'ไม่มีข้อมูล';type:varchar(36); comment:'ระยะล้อหน้าถีงหลัง'"`
	Length              string `db:"length"  json:"length" gorm:"default:'ไม่มีข้อมูล';type:varchar(50); comment:'ความยาว'"`
	Weight              string `db:"weight"  json:"weight" gorm:"default:'ไม่มีข้อมูล';type:varchar(50); comment:'น้ำหนัก'"`
	BootSpace           string `db:"boot_space"  json:"boot_space" gorm:"default:'ไม่มีข้อมูล';type:varchar(50); comment:'พื้นที่ความจุกระโปรงท้าย'"`
	SpareTyreID         string `db:"spare_tyre_id"  json:"spare_tyre_id" gorm:"default:null; type:varchar(36); comment:'ล้ออะไหล่'"`
	AssemblyID          string `db:"assembly_id"  json:"assembly_id" gorm:"default:'ไม่มีข้อมูล'; type:varchar(36); comment:'ประเทศที่ผลิต'"`
	WarrantyID          string `db:"warranty_id"  json:"warranty_id" gorm:"default:'ไม่มีข้อมูล'; type:varchar(36); comment:'การรับประกัน'"`
	ColorID             string `db:"color_id"  json:"color_id" gorm:"default:'ไม่มีข้อมูล'; type:varchar(36); comment:'สี'"`
	EngineCapacityID    string `db:"engine_capacity_id"  json:"engine_capacity_id" gorm:"default:'ไม่มีข้อมูล'; type:varchar(36); comment:'การรับประกัน'"`
}
