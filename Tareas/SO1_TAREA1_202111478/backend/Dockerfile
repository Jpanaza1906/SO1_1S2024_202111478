# Fase de construcción
FROM golang:latest AS builder

WORKDIR /app

# Copiar solo los archivos necesarios para descargar las dependencias
COPY go.mod go.sum ./

# Descargar dependencias
RUN go mod download

# Copiar todo el código fuente
COPY . .

# Compilar la aplicación
RUN go build -o main .

# Fase de ejecución
FROM golang:latest

WORKDIR /app

# Copiar solo los archivos necesarios desde la fase de construcción
COPY --from=builder /app/main .

# Exponer el puerto en el que tu aplicación escucha
EXPOSE 5000

# Comando para ejecutar la aplicación
CMD ["./main"]
