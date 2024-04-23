const express = require('express');
const cors = require('cors');
const morgan = require('morgan');
const mongoose = require('mongoose');

// Create express app
const app = express();
const PORT = 5000;

// Peticiones
app.use(cors());

// Json format
app.use(morgan('dev'));
app.use(express.json());

// Connect to MongoDB
// IP de mongo-svc en el cluster 34.16.29.86
const ip = '34.16.29.86'

mongoose.connect(`mongodb://${ip}:27017/so1p2`)
const db = mongoose.connection;

db.on('error', console.error.bind(console, 'Error connecting to MongoDB:'));
db.once('open', () => {
    console.log('Connected to MongoDB');
});

const Band = mongoose.model('bands', {
    name: String,
    album: String,
    year: String,
    rank: String,
    date: Date
});

// Routes
app.get('/logs', async (req, res) => {
    try {
        const logs = await Band.find().sort({date: -1}).limit(20);
        res.json(logs);
    } catch (error) {
        console.error('Error getting logs:', error);
        res.status(500).send('Error getting logs');
    }
});

// Start express app
app.listen(PORT, () => {
    console.log(`Server running on port ${PORT}`);
});