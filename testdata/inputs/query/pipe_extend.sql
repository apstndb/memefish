SELECT * FROM (
  SELECT * FROM Singers
  |> EXTEND FirstName, LastName AS family_name, FirstName given_name
  |> EXTEND LENGTH(family_name) AS family_name_length
  |> EXTEND Singers.* EXCEPT (SingerId) REPLACE (UPPER(FirstName) AS FirstName),
)
