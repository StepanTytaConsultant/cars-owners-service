package model

type Car struct {
	ID            int64   `gorm:"column:id" csv:"id"`
	CarManufactur *string `gorm:"column:car_manufactur" csv:"car_manufactur"`
	CarModel      *string `gorm:"column:car_model" csv:"car_model"`
	CarModelYear  *int32  `gorm:"column:car_model_year" csv:"car_model_year"`
	OwnerID       int64   `gorm:"column:owner_id" csv:"onwer_id"`
}

func (Car) TableName() string {
	return "car"
}
