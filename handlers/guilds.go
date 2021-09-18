package handlers

import (
	"fmt"
	"kitsurest/database"
	"kitsurest/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// GET

// GetAllGuilds
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

// GetGuildByID
func GetGuildByID(c *fiber.Ctx) error {
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

// GetGuildByGuildID
func GetGuildByGuildID(c *fiber.Ctx) error {
	db := database.DBConn
	guildID := c.Params("guildID")

	guild := new(models.Guild)

	if err := db.Where("guilds.guild_id = ?", guildID).First(&guild).Error; err != nil {
		switch err.Error() {
		case "record not found":
			return c.Status(http.StatusNotFound).JSON(ResponseHTTP{
				Success: false,
				Message: fmt.Sprintf("Guild with GuildID %v not found.", guildID),
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

// UpdateGuildByID
func UpdateGuildByID(c *fiber.Ctx) error {
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
