const express = require('express');
const soap = require('soap');
const app = express();

// Configuramos el servidor para que escuche en el puerto 8000
const PORT = 8000;

// Ruta: http://localhost:8000/clisoap1?n=10
app.get('/clisoap1', (req, res) => {
    // 1. URL del archivo WSDL público 
    const wsdlUrl = 'https://www.dataaccess.com/webservicesserver/NumberConversion.wso?WSDL';
    
    // 2. Capturamos el parámetro 'n' proveniente de la URL de la petición
    const numero = req.query.n || '0';

    // 3. Creamos el cliente SOAP de Node de forma asíncrona
    soap.createClient(wsdlUrl, (err, client) => {
        if (err) {
            return res.status(500).send("Error al conectar con el servicio SOAP");
        }
        
        // 4. Invocamos el método NumberToWords enviando el parámetro estructurado
        client.NumberToWords({ ubiNum: numero }, (errSoap, result) => {
            if (errSoap) {
                return res.status(500).send("Error en la ejecución del servicio web");
            }
            
            // 5. Enviamos la respuesta obtenida en inglés al navegador (ej: "ten")
            res.send(result.NumberToWordsResult);
        });
    });
});

// Inicializamos el servidor web
app.listen(PORT, () => {
    console.log(`Servidor de Node corriendo en http://localhost:${PORT}`);
});
