SELECT * FROM (
  SELECT * FROM Singers
  |> RENAME FirstName AS given_name, LastName family_name,
)
