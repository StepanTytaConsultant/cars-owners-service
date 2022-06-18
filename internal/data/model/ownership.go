package model

type Ownership struct {
	ID            int64   `gorm:"column:id" csv:"id"`
	CarManufactur *string `gorm:"column:car_manufactur" csv:"car_manufactur"`
	CarModel      *string `gorm:"column:car_model" csv:"car_model"`
	CarModelYear  *int32  `gorm:"column:car_model_year" csv:"car_model_year"`
	FirstName     *string `gorm:"column:first_name" csv:"first_name"`
	LastName      *string `gorm:"column:last_name" csv:"last_name"`
	Email         *string `gorm:"column:email" csv:"email"`
	Gender        *string `gorm:"column:gender" csv:"gender"`
	Address       *string `gorm:"column:address" csv:"address"`
}
