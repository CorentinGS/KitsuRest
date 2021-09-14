package routes

import (
	"kitsurest/handlers"

	"github.com/gofiber/fiber/v2"
)

func New() *fiber.App {
	app := fiber.New()

	api := app.Group("/api")

	v1 := api.Group("/v1", func(c *fiber.Ctx) error {
		return c.Next()
	})

	api.Get("/", func(c *fiber.Ctx) error {
		return fiber.NewError(782, "This is not a valid route")
	})

	// Messages
	v1.Get("/messages", handlers.GetAllMessages)
	v1.Get("/messages/last", handlers.GetLastMessage)
	v1.Get("/messages/:id", handlers.GetMessageById)
	v1.Get("/messages/user/:userId", handlers.GetMessagesByUser)
	v1.Post("/messages/new", handlers.CreateNewMessage)

	// Users
	v1.Get("/users/userid/:userId", handlers.GetUserByUserId)
	v1.Get("/users/id/:id", handlers.GetUserById)
	v1.Post("/users/new", handlers.CreateNewUser)
	v1.Put("/users/userid/:userId", handlers.UpdateUserByUserId)
	v1.Put("/users/id/:id", handlers.UpdateUserById)

	// Channels
	v1.Get("/channels", handlers.GetAllChannels)
	v1.Get("/channels/id/:id", handlers.GetChannelById)
	v1.Get("/channels/channelid/:channelId", handlers.GetChannelByChannelId)
	v1.Post("/channels/new", handlers.CreateChannel)
	v1.Put("/channels/id/:id", handlers.UpdateChannelById)

	// Guilds
	v1.Get("/guilds", handlers.GetAllGuilds)
	v1.Get("/guilds/id/:id", handlers.GetGuildById)
	v1.Get("/guilds/guildid/:guildId", handlers.GetGuildByGuildId)
	v1.Post("/guilds/new", handlers.CreateGuild)
	v1.Put("/guilds/id/:id", handlers.UpdateGuildById)

	return app

}
