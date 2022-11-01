package models

type User struct {
	Base
	Email             string `json:"email" gorm:"unique"`
	Password          string `json:"password"`
	BloodGroup        uint
	EmerContactNumber uint64
	MaritalStatus     uint
}
