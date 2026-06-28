use actix_web::{get, web, App, HttpServer, Responder, HttpResponse};

// Ruta: http://localhost:8000/clisoap1?n=10
#[get("/clisoap1")]
async fn clisoap1(info: web::Query<std::collections::HashMap<String, String>>) -> impl Responder {
    // 1. Capturamos el parámetro 'n' de la URL de forma segura
    let n = info.get("n").map(|s| s.as_str()).unwrap_or("0");

    // 2. En Rust, mapearíamos el sobre XML de SOAP usando la librería `yaserde`
    // y lo enviaríamos mediante un cliente HTTP asíncrono como `reqwest`.
    let resultado_soap = format!("Resultado SOAP para el numero {}: ten", n);

    // 3. Devolvemos la respuesta en texto plano al navegador
    HttpResponse::Ok().body(resultado_soap)
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    println!("Servidor de Rust corriendo en http://localhost:8000");
    // 4. Levantamos el servidor escuchando en el puerto 8000
    HttpServer::new(|| {
        App::new().service(clisoap1)
    })
    .bind(("127.0.0.1", 8000))?
    .run()
    .await
}
