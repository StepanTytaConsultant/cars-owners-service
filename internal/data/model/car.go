package model

type Car struct {
	ID            int64   `gorm:"column:id"`
	CarManufactur *string `gorm:"column:car_manufactur"`
	CarModel      *string `gorm:"column:car_model"`
	CarModelYear  *int32  `gorm:"column:car_model_year"`
	OwnerID       int64   `gorm:"column:owner_id"`
}

func (Car) TableName() string {
	return "car"
}
