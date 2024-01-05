package model

import "time"

type Company struct {
	ID          int `gorm:"primaryKey"`
	UserID      int `gorm:"foreignKey:UserRefer"`
	User        User
	CompanyName string    `gorm:"type:varchar(255);not null"`
	DirectorFIO string    `gorm:"type:varchar(255); not null"`
	Email       string    `gorm:"uniqueIndex;not null"`
	Address     string    `gorm:"not null"`
	Website     string    `gorm:"type:varchar(255);not null"`
	CreatedAt   time.Time `gorm:"not null"`
	UpdatedAt   time.Time `gorm:"not null"`
}
