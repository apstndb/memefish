FROM Sales
|> AGGREGATE COUNT(*) AS n, SUM(amount) total_sales DESC
