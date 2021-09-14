package handlers

import (
	"kitsurest/database"
	"kitsurest/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// GET

// Get user by id
func GetUserById(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	user := new(models.User)

	if err := db.First(&user, id).Error; err != nil {
		return c.Status(http.StatusServiceUnavailable).JSON(ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Success get user by ID.",
		Data:    *user,
	})
}

// Get user by userId
func GetUserByUserId(c *fiber.Ctx) error {
	userId := c.Params("userId")
	db := database.DBConn

	user := new(models.User)

	if err := db.Where("users.user_id = ?", userId).First(&user).Error; err != nil {
		return c.Status(http.StatusServiceUnavailable).JSON(ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Success get user by UserId.",
		Data:    *user,
	})
}

// POST

// Create a new user
func CreateNewUser(c *fiber.Ctx) error {
	db := database.DBConn

	user := new(models.User)

	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	db.Create(user)

	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Success register an user",
		Data:    *user,
	})
}

// PUT

// Update user by id
func UpdateUserById(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")

	user := new(models.User)

	if err := db.First(&user, id).Error; err != nil {
		return c.Status(http.StatusServiceUnavailable).JSON(ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	if err := UpdateUser(c, user); err != nil {
		return c.Status(http.StatusServiceUnavailable).JSON(ResponseHTTP{
			Success: false,
			Message: "Couldn't update the user",
			Data:    nil,
		})
	}

	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Success update user by Id.",
		Data:    *user,
	})
}

// Update user by userId
func UpdateUserByUserId(c *fiber.Ctx) error {
	db := database.DBConn
	userId := c.Params("userId")

	user := new(models.User)

	if err := db.Where("users.user_id = ?", userId).First(&user).Error; err != nil {
		return c.Status(http.StatusServiceUnavailable).JSON(ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	if err := UpdateUser(c, user); err != nil {
		return c.Status(http.StatusServiceUnavailable).JSON(ResponseHTTP{
			Success: false,
			Message: "Couldn't update the user",
			Data:    nil,
		})
	}

	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Success update user by UserId.",
		Data:    *user,
	})
}

// Update user in db
func UpdateUser(c *fiber.Ctx, u *models.User) error {
	db := database.DBConn

	if err := c.BodyParser(&u); err != nil {
		return c.Status(http.StatusBadRequest).JSON(ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	db.Save(u)

	return nil
}
