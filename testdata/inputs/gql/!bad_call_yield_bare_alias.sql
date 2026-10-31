GRAPH FinGraph
MATCH (p:Person)
CALL my_tvf(p) YIELD Id tvf_Id
RETURN p.Id, tvf_Id
