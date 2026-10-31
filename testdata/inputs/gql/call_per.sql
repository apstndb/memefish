GRAPH FinGraph
MATCH (n)-[e]->(m)
RETURN n, e, m
NEXT
CALL PER () PageRank() YIELD n, rank
RETURN n.id, rank
