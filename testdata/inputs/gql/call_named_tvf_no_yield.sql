GRAPH FinGraph
MATCH (p:Person)
CALL graph.neighbors(p)
RETURN p.name
