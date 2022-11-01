package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/soolllv/densys_milestone2/models"
	// "github.com/soolllv/densys_milestone2/database"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	user := models.User{
		Base.Name:         data["name"],
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

	// database.DB.Create(&user)

	return c.JSON(user)
}

func Hello(c *fiber.Ctx) error {
	return c.SendString("SUP mate")
}
