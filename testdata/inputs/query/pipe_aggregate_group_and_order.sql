FROM Sales
|> AGGREGATE SUM(amount) AS total ASC
   GROUP AND ORDER BY category, region area DESC
