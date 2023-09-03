// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: chat_room.sql

package db

import (
	"context"
)

const createChatRoom = `-- name: CreateChatRoom :one
INSERT INTO chat_rooms (
    name
) VALUES ($1)
RETURNING id, name
`

func (q *Queries) CreateChatRoom(ctx context.Context, name string) (ChatRoom, error) {
	row := q.db.QueryRowContext(ctx, createChatRoom, name)
	var i ChatRoom
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const deleteChatRoom = `-- name: DeleteChatRoom :exec
DELETE FROM chat_rooms
WHERE id=$1
`

func (q *Queries) DeleteChatRoom(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteChatRoom, id)
	return err
}

const getUserChatRooms = `-- name: GetUserChatRooms :many
SELECT user_chat_rooms.id,user_chat_rooms.chat_room_id,chat_rooms.name FROM user_chat_rooms
INNER JOIN chat_rooms On user_chat_rooms.chat_room_id = chat_rooms.id
WHERE user_id = $1
`

type GetUserChatRoomsRow struct {
	ID         int64  `json:"id"`
	ChatRoomID int64  `json:"chat_room_id"`
	Name       string `json:"name"`
}

func (q *Queries) GetUserChatRooms(ctx context.Context, userID int64) ([]GetUserChatRoomsRow, error) {
	rows, err := q.db.QueryContext(ctx, getUserChatRooms, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetUserChatRoomsRow{}
	for rows.Next() {
		var i GetUserChatRoomsRow
		if err := rows.Scan(&i.ID, &i.ChatRoomID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}