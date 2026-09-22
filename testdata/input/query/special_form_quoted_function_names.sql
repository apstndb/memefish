SELECT `SAFE_CAST`(1), `safe_cast`(2), `Safe_Cast`(3),
       `REPLACE_FIELDS`(p), `replace_fields`(p), `Replace_Fields`(p),
       schema.`SAFE_CAST`(1), schema.`REPLACE_FIELDS`(p), `ordinary`(1),
       SAFE_CAST("1" AS INT64), REPLACE_FIELDS(p, 1 AS x)
FROM T
