GRAPH FinGraph
MATCH (a:Account)
RETURN a.id AS id
FULL UNION ALL
MATCH (a:Account)
RETURN a.id AS id
