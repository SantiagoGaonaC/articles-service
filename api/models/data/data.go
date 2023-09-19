package data

import (
	"encoding/json"
	"log"
	"products-service/api/models/database"
	models "products-service/api/models/entities"
)

func InsertProductsFromJSON() {
	db, err := database.ConnectToDatabase()
	if err != nil {
		log.Fatal("Error al conectar con la base de datos:", err)
	}

	jsonData := ProductsJSON
	// json a lista de ProductResponse
	var products []models.Product
	if err := json.Unmarshal([]byte(jsonData), &products); err != nil {
		log.Fatal("Error al parsear el JSON:", err)
	}

	for _, product := range products {
		if err := db.Create(&product).Error; err != nil {
			log.Printf("Error al insertar producto en la base de datos: %v\n", err)
		} else {
			log.Printf("Producto insertado en la base de datos: %s\n", product.ProductName)
		}
	}

	log.Println("Proceso de inserción de productos completado.")

	username := "admin"
	password := "Clave123"
	user := models.User{
		Username: username,
		Password: password,
	}

	if err := db.Create(&user).Error; err != nil {
		log.Printf("Error al insertar usuario en la base de datos: %v\n", err)
	} else {
		log.Printf("Usuario insertado en la base de datos: %s\n", user.Username)
	}
	log.Println("Proceso de inserción de usuario completado.")
}
