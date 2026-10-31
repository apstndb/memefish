FROM LeftTable AS l
|> JOIN (SELECT 1 AS id) AS r ON l.id = r.id
