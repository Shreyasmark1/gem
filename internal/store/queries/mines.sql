-- name: CreateMine :one
INSERT INTO mines (
  owner_id,
  mine_name,
  created_by
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateMine :one
UPDATE mines
SET 
  mine_name = $2,
  updated_by = $3,
  updated_at = CURRENT_TIMESTAMP
WHERE mine_id = $1
RETURNING *;

-- name: GetMine :one
SELECT * FROM mines
WHERE mine_id = $1 LIMIT 1;
