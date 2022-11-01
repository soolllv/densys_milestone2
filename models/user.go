package models

type User struct {
	Base
	Email             string `json:"email" gorm:"unique"`
	BloodGroup        uint   `json:"bloodGroup"`
	EmerContactNumber uint64 `json:"emerContactNum"`
	MaritalStatus     uint   `json:"maritalSts"`
}
