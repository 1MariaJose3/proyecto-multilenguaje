require 'sinatra'
require 'numbers_and_words'

# Configuramos el servidor en el puerto 8000
set :port, 8000

# Ruta: http://localhost:8000/conintl?n=10
get '/conintl' do
  # 1. Configuramos el idioma de la librería de conversión a español (:es)
  I18n.locale = :es
  
  # 2. Convertimos el parámetro 'n' de la URL a un número entero (.to_i)
  numero = params[:n].to_i
  
  # 3. Usamos el método de la librería para transformar el número a palabras directamente en el servidor
  resultado = numero.to_words
  
  # 4. Devolvemos las letras al navegador (por ejemplo: "diez")
  resultado
end