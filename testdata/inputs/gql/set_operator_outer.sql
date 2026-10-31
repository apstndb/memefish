GRAPH FinGraph
MATCH (a:Account)
RETURN a.id AS id
OUTER UNION ALL
MATCH (a:Account)
RETURN a.id AS id
