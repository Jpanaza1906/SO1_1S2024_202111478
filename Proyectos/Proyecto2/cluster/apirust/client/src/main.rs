//Importaciones
use rocket::{routes, serde::json::Json};
use rocket::post;
use reqwest::Client;
use serde::{Deserialize, Serialize};

//Estructura de datos
#[derive(Debug, Serialize, Deserialize)]
struct Data {
    name: String,
    album: String,
    year: String,
    rank: String,
}

//Funcion para mandar la data /send_data
#[post("/send_data", data = "<data>")]
async fn send_data(data: Json<Data>) -> String {
    let client = Client::new();
    let server_url = "http://localhost:8080/data";
    let response = client.post(server_url)
        .json(&data.into_inner())
        .send()
        .await;

    match response {
        Ok(_) => "Data sent successfully".to_string(),
        Err(_) => "Failed to send data".to_string(),
    }
}

//Funcion principal
#[rocket::main]
async fn main() {
    let rocket = rocket::build()
        .mount("/", routes![send_data])
        .launch();

    let _ = rocket.await;
    
}