GRAPH FinGraph
MATCH (a:Account)
RETURN a.id AS id
LEFT UNION ALL
MATCH (a:Account)
RETURN a.id AS id
