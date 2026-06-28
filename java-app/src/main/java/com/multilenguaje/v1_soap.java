package com.multilenguaje;

import okhttp3.*;

public class v1_soap {
    public static String ejecutar(String numero) {
        if (numero == null || numero.trim().isEmpty()) numero = "0";
        
        try {
            OkHttpClient client = new OkHttpClient();

            // Enviamos el sobre plano y directo
            String soapEnvelope = "<?xml version=\"1.0\" encoding=\"utf-8\"?>\n" +
                    "<soap:Envelope xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\">\n" +
                    "  <soap:Body>\n" +
                    "    <NumberToWords xmlns=\"http://www.dataaccess.com/webservicesserver/\">\n" +
                    "      <ubiNum>" + numero + "</ubiNum>\n" +
                    "    </NumberToWords>\n" +
                    "  </soap:Body>\n" +
                    "</soap:Envelope>";

            RequestBody body = RequestBody.create(soapEnvelope, MediaType.parse("text/xml; charset=utf-8"));
            
            Request request = new Request.Builder()
                    .url("https://www.dataaccess.com/webservicesserver/NumberConversion.wso")
                    .post(body)
                    .build();

            try (Response response = client.newCall(request).execute()) {
                if (response.body() != null) {
                    // RETORNAMOS TODO EL XML COMPLETO para inspeccionarlo en el navegador
                    return response.body().string();
                }
                return "Error: Cuerpo de respuesta vacío";
            }
        } catch (Exception e) {
            return "Error en la petición SOAP: " + e.getMessage();
        }
    }
}