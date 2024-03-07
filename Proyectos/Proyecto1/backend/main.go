package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"

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
	//Vaciar la tabla monitor de la base de datos
	_, err := conexion.Exec("TRUNCATE TABLE monitor")
	if err != nil {
		log.Fatal("Error al vaciar la tabla monitor: ", err)
	}

	//================== Crear el router del servidor ==================
	router := mux.NewRouter().StrictSlash(true)
	//================== Rutas del servidor ==================
	router.HandleFunc("/monitor", monitor).Methods("GET")
	//router.HandleFunc("/", indexRoute)
	//router.HandleFunc("/Registrar", registro).Methods("POST")
	//router.HandleFunc("/Estudiantes", getEstudiantes).Methods("GET")

	//================== Exponer el puerto del servidor ==================

	//hacer una go routine para que el servidor este escuchando en el puerto 8000

	go func() {
		fmt.Println("Server on port", 8000)
		handler := cors.Default().Handler(router)
		log.Fatal((http.ListenAndServe(":8000", handler)))
		http.Handle("/", router)
	}()

	select {}
}

// ============================================= FUNCIONES de RUTAS =============================================

// ============================================= Funcion para el monitoreo del sistema==================================
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

	// Insertar los datos en la base de datos

	err = insertData(ramPercentage, cpuPercentage)
	if err != nil {
		fmt.Println("Error al insertar los datos en la base de datos: ", err)
		http.Error(w, "Error al insertar los datos en la base de datos", http.StatusInternalServerError)
		return
	}

	//Consultar los ultimos 30 registros de mi tabla monitor en la base de datos
	registros, err := getRegistros()
	if err != nil {
		fmt.Println("Error al obtener los registros de la base de datos: ", err)
		http.Error(w, "Error al obtener los registros de la base de datos", http.StatusInternalServerError)
		return
	}

	//variables para almacenar los registros
	var ramData []int
	var cpuData []int
	var labels []string

	//recorrer los registros de atras hacia adelante
	for i := len(registros) - 1; i >= 0; i-- {
		labels = append(labels, registros[i].Fecha.Format("2006-01-02 15:04:05"))
		ramData = append(ramData, registros[i].Usoram)
		cpuData = append(cpuData, registros[i].Usocpu)
	}

	// construir el objeto JSON

	dataHistorial := map[string]interface{}{
		"labels": labels,
		"datasets": []map[string]interface{}{
			{
				"label":           "RAM",
				"data":            ramData,
				"borderColor":     "#94d2bd",
				"backgroundColor": "#94d2bd",
				"borderWidth":     1,
				"tension":         0.5,
				"fill":            false,
				"pointRadius":     1,
			},
			{
				"label":           "CPU",
				"data":            cpuData,
				"borderColor":     "#ee9b00",
				"backgroundColor": "#ee9b00",
				"borderWidth":     1,
				"tension":         0.5,
				"fill":            false,
				"pointRadius":     1,
			},
		},
	}

	// Estructura para la respuesta JSON

	type Response struct {
		DataHistorial map[string]interface{} `json:"data_historial"`
		Data          map[string][]int       `json:"data"`
	}

	// Construir la estructura de respuesta

	response := Response{
		DataHistorial: dataHistorial,
		Data: map[string][]int{
			"cpu_percentage": {complementoCPU, cpuPercentage},
			"ram_percentage": {complementoRAM, ramPercentage},
		},
	}

	// Convertir la estructura a formato JSON
	jsonData, err := json.Marshal(response)
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

// Estructura para los datos del sistema
type SystemData struct {
	CPU_percentage int `json:"cpu_percentage"`
	RAM_percentage int `json:"ram_percentage"`
}

// Estructura para los registros
type Registro struct {
	ID     int       `json:"id"`
	Usoram int       `json:"usoram"`
	Usocpu int       `json:"usocpu"`
	Fecha  time.Time `json:"fecha"`
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

// Funcion para insertar los datos en la base de datos
func insertData(ramPercentage int, cpuPercentage int) error {
	//Preparar la consulta SQL para insertar los datos
	query := "INSERT INTO monitor (usoram, usocpu, fecha) VALUES (?,?,?)"
	_, err := conexion.Exec(query, ramPercentage, cpuPercentage, time.Now())
	if err != nil {
		return err
	}
	return nil
}

// Funcion para obtener los ultimos 30 registros de la tabla monitor
func getRegistros() ([]Registro, error) {
	//Preparar la consulta SQL para obtener los registros
	query := "SELECT * FROM monitor ORDER BY id DESC LIMIT 30"

	//Ejecutar la consulta
	rows, err := conexion.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	//Crear un slice para almacenar los registros
	var registros []Registro
	for rows.Next() {
		var registro Registro
		err := rows.Scan(&registro.ID, &registro.Usoram, &registro.Usocpu, &registro.Fecha)
		if err != nil {
			return nil, err
		}
		registros = append(registros, registro)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return registros, nil
}
