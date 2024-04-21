//Importaciones
use rocket::response::status::BadRequest;
use rocket::serde::json::{json, Value as JsonValue};
use rocket::serde::json::Json;

//Estructura de datos
#[derive(rocket::serde::Deserialize)]
struct Data {
    name: String,
    album: String,
    year: String,
    rank: String
}

//Definir el endpoint /data
#[rocket::post("/data", data = "<data>")]
fn receive_data(data: Json<Data>) -> Result<String, BadRequest<String>> {
    let receive_data = data.into_inner();
    let response = JsonValue::from(json!({
        "message": format!("Received data: Name: {}, Album: {}, Year: {}, Rank: {}", receive_data.name, receive_data.album, receive_data.year, receive_data.rank)
    }));
    Ok(response.to_string())
}

//Funcion main
#[rocket::main]
async fn main() {
    let config = rocket::Config{
        port : 8080,
        ..rocket::Config::default()
    };
    rocket::custom(config)
        .mount("/", rocket::routes![receive_data])
        .launch()
        .await
        .unwrap();
}