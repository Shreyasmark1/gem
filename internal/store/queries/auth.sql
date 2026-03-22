-- name: CreateAuthToken :one
INSERT INTO auth (email_id, token, expires_at)
VALUES ($1, $2, $3)
RETURNING *;

-- name: FindAuthToken :one
SELECT * FROM auth
WHERE token = $1 AND email_id = $2 AND is_expired = FALSE AND expires_at > NOW();

-- name: MarkAuthTokenExpired :exec
UPDATE auth
SET is_expired = TRUE
WHERE token = $1;