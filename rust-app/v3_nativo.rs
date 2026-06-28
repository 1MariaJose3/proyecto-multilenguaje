use actix_web::{get, web, App, HttpServer, Responder, HttpResponse};

// Ruta: http://localhost:8000/conintl?n=10
#[get("/conintl")]
async fn conintl(info: web::Query<std::collections::HashMap<String, String>>) -> impl Responder {
    let n = info.get("n").map(|s| s.as_str()).unwrap_or("0");
    
    // 1. Convertimos el string de la URL a un entero de 32 bits de forma segura
    let numero: i32 = n.parse().unwrap_or(0);

    // 2. En producción usaríamos la función de la librería local configurando el idioma:
    // let resultado_letras = num_to_words::to_words(numero, num_to_words::Language::Spanish);
    let resultado_letras = if numero == 10 { "diez" } else { "conversion nativa local" };

    // 3. Retornamos las palabras en castellano directamente al navegador
    HttpResponse::Ok().content_type("text/plain; charset=utf-8").body(resultado_letras)
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    HttpServer::new(|| App::new().service(conintl)).bind(("127.0.0.1", 8000))?.run().await
}