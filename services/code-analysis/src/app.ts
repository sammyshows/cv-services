const babelParser = require('@babel/parser');
const traverse = require('@babel/traverse').default;
const express = require('express');
const cors = require('cors');
import { neo4jDriver, getNeo4jSession } from './database'

const app = express();
const PORT = 3000;

app.use(express.json());
app.use(cors()); // Enable CORS for all routes


app.get('/', (req, res) => {
    console.log('Request received');
    res.send('Code Analyzer API is running...');
});

app.listen(PORT, '0.0.0.0', () => {
    console.log(`Server is running on port ${PORT}`);
});

process.on('exit', async () => {
    await neo4jDriver.close();
});




app.post('/parse/javascript', async (req, res) => {
    const { code } = req.body;
    const neo4j = getNeo4jSession();
    console.log('Request received');
    
    try {
        const ast = babelParser.parse(code, { sourceType: 'module' });
        let queries = [];
        let callStack = [];  // Stack to track function calls and their order

        traverse(ast, {
            FunctionDeclaration(path) {
                const functionName = path.node.id ? path.node.id.name : 'anonymous';
                const functionIdQuery = `CREATE (f:Function {name: "${functionName}", line: ${path.node.loc.start.line}}) RETURN id(f) AS id`;
                queries.push(functionIdQuery);
            },
            CallExpression(path) {
                const callerFunction = path.findParent((p) => p.isFunctionDeclaration());
                const callerName = callerFunction ? callerFunction.node.id.name : 'anonymous';
                const calledFunctionName = path.node.callee.name;
                const line = path.node.loc.start.line;

                if (path.findParent((p) => p.isConditionalExpression() || p.isIfStatement())) {
                    // This call is part of a conditional statement
                    const conditionalType = path.findParent((p) => p.isIfStatement()) ? 'if' : 'else';
                    const relationshipQuery = `MATCH (caller:Function {name: "${callerName}"}), (callee:Function {name: "${calledFunctionName}"})
                    CREATE (caller)-[:CONDITIONAL_CALL {type: "${conditionalType}", line: ${line}}]->(callee)`;
                    queries.push(relationshipQuery);
                } else {
                    // Regular function call
                    callStack.push({ caller: callerName, callee: calledFunctionName, line: line });
                    const relationshipQuery = `MATCH (caller:Function {name: "${callerName}"}), (callee:Function {name: "${calledFunctionName}"})
                    CREATE (caller)-[:CALLS {line: ${line}, order: ${callStack.length}}]->(callee)`;
                    queries.push(relationshipQuery);
                }
            },
            IfStatement(path) {
                const parentFunctionName = path.findParent((p) => p.isFunctionDeclaration())?.node.id.name || 'anonymous';
                // Extract a meaningful description of the condition
                let conditionDescription = '';
                if (path.node.test.type === 'CallExpression' && path.node.test.callee.name) {
                    // When the test condition is a function call
                    conditionDescription = `${path.node.test.callee.name}(${path.node.test.arguments.map(arg => arg.name || arg.value).join(', ')})`;
                } else {
                    // Generic fallback for other types of conditions
                    conditionDescription = path.node.test.type;
                }

                const ifLine = path.node.loc.start.line;

                const ifQuery = `MATCH (func:Function {name: "${parentFunctionName}"})
                CREATE (func)-[:CONDITIONAL {type: "if", line: ${ifLine}, condition: "${conditionDescription}" }]->(func)`;
                queries.push(ifQuery);

                // Handling else
                if (path.node.alternate) {
                    const elseLine = path.node.alternate.loc.start.line;
                    const elseQuery = `MATCH (func:Function {name: "${parentFunctionName}"})
                    CREATE (func)-[:CONDITIONAL {type: "else", line: ${elseLine} }]->(func)`;
                    queries.push(elseQuery);
                }
            }
        });

        // Execute all Cypher queries
        const results = [];
        console.log('Generated Cypher queries:', queries);
        for (const query of queries) {
            const result = await neo4j.run(query);
            results.push(result.records);
        }
        console.log('Successfully executed all queries.');
        res.json({ ast, results });
    } catch (error) {
        console.error('Error:', error);
        res.status(400).json({ error: error.message });
    } finally {
        await neo4j.close();
    }
});


app.get('/functions', async (req, res) => {
    const query = `
        MATCH (f:Function)
        OPTIONAL MATCH (f)-[call:CALLS]->(called:Function)
        OPTIONAL MATCH (f)-[cond:CONDITIONAL]->(cf:Function)
        RETURN f AS Function, 
               collect(DISTINCT {type: call.type, callee: called.name, line: call.line, order: call.order}) AS Calls,
               collect(DISTINCT {type: cond.type, condition: cf.name, line: cond.line}) AS Conditionals
        ORDER BY f.line`;

    const neo4j = getNeo4jSession();
    try {
        const result = await neo4j.run(query);
        const formattedResult = result.records.map(record => ({
            function: record.get('Function').properties,
            calls: record.get('Calls'),
            conditionals: record.get('Conditionals')
        }));
        console.log(formattedResult);
        res.json(formattedResult);
    } catch (error) {
        console.error('Error fetching function data:', error);
        res.status(500).json({ error: error.message });
    } finally {
        await neo4j.close();
    }
});

