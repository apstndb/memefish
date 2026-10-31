GRAPH FinGraph
MATCH (p:Person)
CALL (p) {
  GRAPH OtherGraph
  MATCH (p)-[:Owns]->(a:Account)
  RETURN a.Id AS account_Id
}
RETURN p.name, account_Id
