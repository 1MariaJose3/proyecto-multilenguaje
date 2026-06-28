use strict;
use warnings;
use Plack::Request;
use SOAP::Lite;

# Definimos que la aplicación correrá en un entorno web básico
my $app = sub {
    my $env = shift;
    my $req = Plack::Request->new($env);
    
    # 1. Capturamos el parámetro 'n' desde la URL (ej. ?n=10)
    my $n = $req->param('n') // 0;
    
    # 2. Conectamos y configuramos el servicio web SOAP público
    my $soap = SOAP::Lite
        ->proxy('https://www.dataaccess.com/webservicesserver/NumberConversion.wso')
        ->uri('http://www.dataaccess.com/webservicesserver/');
        
    # 3. Llamamos al método NumberToWords enviando el parámetro
    my $som = $soap->NumberToWords(
        SOAP::Data->name('ubiNum' => $n)
    );
    
    # 4. Extraemos el resultado en texto (ej. "ten")
    my $resultado_en_ingles = $som->result;
    
    # 5. Devolvemos la respuesta al navegador en texto plano
    return [
        200,
        [ 'Content-Type' => 'text/plain; charset=utf-8' ],
        [ $resultado_en_ingles ]
    ];
};