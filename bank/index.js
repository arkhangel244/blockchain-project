const express = require('express');
const app = express();
const axios = require('axios');


// Define routes
app.post('/adduser', async (req, res) => {
    try {
        // Forward the request to the Python server
        const response = await axios.post('http://localhost:12345/adduser', req.body);
        res.status(response.status).json(response.data);
    } catch (error) {
        console.error('Error:', error);
        res.status(500).json({ error: 'Internal server error' });
    }
});

app.post('/signin', async (req, res) => {
    try {
        // Forward the request to the Python server
        const response = await axios.post('http://localhost:12345/signin', req.body);
        res.status(response.status).json(response.data);
    } catch (error) {
        console.error('Error:', error);
        res.status(500).json({ error: 'Internal server error' });
    }
});

// Start the server
const PORT = 3000;
app.listen(PORT, () => {
    console.log(`Node.js server is running on port ${PORT}`);
});

