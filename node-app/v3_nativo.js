const express = require('express');
const app = express();
const PORT = 8000;

// Función nativa en JavaScript para convertir números a palabras en español
function numeroALetras(num) {
    if (num === 0) return "cero";

    const unidades = ["", "uno", "dos", "tres", "cuatro", "cinco", "seis", "siete", "ocho", "nueve"];
    const decenas = ["", "diez", "veinte", "treinta", "cuarenta", "cincuenta", "sesenta", "setenta", "ochenta", "noventa"];
    const especiales = ["diez", "once", "doce", "trece", "catorce", "quince", "dieciséis", "diecisiete", "dieciocho", "diecinueve"];
    const decenasMod = ["", "veinti", "treinta y ", "cuarenta y ", "cincuenta y ", "sesenta y ", "setenta y ", "ochenta y ", "noventa y "];

    if (num < 10) return unidades[num];
    if (num >= 10 && num < 20) return especiales[num - 10];
    
    if (num >= 20 && num < 100) {
        const d = Math.floor(num / 10);
        const u = num % 10;
        if (u === 0) {
            return d === 2 ? "veinte" : decenas[d];
        }
        return decenasMod[d - 1] + unidades[u];
    }
    
    if (num === 100) return "cien";
    
    // Soporte básico para números mayores por si pruebas con valores más altos
    return "número fuera de rango de prueba básica (0-100)";
}

// Ruta requerida por tu caso de estudio: http://localhost:8000/conintl?n=10
app.get('/conintl', (req, res) => {
    // Captura el parámetro 'n', si no existe usa "0" y lo fuerza a número entero
    const numero = parseInt(req.query.n || "0", 10);

    if (isNaN(numero)) {
        return res.status(400).send("Por favor, introduce un número válido.");
    }

    // Ejecuta la función local
    const resultadoPalabras = numeroALetras(numero);
    
    // Envía el resultado al navegador
    res.send(resultadoPalabras);
});

app.listen(PORT, () => {
    console.log(`Servidor nativo local corriendo en http://localhost:${PORT}`);
});