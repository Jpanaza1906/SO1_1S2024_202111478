//Importaciones
use rocket::serde::json::{json, Value as JsonValue};
use rocket::serde::json::Json;
use rocket::config::SecretKey;
use rocket_cors::{AllowedOrigins, CorsOptions};
use rocket::http::Status;

use std::time::Duration;
use rdkafka::producer::{FutureProducer, FutureRecord};
use rdkafka::ClientConfig;

//Estructura de datos
#[derive(rocket::serde::Deserialize)]
struct Data {
    name: String,
    album: String,
    year: String,
    rank: String
}

//Funcion para enviar datos a Kafka
async fn produce(data: &Data) -> Result<(), Box<dyn std::error::Error>> {
    // Configurar la dirección del broker y el nombre del tema Kafka
    let broker_address = "my-cluster-kafka-bootstrap:9092";
    let kafka_topic = "topic-sopes1";
    
    // Configurar el cliente Kafka
    let producer: FutureProducer = ClientConfig::new()
        .set("bootstrap.servers", broker_address)
        .create()?;
    
    // Crear el mensaje a enviar
    let message_value = format!(
        r#"{{"name":"{}","album":"{}","year":"{}","rank":"{}"}}"#,
        data.name, data.album, data.year, data.rank
    );

    // Construir y enviar el mensaje
    let record = FutureRecord::to(kafka_topic)
        .key(&data.rank)
        .payload(&message_value);
    
    match producer.send(record, Duration::from_secs(0)).await {
        Ok(_) => println!("Mensaje enviado exitosamente"),
        Err((e, _)) => eprintln!("Error al enviar mensaje: {}", e),
    }

    Ok(())
}



//Definir el endpoint /data
#[rocket::post("/data", data = "<data>")]
async fn receive_data(data: Json<Data>) -> Result<String, Status> {
    let received_data = data.into_inner();
    //Imprimir en consola el mensaje recibido
    println!("Received data: Name: {}, Album: {}, Year: {}, Rank: {}", received_data.name, received_data.album, received_data.year, received_data.rank);

    // Llamar a la función produce con los datos recibidos
    match produce(&received_data).await {
        Ok(_) => {
            let response = JsonValue::from(json!({
                "message": "Data received and sent to Kafka successfully"
            }));
            Ok(response.to_string())
        },
        Err(e) => {
            eprintln!("Error while producing message to Kafka: {}", e);
            Err(Status::InternalServerError)
        }
    }
}

//Funcion main
#[rocket::main]
async fn main() {
    let secret_key = SecretKey::generate();

    // configuracion de opciones CORS
    let _cors = CorsOptions::default()
        .allowed_origins(AllowedOrigins::all())
        .to_cors()
        .expect("failed to create CORS fairing");

    let config = rocket::Config{
        address : "0.0.0.0".parse().unwrap(),
        port : 8080,
        secret_key : secret_key.unwrap(),
        ..rocket::Config::default()
    };
    rocket::custom(config)
        .mount("/", rocket::routes![receive_data])
        .launch()
        .await
        .unwrap();
}