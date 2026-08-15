FROM LeftTable AS l
|> INNER JOIN FirstTable AS r1 ON l.id = r1.id
|> RIGHT JOIN SecondTable AS r2 ON l.id = r2.id
|> FULL OUTER JOIN ThirdTable AS r3 ON l.id = r3.id
