package main

import (
	"log"
	data "products-service/api/models/data"
	"products-service/api/models/database"
	models "products-service/api/models/entities"
	"products-service/api/routes"
)

func main() {
	db, err := database.ConnectToDatabase()
	if err != nil {
		log.Fatal("Error al conectar con la base de datos:", err)
	}

	//crear las tablas si aún no existen:
	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Favorite{})
	log.Println("Tablas creadas")
	data.InsertProductsFromJSON()

	r := routes.SetupRouter(db)
	log.Fatal(r.Run(":8080"))
}
