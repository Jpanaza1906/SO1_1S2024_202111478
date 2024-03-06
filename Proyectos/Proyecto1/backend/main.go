package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

//============================================= Conexion a la base de datos =============================================

var conexion = ConexionMysql()

func ConexionMysql() *sql.DB {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)

	conexion, err := sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error al abrir la conexion con la base de datos: ", err)
	}

	_, err = conexion.Exec("SELECT 1")
	if err != nil {
		log.Fatal("La conexion con la base de datos no es correcta: ", err)
	}

	fmt.Println("Conexión exitosa con la base de datos")

	return conexion
}

//============================================= Inicialización del servidor =============================================

func main() {
	router := mux.NewRouter().StrictSlash(true)
	//================== Rutas del servidor ==================
	router.HandleFunc("/monitor", monitor).Methods("GET")
	//router.HandleFunc("/", indexRoute)
	//router.HandleFunc("/Registrar", registro).Methods("POST")
	//router.HandleFunc("/Estudiantes", getEstudiantes).Methods("GET")

	//================== Exponer el puerto del servidor ==================
	fmt.Println("Server on port", 8000)
	handler := cors.Default().Handler(router)
	log.Fatal((http.ListenAndServe(":8000", handler)))
	http.Handle("/", router)
}

// ============================================= FUNCIONES =============================================
func monitor(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("Monitor")

	// Obtener datos de la RAM
	ramPercentage, err := getRAMdata()
	if err != nil {
		fmt.Println("Error al obtener datos de la RAM: ", err)
		http.Error(w, "Error al obtener datos de la RAM", http.StatusInternalServerError)
		return
	}

	// Obtener datos de la CPU
	cpuPercentage, err := getCPUdata()
	if err != nil {
		fmt.Println("Error al obtener datos de la CPU: ", err)
		http.Error(w, "Error al obtener datos de la CPU", http.StatusInternalServerError)
		return
	}

	complementoCPU := 100 - cpuPercentage
	complementoRAM := 100 - ramPercentage

	data := map[string][]int{
		"cpu_percentage": {complementoCPU, cpuPercentage},
		"ram_percentage": {complementoRAM, ramPercentage},
	}

	// Convertir la estructura a formato JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error al convertir los datos a JSON: ", err)
		http.Error(w, "Error al convertir los datos a JSON", http.StatusInternalServerError)
		return
	}

	// Enviar la respuesta
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)

}

type SystemData struct {
	CPU_percentage int `json:"cpu_percentage"`
	RAM_percentage int `json:"ram_percentage"`
}

// Funcion para obtener datos de la RAM
func getRAMdata() (int, error) {
	cmd := exec.Command("sh", "-c", "cat /proc/ram_so1_1s2024")
	stdout, err := cmd.CombinedOutput()

	if err != nil {
		return 0, err
	}

	// Convertir la salida a formato JSON
	var data SystemData
	err = json.Unmarshal(stdout, &data)
	if err != nil {
		return 0, err
	}

	return data.RAM_percentage, nil
}

// Funcion para obtener datos de la CPU
func getCPUdata() (int, error) {
	cmd := exec.Command("sh", "-c", "cat /proc/cpu_so1_1s2024")
	stdout, err := cmd.CombinedOutput()

	if err != nil {
		return 0, err
	}

	// Convertir la salida a formato JSON
	var data SystemData
	err = json.Unmarshal(stdout, &data)
	if err != nil {
		return 0, err
	}

	return data.CPU_percentage, nil
}
