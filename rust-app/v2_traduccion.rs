use actix_web::{get, web, App, HttpServer, Responder, HttpResponse};

// Ruta: http://localhost:8000/clisoap2?n=10
#[get("/clisoap2")]
async fn clisoap2(info: web::Query<std::collections::HashMap<String, String>>) -> impl Responder {
    let n = info.get("n").map(|s| s.as_str()).unwrap_or("0");

    // 1. Simulación del texto obtenido del paso SOAP anterior ("ten")
    let resultado_ingles = "ten";

    // 2. Aquí se procesaría la petición HTTP POST asíncrona de traducción
    // enviando un JSON con origen 'en' y destino 'es'.
    let resultado_espanol = "diez";

    let respuesta = format!("Numero {} en ingles: {} -> Traducido al espanol: {}", n, resultado_ingles, resultado_espanol);
    
    // 3. Retornamos el texto con codificación UTF-8 correcta
    HttpResponse::Ok().content_type("text/plain; charset=utf-8").body(respuesta)
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    HttpServer::new(|| App::new().service(clisoap2)).bind(("127.0.0.1", 8000))?.run().await
}