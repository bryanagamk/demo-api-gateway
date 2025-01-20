const express = require('express');
const app = express();
const port = 9002;

app.get('/', (req, res) => {
    res.send('Hello World from Service 2 (Express)');
});

app.listen(port, () => {
    console.log(`Service 2 (Express) running on http://localhost:${port}`);
});