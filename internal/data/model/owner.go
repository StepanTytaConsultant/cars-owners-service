package model

type Owner struct {
	ID        int64   `gorm:"column:id" csv:"id"`
	FirstName *string `gorm:"column:first_name" csv:"first_name"`
	LastName  *string `gorm:"column:last_name" csv:"last_name"`
	Email     *string `gorm:"column:email" csv:"email"`
	Gender    *string `gorm:"column:gender" csv:"gender"`
	Address   *string `gorm:"column:address" csv:"address"`
	City      *string `gorm:"column:city" csv:"city"`
}

func (Owner) TableName() string {
	return "owner"
}
