FROM Sales
|> AGGREGATE COUNT(*) AS n,
|> WHERE n > 0
