GRAPH FinGraph
CALL PER () (p) {
  MATCH (p:Person)-[:Owns]->(a:Account)
  RETURN a
}
RETURN a
