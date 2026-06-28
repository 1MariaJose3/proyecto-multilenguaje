const express = require('express');
const soap = require('soap');
const https = require('https');
const app = express();
const port = 8000;

const wsdlUrl = 'https://www.dataaccess.com/webservicesserver/NumberConversion.wso?WSDL';

// Ruta exacta: http://localhost:8000/clisoap2?n=10
app.get('/clisoap2', (req, res) => {
    const numero = req.query.n || "0";

    // 1. Consumir el servicio SOAP
    soap.createClient(wsdlUrl, (err, client) => {
        if (err) return res.status(500).send("Error al cargar el WSDL");

        client.NumberToWords({ ubiNum: numero }, (err, result) => {
            if (err) return res.status(500).send("Error en la petición SOAP");

            const resultadoEnIngles = result.NumberToWordsResult.trim();

            // 2. Traducción rápida mediante API HTTP MyMemory sin librerías problemáticas
            const urlTraduccion = `https://api.mymemory.translated.net/get?q=${encodeURIComponent(resultadoEnIngles)}&langpair=en|es`;

            https.get(urlTraduccion, (respuestaNet) => {
                let data = '';
                respuestaNet.on('data', (chunk) => { data += chunk; });
                
                respuestaNet.on('end', () => {
                    const datosJson = JSON.parse(data);
                    const resultadoEnEspanol = datosJson.responseData.translatedText;
                    
                    // 3. Devolver la palabra en español al navegador (ej. "diez")
                    res.send(resultadoEnEspanol);
                });
            }).on('error', (e) => {
                res.status(500).send("Error al traducir: " + e.message);
            });
        });
    });
});

app.listen(port, () => {
    console.log(`Servidor de Traducción corriendo en http://localhost:${port}`);
});