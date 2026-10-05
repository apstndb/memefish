(
  FROM UNNEST([1]) AS x
  |> AGGREGATE COUNT(*) AS n,
)
