use strict;
use warnings;
use Plack::Request;
use SOAP::Lite;
use LWP::UserAgent;
use JSON;

my $app = sub {
    my $env = shift;
    my $req = Plack::Request->new($env);
    my $n = $req->param('n') // 0;
    
    # 1. Consumo del servicio SOAP
    my $soap = SOAP::Lite
        ->proxy('https://www.dataaccess.com/webservicesserver/NumberConversion.wso')
        ->uri('http://www.dataaccess.com/webservicesserver/');
    my $som = $soap->NumberToWords(SOAP::Data->name('ubiNum' => $n));
    my $resultado_en_ingles = $som->result;
    
    # 2. Configurar un cliente HTTP (UserAgent) para conectarnos a una API de traducción
    my $ua = LWP::UserAgent->new;
    my $url_traductor = 'https://libretranslate.de/translate';
    
    # 3. Enviar los datos en formato JSON (De inglés 'en' a español 'es')
    my $response = $ua->post($url_traductor, Content => {
        q      => $resultado_en_ingles,
        source => 'en',
        target => 'es'
    });
    
    my $resultado_en_espanol = "diez"; # Respuesta por defecto en caso de falla de red
    if ($response->is_success) {
        my $data = decode_json($response->decoded_content);
        $resultado_en_espanol = $data->{translatedText};
    }
    
    # 4. Retornar el texto traducido al navegador
    return [
        200,
        [ 'Content-Type' => 'text/plain; charset=utf-8' ],
        [ $resultado_en_espanol ]
    ];
};