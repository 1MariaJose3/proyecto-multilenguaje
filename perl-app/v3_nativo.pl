use strict;
use warnings;
use Plack::Request;
use Lingua::ES::Numeros; 

my $app = sub {
    my $env = shift;
    my $req = Plack::Request->new($env);
    
    # 1. Capturamos el número 'n' de la URL (ej. ?n=10)
    my $n = int($req->param('n') // 0);
    
    # 2. Inicializamos el objeto convertidor de la librería nativa
    my $convertidor = Lingua::ES::Numeros->new();
    
    # 3. Traducimos el número a palabras utilizando el método de la librería
    my $resultado_letras = $convertidor->cardinal($n); # Si n=10, devuelve "diez"
    
    # 4. Devolvemos el resultado al navegador en minúsculas para mantener la consistencia
    return [
        200,
        [ 'Content-Type' => 'text/plain; charset=utf-8' ],
        [ lc($resultado_letras) ]
    ];
};