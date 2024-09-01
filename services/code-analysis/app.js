const babelParser = require('@babel/parser');
const express = require('express');
const cors = require('cors');
const app = express();
const PORT = 3000;

app.use(express.json());
app.use(cors()); // Enable CORS for all routes

app.get('/', (req, res) => {
    res.send('Code Analyzer API is running...');
});

app.listen(PORT, '0.0.0.0', () => {
    console.log(`Server is running on port ${PORT}`);
});









// Endpoint to parse JavaScript code and return the AST
app.post('/parse/javascript', (req, res) => {
    console.log('Request received', req.body);
    const { code } = req.body;
    try {
        const ast = babelParser.parse(code, {
            sourceType: 'module'
        });

        res.json(ast);
    } catch (error) {
        res.status(400).json({ error: error.message });
    }
});