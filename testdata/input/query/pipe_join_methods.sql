FROM LeftTable AS l
|> HASH JOIN RightTable AS r ON l.id = r.id
|> LOOKUP JOIN ThirdTable AS t USING(id)
