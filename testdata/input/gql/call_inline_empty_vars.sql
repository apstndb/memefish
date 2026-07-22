GRAPH FinGraph
MATCH (p:Person)
CALL () {
  MATCH (x:Person)
  RETURN COUNT(*) AS total_persons
}
RETURN p.name, total_persons
