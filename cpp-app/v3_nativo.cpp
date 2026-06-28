#include "crow.h"
#include <string>

// Función local para simular la conversión nativa a palabras
std::string ConvertirNumeroALetras(std::string n) {
    if (n == "10") {
        return "diez";
    }
    return "conversion nativa local";
}

int main() {
    crow::SimpleApp app;

    // Ruta: http://localhost:8000/conintl?n=10
    CROW_ROUTE(app, "/conintl")
    ([](const crow::request& req) {
        auto n_param = req.url_params.get("n");
        std::string n = n_param ? n_param : "0";

        // 1. En un entorno real se usaría ICU RuleBasedNumberFormat configurando el locale "es":
        // UErrorCode status = U_ZERO_ERROR;
        // RuleBasedNumberFormat* formatter = new RuleBasedNumberFormat(URBNF_SPELLOUT, Locale("es"), status);
        
        // 2. Procesamos el parámetro de forma nativa en el servidor
        std::string resultado_letras = ConvertirNumeroALetras(n);

        // 3. Devolvemos el resultado al navegador
        return crow::response(resultado_letras);
    });

    app.port(8000).multithreaded().run();
}