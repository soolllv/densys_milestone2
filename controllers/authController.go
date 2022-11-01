package controllers

import (
	"strconv"

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

func EditDoctor(c *fiber.Ctx) error {
	// database.DB.First(c.Params("id"));
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var curDoc models.Doctor
	intVal, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}
	curDoc.ID = uint(intVal)
	database.DB.Take(&curDoc)

	curDoc.Name = data["name"]
	curDoc.Surname = data["surname"]
	curDoc.MiddleName = data["middleName"]
	curDoc.DataOfBirth = data["dateOfBirth"]
	curDoc.IIN_Number = data["iinNum"]
	curDoc.ID_Number = data["idNum"]
	curDoc.ContactNumber = data["contactNum"]
	curDoc.Address = data["address"]
	curDoc.DepartmentID = data["departmentId"]
	curDoc.SpecDetailsID = data["specDetailsID"]
	curDoc.Experience = data["experience"]
	curDoc.Photo = data["photo"]
	curDoc.Category = data["category"]
	curDoc.PriceOfApp = data["priceOfApp"]
	curDoc.Degree = data["degree"]
	curDoc.Rating = data["rating"]

	database.DB.Save(&curDoc)
	return c.JSON(curDoc)
}

func GetDoctor(c *fiber.Ctx) error {

	// var curDoc map[string]string
	var curDoc models.Doctor
	intVal, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}
	curDoc.ID = uint(intVal)
	database.DB.Take(&curDoc)

	return c.JSON(curDoc)

}

func GetUser(c *fiber.Ctx) error {

	// var curDoc map[string]string
	var curUser models.User
	intVal, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}
	curUser.ID = uint(intVal)
	database.DB.Take(&curUser)

	return c.JSON(curUser)

}
func EditUser(c *fiber.Ctx) error {
	// database.DB.First(c.Params("id"));
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var curUser models.User
	intVal, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}
	curUser.ID = uint(intVal)
	database.DB.Take(&curUser)

	curUser.Name = data["name"]
	curUser.Surname = data["surname"]
	curUser.MiddleName = data["middleName"]
	curUser.DataOfBirth = data["dateOfBirth"]
	curUser.IIN_Number = data["iinNum"]
	curUser.ID_Number = data["idNum"]
	curUser.ContactNumber = data["contactNum"]
	curUser.Address = data["address"]
	curUser.Email = data["email"]
	curUser.BloodGroup = data["bloodGroup"]
	curUser.EmerContactNumber = data["emerContactNum"]
	curUser.MaritalStatus = data["maritalSts"]

	database.DB.Save(&curUser)
	return c.JSON(curUser)
}

func ListDoctors(c *fiber.Ctx) error {
	var docs []models.Doctor
	database.DB.Find(&docs)

	return c.JSON(docs)
}

func ListUsers(c *fiber.Ctx) error {
	var users []models.User
	database.DB.Find(&users)

	return c.JSON(users)
}

func Hello(c *fiber.Ctx) error {
	return c.SendString("SUP mate")
}
