-- name: GetUserChatRooms :many
SELECT user_chat_rooms.id,user_chat_rooms.chat_room_id,chat_rooms.name FROM user_chat_rooms
INNER JOIN chat_rooms On user_chat_rooms.chat_room_id = chat_rooms.id
WHERE user_id = $1;

-- name: CreateChatRoom :one
INSERT INTO chat_rooms (
    name
) VALUES ($1)
RETURNING *;

-- name: DeleteChatRoom :exec
DELETE FROM chat_rooms
WHERE id=$1;