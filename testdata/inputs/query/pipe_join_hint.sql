FROM LeftTable
|> JOIN@{JOIN_TYPE=HASH_JOIN} RightTable USING(id)
