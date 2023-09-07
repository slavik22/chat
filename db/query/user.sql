-- name: CreateUser :one
INSERT INTO users (
    name,
    login,
    hashed_password
) VALUES ($1, $2, $3) RETURNING *;

-- name: GetUsers :many
SELECT * FROM users;

-- name: GetUserByLogin :one
SELECT * FROM users
WHERE login = $1 LIMIT 1;

-- name: GetUserById :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: UpdateUser :one
UPDATE users
SET name = $2, login = $3, hashed_password = $4
WHERE id = $1
RETURNING *;

-- name: GetFriends :many
SELECT users.id, users.name, users.login FROM friends
    INNER JOIN users On friends.friend_id = users.id
                                         WHERE user_id = $1;

-- name: AddFriend :exec
INSERT INTO friends (
    user_id,
    friend_id
) VALUES ($1, $2), ($2,$1);

-- name: DeleteFriend :exec
DELETE FROM friends
       WHERE user_id=$1 AND friend_id=$2;

-- name: GetBlackList :many
SELECT users.id, users.name, users.login FROM black_list
    INNER JOIN users On black_list.friend_id = users.id
                                         WHERE user_id = $1;

-- name: GetUserFromBlackList :one
SELECT * FROM black_list
WHERE user_id = $1 AND friend_id = $2 LIMIT 1;

-- name: AddBlackList :exec
INSERT INTO black_list (
    user_id,
    friend_id
) VALUES ($1, $2);

-- name: DeleteBlackList :exec
DELETE FROM black_list
WHERE user_id=$1 AND friend_id=$2;