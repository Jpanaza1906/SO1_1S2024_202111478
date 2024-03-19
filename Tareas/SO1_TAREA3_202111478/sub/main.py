import redis
import json

# Conexión a Redis local
conexion_redis = redis.StrictRedis(host='10.91.212.67', port=6379, decode_responses=True)

# Función para manejar los mensajes recibidos
def handle_message(message):
    try:
        # Convertir el mensaje JSON a un diccionario Python
        data = json.loads(message['data'])
        print(f"Mensaje recibido en el canal {message['channel']}: {data['msg']}")
    except Exception as e:
        print(f"Error al procesar el mensaje: {e}")

# Crear un objeto de suscripción
subscripcion = conexion_redis.pubsub()

# Suscribirse al canal 'test'
subscripcion.subscribe('test')

# Escuchar mensajes
for message in subscripcion.listen():
    if message['type'] == 'message':
        handle_message(message)
