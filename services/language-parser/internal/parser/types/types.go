package types

// NormalisedAST represents a standard structure for functions, variables, and relationships
type NormalisedAST struct {
	Functions     []Function     `json:"functions"`
	Variables     []Variable     `json:"variables"`
	Relationships []Relationship `json:"relationships"`
}

type Function struct {
	Name string `json:"name"`
	Line int    `json:"line"`
}

type Variable struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

type Relationship struct {
	Source string `json:"source"`
	Target string `json:"target"`
	Type   string `json:"type"`
	Line   int    `json:"line"`
}

type FunctionNode struct {
	Name string
	Line int
}

type CallNode struct {
	CalleeName string
	Line       int
}

type IfNode struct {
	Line     int
	HasElse  bool
	ElseLine int
}
