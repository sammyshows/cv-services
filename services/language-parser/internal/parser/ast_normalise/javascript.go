package ast_normalise

// // NormaliseJavaScriptAST converts the JavaScript AST into a normalized structure
// func NormaliseJavaScriptAST(ast interface{}) (types.NormalisedAST, error) {
// 	astMap, ok := ast.(map[string]interface{})
// 	if !ok {
// 		return types.NormalisedAST{}, fmt.Errorf("invalid AST format")
// 	}

// 	var normalized types.NormalisedAST
// 	callStack := []types.Relationship{}

// 	// For each function declaration in the AST
// 	for _, functionNode := range getFunctionDeclarations(astMap) {
// 		functionName := functionNode.Name
// 		if functionName == "" {
// 			functionName = "anonymous"
// 		}
// 		lineNumber := functionNode.Line

// 		// Add to Functions array
// 		normalized.Functions = append(normalized.Functions, types.Function{
// 			Name: functionName,
// 			Line: lineNumber,
// 		})
// 	}

// 	// For each call expression in the AST
// 	for _, callNode := range getCallExpressions(astMap) {
// 		callerFunction := findParentFunction(callNode)
// 		callerName := "anonymous"
// 		if callerFunction != nil {
// 			callerName = callerFunction.Name
// 		}
// 		calledFunctionName := callNode.CalleeName
// 		line := callNode.Line

// 		// Check if the call is part of a conditional statement
// 		if isConditionalExpression(callNode) {
// 			conditionalType := getConditionalType(callNode)
// 			normalized.Relationships = append(normalized.Relationships, types.Relationship{
// 				Source: callerName,
// 				Target: calledFunctionName,
// 				Type:   fmt.Sprintf("CONDITIONAL_CALL (%s)", conditionalType),
// 				Line:   line,
// 			})
// 		} else {
// 			callStack = append(callStack, types.Relationship{
// 				Source: callerName,
// 				Target: calledFunctionName,
// 				Type:   "CALLS",
// 				Line:   line,
// 			})
// 		}
// 	}

// 	// For each if statement in the AST
// 	for _, ifNode := range getIfStatements(astMap) {
// 		parentFunction := findParentFunction(ifNode)
// 		parentFunctionName := "anonymous"
// 		if parentFunction != nil {
// 			parentFunctionName = parentFunction.Name
// 		}

// 		// Get condition description
// 		conditionDescription := extractConditionDescription(ifNode)

// 		// Create relationships for if and else conditions
// 		ifLine := ifNode.Line
// 		normalized.Relationships = append(normalized.Relationships, types.Relationship{
// 			Source: parentFunctionName,
// 			Target: parentFunctionName,
// 			Type:   fmt.Sprintf("CONDITIONAL (if, %s)", conditionDescription),
// 			Line:   ifLine,
// 		})

// 		// Handle else case if available
// 		if ifNode.HasElse {
// 			elseLine := ifNode.ElseLine
// 			normalized.Relationships = append(normalized.Relationships, types.Relationship{
// 				Source: parentFunctionName,
// 				Target: parentFunctionName,
// 				Type:   "CONDITIONAL (else)",
// 				Line:   elseLine,
// 			})
// 		}
// 	}

// 	return normalized, nil
// }

// // getFunctionDeclarations returns all the function declarations in the AST
// func getFunctionDeclarations(ast []types.CallNode) []types.FunctionNode {
// 	var functions []types.FunctionNode

// 	// Traverse the Babel AST body (assuming AST structure based on Babel output)
// 	body := ast["program"].(map[string]interface{})["body"].([]interface{})

// 	for _, node := range body {
// 		nodeMap := node.(map[string]interface{})
// 		if nodeMap["type"] == "FunctionDeclaration" {
// 			name := nodeMap["id"].(map[string]interface{})["name"].(string)
// 			line := int(nodeMap["loc"].(map[string]interface{})["start"].(map[string]interface{})["line"].(float64))

// 			functions = append(functions, types.FunctionNode{Name: name, Line: line})
// 		}
// 	}

// 	return functions
// }

// // getCallExpressions returns all the call expressions in the AST
// func getCallExpressions(ast map[string]interface{}) []types.CallNode {
// 	var calls []types.CallNode

// 	body := ast["program"].(map[string]interface{})["body"].([]interface{})

// 	for _, node := range body {
// 		nodeMap := node.(map[string]interface{})
// 		if nodeMap["type"] == "ExpressionStatement" {
// 			expression := nodeMap["expression"].(map[string]interface{})
// 			if expression["type"] == "CallExpression" {
// 				calleeName := expression["callee"].(map[string]interface{})["name"].(string)
// 				line := int(nodeMap["loc"].(map[string]interface{})["start"].(map[string]interface{})["line"].(float64))

// 				calls = append(calls, types.CallNode{CalleeName: calleeName, Line: line})
// 			}
// 		}
// 	}

// 	return calls
// }

// // getIfStatements returns all the if statements in the AST
// func getIfStatements(ast map[string]interface{}) []types.IfNode {
// 	var ifStatements []types.IfNode

// 	body := ast["program"].(map[string]interface{})["body"].([]interface{})

// 	for _, node := range body {
// 		nodeMap := node.(map[string]interface{})
// 		if nodeMap["type"] == "IfStatement" {
// 			line := int(nodeMap["loc"].(map[string]interface{})["start"].(map[string]interface{})["line"].(float64))
// 			var hasElse bool
// 			var elseLine int

// 			if nodeMap["alternate"] != nil {
// 				hasElse = true
// 				elseLine = int(nodeMap["alternate"].(map[string]interface{})["loc"].(map[string]interface{})["start"].(map[string]interface{})["line"].(float64))
// 			}

// 			ifStatements = append(ifStatements, types.IfNode{Line: line, HasElse: hasElse, ElseLine: elseLine})
// 		}
// 	}

// 	return ifStatements
// }

// // findParentFunction finds the parent function of a given node (this would need context traversal)
// func findParentFunction(node map[string]interface{}) *types.FunctionNode {
// 	// Simulating traversal up the tree; in practice, you'd need access to the parent node or use a stack
// 	// In this case, let's assume you have a field in the node that tracks the parent function
// 	if node["parentFunction"] != nil {
// 		parent := node["parentFunction"].(map[string]interface{})
// 		name := parent["name"].(string)
// 		line := int(parent["loc"].(map[string]interface{})["start"].(map[string]interface{})["line"].(float64))

// 		return &types.FunctionNode{Name: name, Line: line}
// 	}

// 	return nil
// }

// // isConditionalExpression checks if the given node is part of a conditional statement
// func isConditionalExpression(node map[string]interface{}) bool {
// 	// Check if node is inside a conditional expression or statement
// 	if node["type"] == "ConditionalExpression" || node["type"] == "IfStatement" {
// 		return true
// 	}
// 	return false
// }

// // getConditionalType returns the type of conditional (e.g., if, else)
// func getConditionalType(node map[string]interface{}) string {
// 	if node["type"] == "IfStatement" {
// 		return "if"
// 	} else if node["type"] == "ConditionalExpression" {
// 		return "else"
// 	}
// 	return "unknown"
// }

// // extractConditionDescription generates a description of the condition for if statements
// func extractConditionDescription(node types.IfNode) string {
// 	// Implement logic to generate a condition description based on the structure of IfNode
// 	return fmt.Sprintf("If statement at line %d", node.Line)
// }
