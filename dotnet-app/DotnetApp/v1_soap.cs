using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Http;
using System.Linq;
using System.Net.Http;
using System.Text;
using System.Xml.Linq;

public static class V1Soap
{
    public static void RegistrarRuta(WebApplication app)
    {
        // Ruta exacta requerida: http://127.0.0.1:8000/clisoap1?n=10
        app.MapGet("/clisoap1", async (HttpContext context) =>
        {
            string numero = context.Request.Query["n"].ToString() ?? "0";

            // Estructura XML limpia para el servicio SOAP de internet
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
            
            var respuesta = await cliente.PostAsync("https://www.dataaccess.com/webservicesserver/NumberConversion.wso", contenido);
            
            if (respuesta.IsSuccessStatusCode)
            {
                string xmlResultado = await respuesta.Content.ReadAsStringAsync();
                var documento = XDocument.Parse(xmlResultado);
                XNamespace ns = "http://www.dataaccess.com/webservicesserver/";
                string textoIngles = documento.Descendants(ns + "NumberToWordsResult").FirstOrDefault()?.Value ?? "Error";

                return Results.Text(textoIngles.Trim(), "text/plain; charset=utf-8");
            }

            return Results.Text("Error al conectar con el servicio SOAP", "text/plain");
        });
    }
}