-- name: AddUserToChat :exec
INSERT INTO user_chat_rooms (
    chat_room_id,
    user_id
) VALUES ($1, $2);

-- name: DeleteUserFromChat :exec
DELETE FROM user_chat_rooms
WHERE chat_room_id=$1 AND user_id=$2;
