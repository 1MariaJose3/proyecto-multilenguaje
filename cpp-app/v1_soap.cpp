#include "crow.h" // Librería de micro-framework web para C++
#include <string>

int main() {
    crow::SimpleApp app;

    // Ruta: http://localhost:8000/clisoap1?n=10
    CROW_ROUTE(app, "/clisoap1")
    ([](const crow::request& req) {
        // 1. Capturamos el parámetro 'n' desde la URL
        auto n_param = req.url_params.get("n");
        std::string n = n_param ? n_param : "0";

        // 2. En producción, aquí se invocarían las funciones generadas por gSOAP:
        // NumberConversionSoapProxy service;
        // _ns1__NumberToWords request;
        // _ns1__NumberToWordsResponse response;
        // request.ubiNum = std::stoull(n);
        // service.NumberToWords(&request, response);
        
        // 3. Simulamos la respuesta exitosa del servicio SOAP remoto
        std::string resultado_ingles = "Resultado SOAP para el numero " + n + ": ten";
        
        return crow::response(resultado_ingles);
    });

    // 4. Configuramos el servidor para que escuche en el puerto 8000
    app.port(8000).multithreaded().run();
}