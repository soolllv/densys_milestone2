package models

type User struct {
	Base
	Email             string `json:"email" gorm:"unique"`
	BloodGroup        uint
	EmerContactNumber uint64
	MaritalStatus     uint
}
