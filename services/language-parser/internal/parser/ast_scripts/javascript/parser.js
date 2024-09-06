const babelParser = require('@babel/parser');
const traverse = require('@babel/traverse').default;

// Read the JavaScript code from stdin
let input = '';

process.stdin.on('data', function(chunk) {
    input += chunk;
});

process.stdin.on('end', function() {
    // Parse the JavaScript code to generate the AST
    const ast = babelParser.parse(input, {
        sourceType: "module", // You can change this based on your needs
    });

    // Normalized AST structure
    let normalisedAST = {
        functions: [],
        relationships: [],
        conditionals: []
    };

    let callStack = [];  // Track function calls and their order

    traverse(ast, {
        FunctionDeclaration(path) {
            const functionName = path.node.id ? path.node.id.name : 'anonymous';
            const line = path.node.loc.start.line;

            // Add to normalized function list
            normalisedAST.functions.push({
                name: functionName,
                line: line
            });
        },
        CallExpression(path) {
            const callerFunction = path.findParent((p) => p.isFunctionDeclaration());
            const callerName = callerFunction ? callerFunction.node.id.name : 'anonymous';
            const calledFunctionName = path.node.callee.name;
            const line = path.node.loc.start.line;

            // Check if the call is part of a conditional statement
            if (path.findParent((p) => p.isConditionalExpression() || p.isIfStatement())) {
                const conditionalType = path.findParent((p) => p.isIfStatement()) ? 'if' : 'else';

                // Add conditional call to relationships
                normalisedAST.relationships.push({
                    source: callerName,
                    target: calledFunctionName,
                    type: `CONDITIONAL_CALL (${conditionalType})`,
                    line: line
                });
            } else {
                // Regular function call
                callStack.push({ caller: callerName, callee: calledFunctionName, line: line });

                // Add regular call to relationships
                normalisedAST.relationships.push({
                    source: callerName,
                    target: calledFunctionName,
                    type: "CALLS",
                    line: line,
                    order: callStack.length
                });
            }
        },
        IfStatement(path) {
            const parentFunctionName = path.findParent((p) => p.isFunctionDeclaration())?.node.id.name || 'anonymous';

            // Extract a meaningful description of the condition
            let conditionDescription = '';
            if (path.node.test.type === 'CallExpression' && path.node.test.callee.name) {
                conditionDescription = `${path.node.test.callee.name}(${path.node.test.arguments.map(arg => arg.name || arg.value).join(', ')})`;
            } else {
                conditionDescription = path.node.test.type;
            }

            const ifLine = path.node.loc.start.line;

            // Add conditional relationships to conditionals
            normalisedAST.conditionals.push({
                source: parentFunctionName,
                target: parentFunctionName,
                type: `CONDITIONAL (if, ${conditionDescription})`,
                line: ifLine
            });

            // Handle else case if available
            if (path.node.alternate) {
                const elseLine = path.node.alternate.loc.start.line;
                normalisedAST.conditionals.push({
                    source: parentFunctionName,
                    target: parentFunctionName,
                    type: "CONDITIONAL (else)",
                    line: elseLine
                });
            }
        }
    });

    // Output the normalized AST as a JSON string
    console.log(JSON.stringify(normalisedAST, null, 2));
    return normalisedAST;
});
