package model

import "time"

type Employee struct {
	ID             int `gorm:"primaryKey"`
	CompanyID      int `gorm:"foreignKey:CompanyRefer"`
	Company        Company
	Passport       string    `gorm:"uniqueIndex;type:varchar(255);not null"`
	Surname        string    `gorm:"type:varchar(255); not null"`
	Firstname      string    `gorm:"type:varchar(255); not null"`
	PatronymicName string    `gorm:"type:varchar(255); not null"`
	Position       string    `gorm:"type:varchar(255); not null"`
	PhoneNumber    string    `gorm:"type:varchar(255); not null"`
	Address        string    `gorm:"type:varchar(255); not null"`
	CreatedAt      time.Time `gorm:"not null"`
	UpdatedAt      time.Time `gorm:"not null"`
}
