package com.multilenguaje;

import okhttp3.*;
import com.google.gson.JsonObject;
import com.google.gson.JsonParser;
import java.net.URLEncoder;
import java.nio.charset.StandardCharsets;

public class v2_traduccion {
    public static String ejecutar(String numero) {
        String textoIngles = v1_soap.ejecutar(numero);
        if (textoIngles.startsWith("Error")) return textoIngles;

        try {
            OkHttpClient client = new OkHttpClient();
            
            String textoCodificado = URLEncoder.encode(textoIngles, StandardCharsets.UTF_8.name());
            String urlTraductor = "https://api.mymemory.translated.net/get?q=" + textoCodificado + "&langpair=en|es";

            Request request = new Request.Builder()
                    .url(urlTraductor)
                    .get()
                    .build();

            try (Response response = client.newCall(request).execute()) {
                if (response.isSuccessful() && response.body() != null) {
                    String jsonResponse = response.body().string();
                    JsonObject jsonObject = JsonParser.parseString(jsonResponse).getAsJsonObject();
                    String textoEspanol = jsonObject.getAsJsonObject("responseData")
                                                    .get("translatedText").getAsString();
                    
                    return textoEspanol.toLowerCase().trim();
                }
            }
        } catch (Exception e) {
            return "Error en la traducción: " + e.getMessage();
        }
        return "Error al traducir el resultado";
    }
}