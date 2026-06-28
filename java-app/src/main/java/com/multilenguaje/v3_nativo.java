package com.multilenguaje;

public class v3_nativo {
    public static String ejecutar(String numeroStr) {
        int num = 0;
        try {
            num = Integer.parseInt(numeroStr != null ? numeroStr : "0");
        } catch (NumberFormatException e) {
            return "Número inválido";
        }

        if (num == 0) return "cero";

        String[] unidades = {"", "uno", "dos", "tres", "cuatro", "cinco", "seis", "siete", "ocho", "nueve"};
        String[] decenas = {"", "diez", "veinte", "treinta", "cuarenta", "cincuenta", "sesenta", "setenta", "ochenta", "noventa"};
        
        if (num < 10) return unidades[num];
        
        if (num == 11) return "once";
        if (num == 12) return "doce";
        if (num == 13) return "trece";
        if (num == 14) return "catorce";
        if (num == 15) return "quince";
        if (num == 16) return "dieciséis";
        if (num == 17) return "diecisiete";
        if (num == 18) return "dieciocho";
        if (num == 19) return "diecinueve";
        
        if (num == 20) return "veinte";
        if (num > 20 && num < 30) return "veinti" + unidades[num % 10];

        if (num < 100) {
            int u = num % 10;
            int d = num / 10;
            if (u == 0) return decenas[d];
            return decenas[d] + " y " + unidades[u];
        }
        
        if (num == 100) return "cien";

        return "Número fuera de rango (0-100)";
    }
}