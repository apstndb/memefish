FROM LeftTable
|> CROSS JOIN UNNEST([1, 2]) AS value
