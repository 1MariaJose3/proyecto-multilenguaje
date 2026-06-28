package main

import (
	"bytes"
	"encoding/xml"
	"io"
	"net/http"
	"strings"
	"github.com/gofiber/fiber/v2"
)

// Estructura para parsear la respuesta XML del servicio SOAP de forma limpia
type SoapResponseV1 struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		NumberToWordsResponse struct {
			NumberToWordsResult string `xml:"NumberToWordsResult"`
		} `xml:"NumberToWordsResponse"`
	} `xml:"Body"`
}

func init() {
	// Enlazamos la ruta con el servidor principal
	RegistrarV1Soap = func(app *fiber.App) {
		// Ruta exacta: http://127.0.0.1:8000/clisoap1?n=10
		app.Get("/clisoap1", func(c *fiber.Ctx) error {
			numero := c.Query("n", "0")

			// 1. Construimos el sobre XML SOAP para el servidor externo
			sobreSoap := `<?xml version="1.0" encoding="utf-8"?>
			<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
			  <soap:Body>
				<NumberToWords xmlns="http://www.dataaccess.com/webservicesserver/">
				  <ubiNum>` + numero + `</ubiNum>
				</NumberToWords>
			  </soap:Body>
			</soap:Envelope>`

			// 2. Enviamos la petición POST a internet
			resp, err := http.Post("https://www.dataaccess.com/webservicesserver/NumberConversion.wso", "text/xml; charset=utf-8", bytes.NewBufferString(sobreSoap))
			if err != nil {
				return c.Status(500).SendString("Error al conectar con el servicio SOAP")
			}
			defer resp.Body.Close()

			// 3. Leemos y parseamos el XML devuelto
			bodyBytes, _ := io.ReadAll(resp.Body)
			var resultado SoapResponseV1
			if err := xml.Unmarshal(bodyBytes, &resultado); err != nil {
				return c.Status(500).SendString("Error al procesar el XML")
			}

			textoIngles := strings.TrimSpace(resultado.Body.NumberToWordsResponse.NumberToWordsResult)
			c.Set("Content-Type", "text/plain; charset=utf-8")
			return c.SendString(textoIngles)
		})
	}
}