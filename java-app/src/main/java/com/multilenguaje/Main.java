package com.multilenguaje;

import io.javalin.Javalin;

public class Main {
    public static void main(String[] args) {
        // Inicializar servidor en el puerto 8000 e IP local IPv4
        Javalin app = Javalin.create().start("127.0.0.1", 8000);

        // Mapear las rutas hacia tus tres archivos de versión independientes
        app.get("/clisoap1", ctx -> ctx.result(v1_soap.ejecutar(ctx.queryParam("n"))));
        app.get("/clisoap2", ctx -> ctx.result(v2_traduccion.ejecutar(ctx.queryParam("n"))));
        app.get("/conintl", ctx -> ctx.result(v3_nativo.ejecutar(ctx.queryParam("n"))));
    }
}