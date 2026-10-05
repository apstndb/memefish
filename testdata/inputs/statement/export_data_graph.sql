EXPORT DATA OPTIONS (uri = "gs://bucket/output.csv", format = "csv") AS
GRAPH FinGraph
MATCH (n:Account)
RETURN n.id
