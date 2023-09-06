-- name: GetChat :one
SELECT * FROM chats
WHERE id = $1;

-- name: GetUserChats :many
SELECT c.id, user1.name as name1, user2.name as name2
FROM chats c
JOIN users user1 ON c.user1_id = user1.id
JOIN users user2 ON c.user2_id = user2.id
WHERE c.user1_id = $1 OR c.user2_id = $1;

-- name: CreateChatRoom :exec
INSERT INTO chats (
    user1_id,
    user2_id
) VALUES ($1,$2);

-- name: DeleteChatRoom :exec
DELETE FROM chats
WHERE id=$1;