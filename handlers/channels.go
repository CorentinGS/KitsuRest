package handlers

import (
	"fmt"
	"kitsurest/database"
	"kitsurest/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// GET

// Get all channels
func GetAllChannels(c *fiber.Ctx) error {
	db := database.DBConn

	var channels []models.Channel

	if res := db.Find(&channels); res.Error != nil {

		return c.JSON(ResponseHTTP{
			Success: false,
			Message: "Get All channels",
			Data:    nil,
		})
	}
	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Get All channels",
		Data:    channels,
	})
}

// Get channel by channelId
func GetChannelByChannelId(c *fiber.Ctx) error {
	db := database.DBConn
	channelId := c.Params("channelId")

	channel := new(models.Channel)

	if err := db.First(&channel, channelId).Error; err != nil {
		switch err.Error() {
		case "record not found":
			return c.Status(http.StatusNotFound).JSON(ResponseHTTP{
				Success: false,
				Message: fmt.Sprintf("Channel with ChannelID %v not found.", channelId),
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
		Message: "Success get user by UserId.",
		Data:    *channel,
	})
}

// Get channel by id
func GetChannelById(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	channel := new(models.Channel)

	if err := db.First(&channel, id).Error; err != nil {
		switch err.Error() {
		case "record not found":
			return c.Status(http.StatusNotFound).JSON(ResponseHTTP{
				Success: false,
				Message: fmt.Sprintf("Channel with ID %v not found.", id),
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
		Message: fmt.Sprintf("Get channel by ID %v", id),
		Data:    *channel,
	})

}

// POST

// Create a new channel
func CreateChannel(c *fiber.Ctx) error {
	db := database.DBConn

	channel := new(models.Channel)

	if err := c.BodyParser(&channel); err != nil {
		return c.Status(http.StatusBadRequest).JSON(ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	db.Create(channel)

	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Success register a channel",
		Data:    *channel,
	})
}

// PUT

// Update channel by id
func UpdateChannelById(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")

	channel := new(models.Channel)

	if err := db.First(&channel, id).Error; err != nil {
		return c.Status(http.StatusServiceUnavailable).JSON(ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	if err := c.BodyParser(&channel); err != nil {
		return c.Status(http.StatusBadRequest).JSON(ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	db.Save(channel)

	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Success update channel by Id.",
		Data:    *channel,
	})
}
