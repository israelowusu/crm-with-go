package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/israelowusu/crm-with-go.git/database"
	"github.com/israelowusu/crm-with-go.git/lead"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads())
	app.Get("/api/v1/lead/:id", lead.GetLead())
	app.Post("/api/v1/lead", lead.NewLead())
	app.Delete("/api/v1/lead/:id", lead.DeleteLead())
}

func initDatabase() {
	database.DBConn, err = gorm.Open(sqlite.Open("leads.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	fmt.Println("Connection to database is successful!")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database migrated!")
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(":8080")
	defer func() {
		sqlDB, err := database.DBConn.DB()
		if err != nil {
			panic("Failed to get sql.DB from gorm.DB!")
		}
		sqlDB.Close()
	}()

}
