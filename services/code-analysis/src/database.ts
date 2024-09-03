const neo4j = require('neo4j-driver')

// Production (Neo4j Aura)
// const uri = "neo4j+s://04d662a4.databases.neo4j.io:7687";
// const user = "neo4j"; // Replace with your actual username
// const password = "K2RAQlt9Vli93lpOSGjUXInC1_5Eky7DhjoEnH99CEw"; // Replace with your actual password
// const driver = neo4j.driver(uri, neo4j.auth.basic(user, password));
// const session = driver.session();

// Local
const uri = "bolt://neo4j:7687";
const neo4jDriver = neo4j.driver(uri); // No auth needed since authentication is disabled
const getNeo4jSession =  () => neo4jDriver.session();

export { neo4jDriver, getNeo4jSession };