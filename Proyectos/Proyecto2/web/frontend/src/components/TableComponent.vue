<template>
    <div class="title-text">
        <h1>Mongo Last 20 Logs</h1>
    </div>
    <div class="table-container">
        <div class="table-wrapper">
            <table class="custom-table">
                <thead>
                    <tr>
                        <th>Name</th>
                        <th>Album</th>
                        <th>Year</th>
                        <th>Rank</th>
                        <th>Date</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="(log, index) in logs" :key="index">
                        <td>{{ log.name }}</td>
                        <td>{{ log.album }}</td>
                        <td>{{ log.year }}</td>
                        <td>{{ log.rank }}</td>
                        <td>{{ log.date }}</td>
                    </tr>
                </tbody>
            </table>
        </div>
    </div>
</template>

<script>
export default {
    name: 'TableComponent',
    data() {
        return {
            logs: [] // Array para almacenar los logs
        }
    },
    mounted() {
        // Consultar el endpoint cada 1 segundo al montar el componente
        setInterval(this.fetchLogs, 2000);
    },
    methods: {
        // Método para consultar el endpoint y actualizar los logs
        fetchLogs() {
            fetch('http://localhost:5000/logs') // Cambia la URL según tu endpoint
                .then(response => response.json())
                .then(data => {
                    this.logs = data; // Actualiza los logs con los datos obtenidos del endpoint
                })
                .catch(error => {
                    console.error('Error fetching logs:', error);
                });
        }
    }
}
</script>

<style>
/* Estilos CSS para el título */
.title-text {
    text-align: center;
    margin-top: 5vh;
}

/* Estilos CSS para la tabla */
.table-container {
    max-height: 70vh;
    /* Altura máxima de la tabla */
    width: 80%;
    overflow-y: auto;
    /* Hacer la tabla desplazable verticalmente */
    margin-left: auto;
    margin-right: auto;
    margin-top: 5vh;
    border: 1px solid #ddd;
}

.table-wrapper {
    overflow-x: auto;
    /* Hacer que la tabla sea desplazable horizontalmente si es necesario */
}

.custom-table {
    width: 100%;
    /* Ajustar el ancho al 100% del contenedor */
    border-collapse: collapse;
    /* Colapsar los bordes de la tabla */
}

.custom-table th,
.custom-table td {
    padding: 8px;
    /* Agregar relleno a las celdas */
    text-align: center;
    /* Centrar el texto en las celdas */
    border: 1px solid #ddd;
    /* Añadir bordes a las celdas */
}

.custom-table th {
    background-color: #f2f2f2;
    /* Color de fondo para los encabezados */
    color: #000000;
    /* Color de texto para los encabezados */
}
</style>
