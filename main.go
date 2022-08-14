package main

import (
	"carbon/internal/controllers"
	"carbon/internal/database"
	"carbon/internal/helpers"
	"carbon/internal/migrations"
	"carbon/internal/repo"
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	ctx := context.Background()
	app := fiber.New(fiber.Config{
		ServerHeader: "Carbon",
		AppName:      "PasteBin",
	})
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", //should be frontend url but for dev purposes this *
	}))
	db := database.NewDatabase(ctx).ConnectDatabase()
	migrations.RunPasteMigration(db)
	api := app.Group("/api")
	v1 := api.Group("v1")
	paste := v1.Group("/paste")
	pasteRepo := repo.NewPasteRepo(ctx, db)
	pasteController := controllers.NewPasteController(ctx, pasteRepo)
	paste.Post("/", pasteController.Create)
	app.Get("/:url", pasteController.View)
	app.Listen(":" + helpers.GoDotEnvVariable("PORT"))
}
