GRAPH FinGraph
MATCH (p:Person)
CALL graph.neighbors(p, 2) YIELD neighbor AS friend
RETURN p.name, friend
