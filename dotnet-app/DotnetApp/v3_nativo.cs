using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Http;
using Humanizer; // La librería oficial de NuGet que añadimos al .csproj
using System.Globalization;

public static class V3Nativo
{
    public static void RegistrarRuta(WebApplication app)
    {
        // Ruta exacta requerida: http://127.0.0.1:8000/conintl?n=10
        app.MapGet("/conintl", (HttpContext context) =>
        {
            int numero = int.TryParse(context.Request.Query["n"], out var res) ? res : 0;

            // Usamos el método nativo .ToWords en cultura española de la librería NuGet
            string resultadoLetras = numero.ToWords(new CultureInfo("es"));

            return Results.Text(resultadoLetras.ToLower(), "text/plain; charset=utf-8");
        });
    }
}