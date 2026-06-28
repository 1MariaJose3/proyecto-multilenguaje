using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Http;
using System;
using System.Linq;
using System.Net.Http;
using System.Text;
using System.Text.Json;
using System.Xml.Linq;

public static class V2Traduccion
{
    public static void RegistrarRuta(WebApplication app)
    {
        // Ruta exacta requerida: http://127.0.0.1:8000/clisoap2?n=10
        app.MapGet("/clisoap2", async (HttpContext context) =>
        {
            string numero = context.Request.Query["n"].ToString() ?? "0";

            // 1. Consumir el servicio SOAP (Misma lógica real de la V1)
            string sobreSoap = $@"<?xml version=""1.0"" encoding=""utf-8""?>
            <soap:Envelope xmlns:soap=""http://schemas.xmlsoap.org/soap/envelope/"">
              <soap:Body>
                <NumberToWords xmlns=""http://www.dataaccess.com/webservicesserver/"">
                  <ubiNum>{numero}</ubiNum>
                </NumberToWords>
              </soap:Body>
            </soap:Envelope>";

            using var cliente = new HttpClient();
            var contenido = new StringContent(sobreSoap, Encoding.UTF8, "text/xml");
            var respuestaSoap = await cliente.PostAsync("https://www.dataaccess.com/webservicesserver/NumberConversion.wso", contenido);
            
            if (!respuestaSoap.IsSuccessStatusCode)
                return Results.Text("Error en SOAP", "text/plain");

            string xmlResultado = await respuestaSoap.Content.ReadAsStringAsync();
            var documento = XDocument.Parse(xmlResultado);
            XNamespace ns = "http://www.dataaccess.com/webservicesserver/";
            string textoIngles = documento.Descendants(ns + "NumberToWordsResult").FirstOrDefault()?.Value?.Trim() ?? "";

            // 2. Consumir la API real de MyMemory para traducir el resultado al vuelo
            string urlTraductor = $"https://api.mymemory.translated.net/get?q={Uri.EscapeDataString(textoIngles)}&langpair=en|es";
            var respuestaTraduccion = await cliente.GetAsync(urlTraductor);
            
            if (respuestaTraduccion.IsSuccessStatusCode)
            {
                string jsonResultado = await respuestaTraduccion.Content.ReadAsStringAsync();
                using var jsonDoc = JsonDocument.Parse(jsonResultado);
                string textoEspanol = jsonDoc.RootElement.GetProperty("responseData").GetProperty("translatedText").GetString() ?? "Error";

                return Results.Text(textoEspanol.ToLower(), "text/plain; charset=utf-8");
            }

            return Results.Text("Error en la traducción", "text/plain");
        });
    }
}