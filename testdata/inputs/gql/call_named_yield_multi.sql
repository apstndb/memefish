GRAPH FinGraph
MATCH (p:Person)
CALL my_tvf(p) YIELD Id AS tvf_Id, score
RETURN p.Id AS person_Id, tvf_Id, score
