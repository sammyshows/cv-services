package neo4j

import (
	"log"

	"github.com/neo4j/neo4j-go-driver/neo4j"
)

func RunCypherQueries(queries []string) {
	driver, err := neo4j.NewDriver("bolt://localhost:7687", neo4j.NoAuth())
	if err != nil {
		log.Fatal(err)
	}
	defer driver.Close()

	session, err := driver.Session(neo4j.AccessModeWrite)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	for _, query := range queries {
		_, err := session.Run(query, nil)
		if err != nil {
			log.Println("Error running query:", query, err)
		}
	}
}

func GetFunctions() ([]map[string]interface{}, error) {
	driver, err := neo4j.NewDriver("bolt://localhost:7687", neo4j.NoAuth())
	if err != nil {
		return nil, err
	}
	defer driver.Close()

	session, err := driver.Session(neo4j.AccessModeRead)
	if err != nil {
		return nil, err
	}
	defer session.Close()

	query := `
        MATCH (f:Function)
        OPTIONAL MATCH (f)-[call:CALLS]->(called:Function)
        OPTIONAL MATCH (f)-[cond:CONDITIONAL]->(cf:Function)
        RETURN f AS Function, 
               collect(DISTINCT {type: call.type, callee: called.name, line: call.line, order: call.order}) AS Calls,
               collect(DISTINCT {type: cond.type, condition: cf.name, line: cond.line}) AS Conditionals
        ORDER BY f.line`

	result, err := session.Run(query, nil)
	if err != nil {
		return nil, err
	}

	var functions []map[string]interface{}
	for result.Next() {
		functions = append(functions, result.Record().Values)
	}
	return functions, nil
}
