package models

type Doctor struct {
	Base
	DepartmentID  uint
	SpecDetailsID uint
	Experience    uint
	Photo         string
	Category      string
	PriceOfApp    uint
	Degree        string
	Rating        uint
}
