// main.go

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Data struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
	Date string `json:"date"`
}

func handleDataRequest(w http.ResponseWriter, r *http.Request) {
	// Añadir este caso para el método OPTIONS
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		return
	}

	if r.Method == "GET" {
		//se obtiene la fecha y hora actual del sistema
		currentTime := time.Now().Format("2006-01-02 15:04:05")

		sampleData := Data{
			Name: "Jose David Panaza Batres",
			ID:   202111478,
			Date: currentTime,
		}

		jsonData, err := json.Marshal(sampleData)
		if err != nil {
			http.Error(w, "Error al convertir a JSON", http.StatusInternalServerError)
			return
		}

		// Añadir encabezados CORS permitiendo solicitudes desde http://localhost:3000
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	} else {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

func main() {
	// Crear un nuevo manejador CORS
	corsHandler := http.HandlerFunc(handleDataRequest)

	// Usar el manejador CORS para la ruta /data
	http.Handle("/data", corsHandler)

	fmt.Println("Servidor escuchando en el puerto 5000...")
	http.ListenAndServe(":5000", nil)
}
