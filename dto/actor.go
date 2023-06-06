package dto

type Actor struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"not null"`
	Password string `gorm:"not null"`
	RoleID   uint   `gorm:"not null"`
}
