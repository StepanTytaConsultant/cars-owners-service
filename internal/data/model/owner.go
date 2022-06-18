package model

type Owner struct {
	ID        int64   `gorm:"column:id"`
	FirstName *string `gorm:"column:first_name"`
	LastName  *string `gorm:"column:last_name"`
	Email     *string `gorm:"column:email"`
	Gender    *string `gorm:"column:gender"`
	Address   *string `gorm:"column:address"`
	City      *string `gorm:"column:city"`
}

func (Owner) TableName() string {
	return "owner"
}
