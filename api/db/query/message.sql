-- name: GetChatMessages :many
SELECT messages.id,chat_room_id,user_id,content,createdAt,users.name FROM messages
INNER JOIN users On messages.user_id = users.id
WHERE chat_room_id = $1;

-- name: CreateMessage :one
INSERT INTO messages (
    chat_room_id,
    user_id,
    content
) VALUES ($1, $2, $3)
RETURNING *;

-- name: DeleteMessage :exec
DELETE FROM messages
WHERE id=$1;

-- name: UpdateMessage :one
UPDATE messages
SET content = $2
WHERE id = $1
 RETURNING *;
