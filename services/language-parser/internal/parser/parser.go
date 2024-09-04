package parser

import (
	"encoding/json"
	"fmt"
	"language-parser/internal/parser/types"
	"log"
	"os/exec"
	"strings"
)

// ParseCode accepts code and language, then calls the appropriate parser subprocess
func ParseContent(code string, language string) (interface{}, error) {
	var cmd *exec.Cmd

	// Switch logic to decide the parser based on the language
	switch language {
	case "javaScript":
		// Call JavaScript parser subprocess
		// cmd = exec.Command("node", "path_to_js_parser.js")
		log.Println("Reached the JavaScript case of the parse switch statements")
	case "go":
		// Use `goparse` or similar to parse Go code
		cmd = exec.Command("goparse", "--ast")
	case "python":
		// Use a subprocess to parse Python code, possibly using ast in Python
		cmd = exec.Command("python", "path_to_python_parser.py")
	default:
		return nil, fmt.Errorf("unsupported language: %s", language)
	}

	return nil, nil

	cmd.Stdin = strings.NewReader(code)

	// Get the output from the subprocess
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("error running %s parser: %v", language, err)
	}

	var ast interface{}
	// Unmarshal JSON output from the subprocess into the AST structure
	if err := json.Unmarshal(output, &ast); err != nil {
		return nil, fmt.Errorf("failed to parse AST JSON: %v", err)
	}
	return ast, nil
}

// NormaliseAst standardises the AST based on the language
// func NormaliseAST(ast interface{}, language string) (types.NormalisedAST, error) {
func NormaliseAST(ast interface{}, language string) (interface{}, error) {
	// var normalised types.NormalisedAST
	println("Normalising AST", ast)
	switch language {
	case "javascript":
		return ast, nil
		// return ast_normalise.NormaliseJavaScriptAST(ast)
	// case "go":
	// 	return ast_normalise.NormaliseGoAST(ast)
	// case "python":
	// 	return ast_normalise.NormalisePythonAST(ast)
	default:
		return types.NormalisedAST{}, fmt.Errorf("unsupported language: %s", language)
	}
}
