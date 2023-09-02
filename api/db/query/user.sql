-- name: CreateUser :one
INSERT INTO users (
    name,
    login,
    hashed_password
) VALUES ($1, $2, $3) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE login = $1 LIMIT 1;

-- name: UpdateUser :one
UPDATE users
SET name = $2, login = $3, hashed_password = $4
WHERE id = $1
RETURNING *;