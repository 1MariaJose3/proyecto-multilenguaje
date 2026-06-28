require 'sinatra'
require 'savon'
require 'net/http'
require 'json'
require 'uri'

# Configuramos el servidor en el puerto 8000
set :port, 8000

# Ruta exacta pedida: http://localhost:8000/clisoap2?n=10
get '/clisoap2' do
  # 1. Conectamos con el servicio web SOAP público usando Savon
  client = Savon.client(wsdl: "https://www.dataaccess.com/webservicesserver/NumberConversion.wso?WSDL")
  
  # 2. Enviamos el número 'n' capturado desde la URL al servicio SOAP
  response = client.call(:number_to_words, message: { "ubiNum" => params[:n] })
  
  # 3. Extraemos el resultado original en inglés (por ejemplo: "ten")
  resultado_en_ingles = response.body[:number_to_words_response][:number_to_words_result].to_s.strip
  
  # 4. Traducción inmediata usando la API HTTP pública MyMemory
  uri = URI.parse("https://api.mymemory.translated.net/get?q=#{URI.encode_www_form_component(resultado_en_ingles)}&langpair=en|es")
  respuesta_traduccion = Net::HTTP.get(uri)
  datos_json = JSON.parse(respuesta_traduccion)
  
  # 5. Extraer texto traducido en español
  resultado_en_espanol = datos_json["responseData"]["translatedText"]
  
  # 6. Devolvemos el resultado al navegador
  resultado_en_espanol
end