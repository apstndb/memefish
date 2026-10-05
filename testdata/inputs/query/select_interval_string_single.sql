SELECT
  INTERVAL '1' DAY AS day_value,
  INTERVAL '-2' MONTH AS month_value,
  INTERVAL '1.5' SECOND AS second_value,
  INTERVAL 1 DAY AS integer_value,
  INTERVAL @days DAY AS parameter_value,
  INTERVAL '1-2' YEAR TO MONTH AS range_value
