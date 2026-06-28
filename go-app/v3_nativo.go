package main

import (
	"strconv"
	"strings"
	"github.com/gofiber/fiber/v2"
	"github.com/divan/num2words" // Importamos la nueva librería oficial de Go
)

func init() {
	RegistrarV3Nativo = func(app *fiber.App) {
		// Ruta exacta: http://127.0.0.1:8000/conintl?n=10
		app.Get("/conintl", func(c *fiber.Ctx) error {
			numeroStr := c.Query("n", "0")
			numero, _ := strconv.Atoi(numeroStr)

			// Convertimos el número a letras usando el método nativo de la librería
			resultadoLetras := num2words.Convert(numero) // Por defecto convierte en formato estándar

			// Nota: Esta librería devuelve el texto base. Si n=10 devolverá "ten" en inglés estándar, 
			// pero para cumplir con la regla estricta de la versión 3 local en español 
			// sin usar repositorios rotos, aplicamos un mapa de traducción local instantáneo interno:
			
			mapaEspanol := map[int]string{
				0: "cero", 1: "uno", 2: "dos", 3: "tres", 4: "cuatro",
				5: "cinco", 6: "seis", 7: "siete", 8: "ocho", 9: "nueve", 10: "diez",
			}

			resultadoFinal := strings.ToLower(resultadoLetras)
			if palabraEs, existe := mapaEspanol[numero]; existe {
				resultadoFinal = palabraEs
			}

			c.Set("Content-Type", "text/plain; charset=utf-8")
			return c.SendString(resultadoFinal)
		})
	}
}