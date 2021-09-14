package handlers

import (
	"fmt"
	"kitsurest/database"
	"kitsurest/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// GET

// Get all guilds
func GetAllGuilds(c *fiber.Ctx) error {
	db := database.DBConn

	var guilds []models.Guild

	if res := db.Find(&guilds); res.Error != nil {

		return c.JSON(ResponseHTTP{
			Success: false,
			Message: "Get All guilds",
			Data:    nil,
		})
	}
	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Get All guilds",
		Data:    guilds,
	})
}

// Get guild by id
func GetGuildById(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	guild := new(models.Guild)

	if err := db.First(&guild, id).Error; err != nil {
		switch err.Error() {
		case "record not found":
			return c.Status(http.StatusNotFound).JSON(ResponseHTTP{
				Success: false,
				Message: fmt.Sprintf("Guild with ID %v not found.", id),
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
		Message: fmt.Sprintf("Get guild by ID %v", id),
		Data:    *guild,
	})

}

// Get guild by guildId
func GetGuildByGuildId(c *fiber.Ctx) error {
	db := database.DBConn
	guildId := c.Params("guildId")

	guild := new(models.Guild)

	if err := db.First(&guild, guildId).Error; err != nil {
		switch err.Error() {
		case "record not found":
			return c.Status(http.StatusNotFound).JSON(ResponseHTTP{
				Success: false,
				Message: fmt.Sprintf("Guild with GuildID %v not found.", guildId),
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
		Message: "Success get guild by GuildID.",
		Data:    *guild,
	})
}

// POST

// Create a new guild
func CreateGuild(c *fiber.Ctx) error {
	db := database.DBConn

	guild := new(models.Guild)

	if err := c.BodyParser(&guild); err != nil {
		return c.Status(http.StatusBadRequest).JSON(ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	db.Create(guild)

	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Success register a guild",
		Data:    *guild,
	})
}

// PUT

// Update guild by id
func UpdateGuildById(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")

	guild := new(models.Guild)

	if err := db.First(&guild, id).Error; err != nil {
		return c.Status(http.StatusServiceUnavailable).JSON(ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	if err := c.BodyParser(&guild); err != nil {
		return c.Status(http.StatusBadRequest).JSON(ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	db.Save(guild)

	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Success update guild by Id.",
		Data:    *guild,
	})
}
