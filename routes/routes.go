package routes

import (
	"kitsurest/handlers"

	"github.com/gofiber/fiber/v2"
)

// Create new app
func New() *fiber.App {
	app := fiber.New()

	// Api group
	api := app.Group("/api")

	// v1 group "/api/v1"
	v1 := api.Group("/v1", func(c *fiber.Ctx) error {
		return c.Next()
	})

	api.Get("/", func(c *fiber.Ctx) error {
		return fiber.NewError(fiber.StatusForbidden, "This is not a valid route") // Custom error
	})

	// Messages
	v1.Get("/message", handlers.GetAllMessages)                 // Return all messages. Not recommended
	v1.Get("/message/last", handlers.GetLastMessage)            // Return last message
	v1.Get("/message/:id", handlers.GetMessageById)             // Return message by id
	v1.Get("/message/user/:userId", handlers.GetMessagesByUser) // Return every messages from a specific user
	v1.Post("/message/new", handlers.CreateNewMessage)          // Create a new message
	//TODO: Add new routes
	// - Put: Update Message
	// - Get: Filter by guild/channel/user */

	// Users
	v1.Get("/user/userid/:userId", handlers.GetUserByUserId)    // Get user by user_id
	v1.Get("/user/id/:id", handlers.GetUserById)                // Get user by id
	v1.Post("/user/new", handlers.CreateNewUser)                // Create a new user
	v1.Put("/user/userid/:userId", handlers.UpdateUserByUserId) // Update an user using his user_id
	v1.Put("/user/id/:id", handlers.UpdateUserById)             // Update an user using his id

	// Channels
	v1.Get("/channel", handlers.GetAllChannels)                             // Return all channels. Not recommended
	v1.Get("/channel/id/:id", handlers.GetChannelById)                      // Get channel by id
	v1.Get("/channel/channelid/:channelId", handlers.GetChannelByChannelId) // Get channel by channel_id
	v1.Post("/channel/new", handlers.CreateChannel)                         // Create a new channel
	v1.Put("/channel/id/:id", handlers.UpdateChannelById)                   // Update channel by id

	// Guilds
	v1.Get("/guild", handlers.GetAllGuilds)                       // Return all guilds. Not recommended
	v1.Get("/guild/id/:id", handlers.GetGuildById)                // Get guild by id
	v1.Get("/guild/guildid/:guildId", handlers.GetGuildByGuildId) // Get guild by guild_id
	v1.Post("/guild/new", handlers.CreateGuild)                   // Create a new guild
	v1.Put("/guild/id/:id", handlers.UpdateGuildById)             // Update guild by id

	return app

}
