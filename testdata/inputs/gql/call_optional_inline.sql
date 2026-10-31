GRAPH FinGraph
MATCH (p:Person)
OPTIONAL CALL (p) {
  MATCH (p)-[:Owns]->(a:Account)
  RETURN a.Id AS account_Id
  LIMIT 1
}
RETURN p.name, account_Id
