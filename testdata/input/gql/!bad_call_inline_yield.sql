GRAPH FinGraph
CALL (p) {
  MATCH (p:Person)-[:Owns]->(a:Account)
  RETURN a
} YIELD a
RETURN a
