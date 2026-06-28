require 'sinatra'
require 'savon'

# Configuramos el servidor en el puerto 8000
set :port, 8000

# Ruta: http://localhost:8000/clisoap1?n=10
get '/clisoap1' do
  # Conectamos con el servicio web público que te dio tu profesor
  client = Savon.client(wsdl: "https://www.dataaccess.com/webservicesserver/NumberConversion.wso?WSDL")
  
  # Enviamos el número 'n' que viene en la URL
  response = client.call(:number_to_words, message: { "ubiNum" => params[:n] })
  
  # Devolvemos el resultado en texto (ej. "ten")
  response.body[:number_to_words_response][:number_to_words_result]
end



