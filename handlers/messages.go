package handlers

import (
	"fmt"
	"kitsurest/database"
	"kitsurest/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// Response Struct
type ResponseHTTP struct {
	Success bool        `json:"sucess"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// GET

// Get last message
func GetLastMessage(c *fiber.Ctx) error {

	db := database.DBConn

	var message models.Message
	var author models.User

	if res := db.Last(&message); res.Error != nil {
		return c.JSON(ResponseHTTP{
			Success: false,
			Message: "Get last message",
			Data:    nil,
		})
	}
	db.Where("id = ?", message.UserID).First(&author)

	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Get last message",
		Data: fiber.Map{
			"Messages": message,
			"Author":   author,
		},
	})

}

// Get all messages
func GetAllMessages(c *fiber.Ctx) error {

	db := database.DBConn

	var messages []models.Message

	if res := db.Find(&messages); res.Error != nil {

		return c.JSON(ResponseHTTP{
			Success: false,
			Message: "Get All messages",
			Data:    nil,
		})
	}
	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Get All messages",
		Data:    messages,
	})
}

// Get messages from an user
func GetMessagesByUser(c *fiber.Ctx) error {
	userId := c.Params("userId")
	db := database.DBConn

	var messages []models.Message

	if err := db.Joins("Join users ON users.id = messages.user_id AND users.user_id = ?", userId).Find(&messages).Error; err != nil {
		switch err.Error() {
		case "record not found":
			return c.Status(http.StatusNotFound).JSON(ResponseHTTP{
				Success: false,
				Message: fmt.Sprintf("Message with UserId %v not found.", userId),
				Data:    nil,
			})
		default:
			return c.Status(http.StatusServiceUnavailable).JSON(ResponseHTTP{
				Success: false,
				Message: err.Error(),
				Data:    nil,
			})

		}
	}

	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Success get message by UserID.",
		Data:    messages,
	})
}

// Get message by id
func GetMessageById(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	message := new(models.Message)
	if err := db.First(&message, id).Error; err != nil {
		switch err.Error() {
		case "record not found":
			return c.Status(http.StatusNotFound).JSON(ResponseHTTP{
				Success: false,
				Message: fmt.Sprintf("Message with ID %v not found.", id),
				Data:    nil,
			})
		default:
			return c.Status(http.StatusServiceUnavailable).JSON(ResponseHTTP{
				Success: false,
				Message: err.Error(),
				Data:    nil,
			})

		}
	}

	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Success get message by ID.",
		Data:    *message,
	})

}

// POST

// Create a new message
func CreateNewMessage(c *fiber.Ctx) error {
	db := database.DBConn

	message := new(models.Message)

	if err := c.BodyParser(&message); err != nil {
		return c.Status(http.StatusBadRequest).JSON(ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	db.Create(message)

	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Success register a message",
		Data:    *message,
	})

}
