const express = require('express');
const cors = require('cors');
const morgan = require('morgan');
const mongoose = require('mongoose');


const app = express();
const PORT = 5000;

// Peticiones 
app.use(cors());

//Formato json
app.use(morgan('dev'));
app.use(express.json());

//mongodb://MongoDB:27017/clase2
//mongodb://localhost:27017/clase2

// Conexión a la base de datos
mongoose.connect('mongodb://localhost:27017/tarea2');
const db = mongoose.connection;

db.on('error', console.error.bind(console, 'Error de conexion a la base de datos:'));
db.once('open', () => {console.log('Conexión exitosa a la base de datos');});

const Imagen = mongoose.model('imgs',{
    imgb64: String,
    fecha: String
});

//obtener todos las imagenes
app.get('/imagenes', async (req, res) => {
    try{
        const imagenes = await Imagen.find();
        res.json(imagenes);
    } catch (error){
        console.error(error);
        res.status(500).json({error: 'Error al obtener las imagenes'});
    }
});

// Insertar una imagen, con dos campos: imgb64 y fecha
app.post('/imagenes', async (req, res) => {
    const {imgb64, fecha} = req.body;

    if (!imgb64 || !fecha){
        return res.status(400).json({error: 'imgb64 y fecha son campos requeridos'});
    }

    try{
        const nuevaImagen = new Imagen({imgb64, fecha});
        await nuevaImagen.save();
        res.status(201).json(nuevaImagen);
    } catch (error){
        console.error(error);
        res.status(500).json({error: 'Error al guardar la imagen'});
    }
});


app.listen(PORT, () => {
  console.log(`La aplicación está escuchando en el puerto ${PORT}`);
});
