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
const ip = '34.172.114.76';
const dbName = 'so1p2';
const username = 'jpanaza';
const password = 'so1p2';

const uri = `mongodb://${username}:${password}@${ip}:27017/${dbName}`;

mongoose.connect(uri, {
  useNewUrlParser: true,
  useUnifiedTopology: true
});

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

// Rouute with nothing
app.get('/api', (req, res) => {
    res.send('Hello World');
});

// Routes
app.get('/api/logs', async (req, res) => {
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
