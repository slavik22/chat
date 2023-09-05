-- name: GetChatUsers :many
SELECT users.id,users.name, users.login FROM user_chat_rooms
INNER JOIN users ON user_chat_rooms.user_id = users.id
WHERE chat_room_id = $1;

-- name: AddUserToChat :exec
INSERT INTO user_chat_rooms (
    chat_room_id,
    user_id
) VALUES ($1, $2);

-- name: DeleteUserFromChat :exec
DELETE FROM user_chat_rooms
WHERE chat_room_id=$1 AND user_id=$2;
