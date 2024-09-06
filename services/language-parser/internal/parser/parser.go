package parser

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func ParseContent(code string, language string) (interface{}, error) {
	var cmd *exec.Cmd

	// Switch logic to decide the parser based on the language
	switch language {
	case "javaScript":
		// Call JavaScript parser subprocess
		cmd = exec.Command("node", "internal/parser/ast_scripts/javascript/parser.js")
	case "go":
		// Use `goparse` or similar to parse Go code
		cmd = exec.Command("goparse", "--ast")
	case "python":
		// Use a subprocess to parse Python code, possibly using ast in Python
		cmd = exec.Command("python", "internal/parser/ast_scripts/python/parser.py")
	default:
		return nil, fmt.Errorf("unsupported language: %s", language)
	}

	// Pass the code to the parser via stdin
	cmd.Stdin = strings.NewReader(code)

	// Capture the output from the parser
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	// Run the command
	err := cmd.Run()
	if err != nil {
		log.Println("Error running parser:", stderr.String())
		return nil, err
	}

	// Return the output as a string
	return out.String(), nil
}
