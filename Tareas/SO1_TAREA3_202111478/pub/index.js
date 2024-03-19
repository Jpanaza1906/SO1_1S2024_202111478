const Redis = require('ioredis');

const conexion = new Redis({
    host: '10.91.212.67',
    port: 6379,
    connectTimeout: 5000,
});

function funcionPub() {
    // Objeto JSON a enviar
    const jsonMsg = {
        msg: 'Hola a todos'
    };

    // Convertir el objeto JSON a una cadena y enviarlo a través de publish
    conexion.publish('test', JSON.stringify(jsonMsg))
    .then(() => {
        console.log('Mensaje enviado');
    })
    .catch((error) => {
        console.error('Error:', error);
    });
}

setInterval(funcionPub, 3000);
