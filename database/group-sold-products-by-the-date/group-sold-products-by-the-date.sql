WITH a AS (
    SELECT DISTINCT * FROM Activities)

SELECT
    sell_date
     ,COUNT(product) AS num_sold
     ,STRING_AGG(product,',') WITHIN group (ORDER BY product) AS products
FROM a
GROUP BY sell_date