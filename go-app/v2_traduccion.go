package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io"
	"net/http"
	"net/url"
	"strings"
	"github.com/gofiber/fiber/v2"
)

type SoapResponseV2 struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		NumberToWordsResponse struct {
			NumberToWordsResult string `xml:"NumberToWordsResult"`
		} `xml:"NumberToWordsResponse"`
	} `xml:"Body"`
}

// Estructura para decodificar la respuesta JSON del traductor MyMemory
type MyMemoryResponse struct {
	ResponseData struct {
		TranslatedText string `json:"translatedText"`
	} `json:"responseData"`
}

func init() {
	RegistrarV2Traduccion = func(app *fiber.App) {
		// Ruta exacta: http://127.0.0.1:8000/clisoap2?n=10
		app.Get("/clisoap2", func(c *fiber.Ctx) error {
			numero := c.Query("n", "0")

			// 1. Consumir el SOAP (Misma lógica real de la V1)
			sobreSoap := `<?xml version="1.0" encoding="utf-8"?>
			<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
			  <soap:Body>
				<NumberToWords xmlns="http://www.dataaccess.com/webservicesserver/">
				  <ubiNum>` + numero + `</ubiNum>
				</NumberToWords>
			  </soap:Body>
			</soap:Envelope>`

			respSoap, err := http.Post("https://www.dataaccess.com/webservicesserver/NumberConversion.wso", "text/xml; charset=utf-8", bytes.NewBufferString(sobreSoap))
			if err != nil {
				return c.Status(500).SendString("Error en SOAP")
			}
			defer respSoap.Body.Close()

			bodySoap, _ := io.ReadAll(respSoap.Body)
			var resultadoSoap SoapResponseV2
			xml.Unmarshal(bodySoap, &resultadoSoap)
			textoIngles := strings.TrimSpace(resultadoSoap.Body.NumberToWordsResponse.NumberToWordsResult)

			// 2. Consumir la API real de MyMemory para traducir de inglés a español
			urlTraductor := "https://api.mymemory.translated.net/get?q=" + url.QueryEscape(textoIngles) + "&langpair=en|es"
			respTrad, err := http.Get(urlTraductor)
			if err != nil {
				return c.Status(500).SendString("Error en la traducción")
			}
			defer respTrad.Body.Close()

			bodyTrad, _ := io.ReadAll(respTrad.Body)
			var resultadoTrad MyMemoryResponse
			json.Unmarshal(bodyTrad, &resultadoTrad)

			textoEspanol := strings.ToLower(resultadoTrad.ResponseData.TranslatedText)
			c.Set("Content-Type", "text/plain; charset=utf-8")
			return c.SendString(textoEspanol)
		})
	}
}