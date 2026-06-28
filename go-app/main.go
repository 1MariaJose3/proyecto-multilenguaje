package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
)

// Definimos las funciones que configurarán las rutas desde los otros archivos
var RegistrarV1Soap func(*fiber.App)
var RegistrarV2Traduccion func(*fiber.App)
var RegistrarV3Nativo func(*fiber.App)

func main() {
	// Inicializar la aplicación de Fiber
	app := fiber.New()

	// Registrar los endpoints de tus tres archivos de versión
	RegistrarV1Soap(app)
	RegistrarV2Traduccion(app)
	RegistrarV3Nativo(app)

	// Levantar el servidor estrictamente en el puerto 8000 e IP local IPv4
	log.Fatal(app.Listen("127.0.0.1:8000"))
}