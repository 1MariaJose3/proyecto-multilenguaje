#include "crow.h"
#include <string>

int main() {
    crow::SimpleApp app;

    // Ruta: http://localhost:8000/clisoap2?n=10
    CROW_ROUTE(app, "/clisoap2")
    ([](const crow::request& req) {
        auto n_param = req.url_params.get("n");
        std::string n = n_param ? n_param : "0";

        // 1. Resultado en inglés obtenido de la consulta SOAP simulada ("ten")
        std::string resultado_ingles = "ten";

        // 2. Aquí se realizaría la petición HTTP a la API de traducción:
        // cpr::Response r = cpr::Post(cpr::Url{"https://libretranslate.de/translate"}, ...);
        std::string resultado_espanol = "diez"; 

        std::string respuesta = "Numero " + n + " en ingles: " + resultado_ingles + " -> Traducido: " + resultado_espanol;
        
        return crow::response(respuesta);
    });

    app.port(8000).multithreaded().run();
}