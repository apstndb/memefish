GRAPH FinGraph
MATCH (p:Person)
OPTIONAL CALL graph.neighbors(p)
RETURN p.name
