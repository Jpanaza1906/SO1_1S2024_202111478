FROM node:18.18.2-alpine

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app
COPY package*.json ./

# Instala las dependencias
RUN npm install

# Copia el resto de los archivos de la aplicación
COPY . .

# Construye la aplicación para producción
RUN npm run build

# Expone el puerto en el que la aplicación se ejecutará
EXPOSE 3000

# Comando para iniciar la aplicación cuando se ejecute el contenedor
CMD ["npm", "start"]

#Notas:
#docker build -t front_clase1 .
#docker run -p 3000:3000 front_clase1