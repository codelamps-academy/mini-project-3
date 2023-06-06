package dto

// Struct untuk tabel RegisterApproval
type RegisterApproval struct {
	ID           uint   `gorm:"primaryKey"`
	AdminID      uint   `gorm:"not null"`
	SuperAdminID uint   `gorm:"not null"`
	Status       string `gorm:"not null"`
}
