GRAPH FinGraph
MATCH (a:Account)
RETURN a.id AS id
INNER UNION ALL
MATCH (a:Account)
RETURN a.id AS id
