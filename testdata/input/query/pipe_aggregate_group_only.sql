FROM Sales
|> AGGREGATE GROUP BY category AS kind ASC, region DESC
