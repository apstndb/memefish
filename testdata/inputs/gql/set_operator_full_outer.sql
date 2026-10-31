GRAPH FinGraph
MATCH (a:Account)
RETURN a.id AS id
FULL OUTER UNION ALL
MATCH (a:Account)
RETURN a.id AS id
