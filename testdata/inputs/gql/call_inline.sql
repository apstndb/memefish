GRAPH FinGraph
MATCH (p:Person)-[:Owns]->(a:Account)
CALL (p, a) {
  MATCH (p)-[:Owns]->(a:Account)
  RETURN a.Id AS account_Id
  ORDER BY account_Id
  LIMIT 2
}
RETURN p.name AS person_name, a.Id, account_Id
