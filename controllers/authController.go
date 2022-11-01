package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/soolllv/densys_milestone2/database"
	"github.com/soolllv/densys_milestone2/models"
)

func RegisterUser(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	user := models.User{
		Name:              data["name"],
		Surname:           data["surname"],
		MiddleName:        data["middleName"],
		DataOfBirth:       data["dateOfBirth"],
		IIN_Number:        data["iinNum"],
		ID_Number:         data["idNum"],
		ContactNumber:     data["contactNum"],
		Address:           data["address"],
		Email:             data["email"],
		BloodGroup:        data["bloodGroup"],
		EmerContactNumber: data["emerContactNum"],
		MaritalStatus:     data["maritalSts"],
	}

	database.DB.Create(&user)

	return c.JSON(user)
}

func RegisterDoctor(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	doctor := models.Doctor{
		Name:          data["name"],
		Surname:       data["surname"],
		MiddleName:    data["middleName"],
		DataOfBirth:   data["dateOfBirth"],
		IIN_Number:    data["iinNum"],
		ID_Number:     data["idNum"],
		ContactNumber: data["contactNum"],
		Address:       data["address"],
		DepartmentID:  data["departmentId"],
		SpecDetailsID: data["specDetailsID"],
		Experience:    data["experience"],
		Photo:         data["photo"],
		Category:      data["category"],
		PriceOfApp:    data["priceOfApp"],
		Degree:        data["degree"],
		Rating:        data["rating"],
	}

	database.DB.Create(&doctor)

	return c.JSON(doctor)
}

func Hello(c *fiber.Ctx) error {
	return c.SendString("SUP mate")
}
