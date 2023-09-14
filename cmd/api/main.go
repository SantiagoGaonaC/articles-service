package main

import (
	"articles-service/api/models/database"
	models "articles-service/api/models/entities"
	"articles-service/api/routes"
	"log"
)

func main() {
	db, err := database.ConnectToDatabase()
	if err != nil {
		log.Fatal("Error al conectar con la base de datos:", err)
	}

	//crear las tablas si aún no existen:
	db.AutoMigrate(&models.User{}, &models.Article{}, &models.Favorite{})
	log.Println("Tablas creadas")

	r := routes.SetupRouter()
	log.Fatal(r.Run(":8080"))
}
