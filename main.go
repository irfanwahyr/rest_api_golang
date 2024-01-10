package main

import (
	"log"
	"os"
	"react_go_catalog_web/database"
	"react_go_catalog_web/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	database.ConnectDB()
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	port := os.Getenv("PORT")
	app := fiber.New()
	log.Println("CONNECT TO PORT:", port)
	routes.Setup(app)
	app.Listen(":"+port)
}