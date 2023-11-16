-- Inserting 50 Transactions
WITH random_category AS (
  SELECT id FROM category ORDER BY RANDOM() LIMIT 50
)
INSERT INTO transaction (wallet_id_from, wallet_id_to, amount, created_at, category_id)
SELECT
  CASE WHEN mod(id, 3) = 0 THEN 1 END AS wallet_id_from,
  CASE WHEN mod(id, 3) = 1 THEN 2 END AS wallet_id_to,
  COALESCE(
    CASE WHEN mod(id, 3) = 0 THEN (RANDOM() * 989000 + 10000)::numeric(10,2)  -- Spending
         WHEN mod(id, 3) = 1 THEN (RANDOM() * 989000 + 10000)::numeric(10,2)   -- Income
         WHEN mod(id, 3) = 2 THEN (RANDOM() * 989000 + 10000)::numeric(10,2)   -- Transfer
    END,
    0.00
  ) AS amount,
  NOW() - (id * INTERVAL '1 day') AS created_at,
  (SELECT id FROM random_category LIMIT 1) AS category_id
FROM generate_series(1, 50) id;
