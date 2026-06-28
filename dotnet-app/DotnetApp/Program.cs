using Microsoft.AspNetCore.Builder;
using Microsoft.Extensions.Hosting;

var builder = WebApplication.CreateBuilder(args);

// Configurar servidor para escuchar en el puerto 8000 e IP local IPv4
builder.WebHost.UseUrls("http://127.0.0.1:8000");

var app = builder.Build();

// Registrar las rutas vinculadas a tus nuevos archivos
V1Soap.RegistrarRuta(app);
V2Traduccion.RegistrarRuta(app);
V3Nativo.RegistrarRuta(app);

app.Run();