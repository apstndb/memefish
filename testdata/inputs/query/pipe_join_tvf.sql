FROM LeftTable
|> JOIN ReadRightTable(1) USING(id)
