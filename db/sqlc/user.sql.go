// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: user.sql

package db

import (
	"context"
	"database/sql"
)

const addBlackList = `-- name: AddBlackList :exec
INSERT INTO black_list (
    user_id,
    friend_id
) VALUES ($1, $2)
`

type AddBlackListParams struct {
	UserID   int64 `json:"user_id"`
	FriendID int64 `json:"friend_id"`
}

func (q *Queries) AddBlackList(ctx context.Context, arg AddBlackListParams) error {
	_, err := q.db.ExecContext(ctx, addBlackList, arg.UserID, arg.FriendID)
	return err
}

const addFriend = `-- name: AddFriend :exec
INSERT INTO friends (
    user_id,
    friend_id
) VALUES ($1, $2), ($2,$1)
`

type AddFriendParams struct {
	UserID   int64 `json:"user_id"`
	FriendID int64 `json:"friend_id"`
}

func (q *Queries) AddFriend(ctx context.Context, arg AddFriendParams) error {
	_, err := q.db.ExecContext(ctx, addFriend, arg.UserID, arg.FriendID)
	return err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (
    name,
    login,
    hashed_password
) VALUES ($1, $2, $3) RETURNING id, name, login, hashed_password, image_name
`

type CreateUserParams struct {
	Name           string `json:"name"`
	Login          string `json:"login"`
	HashedPassword string `json:"hashed_password"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Name, arg.Login, arg.HashedPassword)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Login,
		&i.HashedPassword,
		&i.ImageName,
	)
	return i, err
}

const deleteBlackList = `-- name: DeleteBlackList :exec
DELETE FROM black_list
WHERE user_id=$1 AND friend_id=$2
`

type DeleteBlackListParams struct {
	UserID   int64 `json:"user_id"`
	FriendID int64 `json:"friend_id"`
}

func (q *Queries) DeleteBlackList(ctx context.Context, arg DeleteBlackListParams) error {
	_, err := q.db.ExecContext(ctx, deleteBlackList, arg.UserID, arg.FriendID)
	return err
}

const deleteFriend = `-- name: DeleteFriend :exec
DELETE FROM friends
       WHERE user_id=$1 AND friend_id=$2
`

type DeleteFriendParams struct {
	UserID   int64 `json:"user_id"`
	FriendID int64 `json:"friend_id"`
}

func (q *Queries) DeleteFriend(ctx context.Context, arg DeleteFriendParams) error {
	_, err := q.db.ExecContext(ctx, deleteFriend, arg.UserID, arg.FriendID)
	return err
}

const getBlackList = `-- name: GetBlackList :many
SELECT users.id, users.name, users.login FROM black_list
    INNER JOIN users On black_list.friend_id = users.id
                                         WHERE user_id = $1
`

type GetBlackListRow struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Login string `json:"login"`
}

func (q *Queries) GetBlackList(ctx context.Context, userID int64) ([]GetBlackListRow, error) {
	rows, err := q.db.QueryContext(ctx, getBlackList, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetBlackListRow{}
	for rows.Next() {
		var i GetBlackListRow
		if err := rows.Scan(&i.ID, &i.Name, &i.Login); err != nil {
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

const getFriends = `-- name: GetFriends :many
SELECT users.id, users.name, users.login FROM friends
    INNER JOIN users On friends.friend_id = users.id
                                         WHERE user_id = $1
`

type GetFriendsRow struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Login string `json:"login"`
}

func (q *Queries) GetFriends(ctx context.Context, userID int64) ([]GetFriendsRow, error) {
	rows, err := q.db.QueryContext(ctx, getFriends, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetFriendsRow{}
	for rows.Next() {
		var i GetFriendsRow
		if err := rows.Scan(&i.ID, &i.Name, &i.Login); err != nil {
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

const getImageName = `-- name: GetImageName :one
SELECT image_name
From users
WHERE id = $1
`

func (q *Queries) GetImageName(ctx context.Context, id int64) (sql.NullString, error) {
	row := q.db.QueryRowContext(ctx, getImageName, id)
	var image_name sql.NullString
	err := row.Scan(&image_name)
	return image_name, err
}

const getUserById = `-- name: GetUserById :one
SELECT id, name, login, hashed_password, image_name FROM users
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUserById(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Login,
		&i.HashedPassword,
		&i.ImageName,
	)
	return i, err
}

const getUserByLogin = `-- name: GetUserByLogin :one
SELECT id, name, login, hashed_password, image_name FROM users
WHERE login = $1 LIMIT 1
`

func (q *Queries) GetUserByLogin(ctx context.Context, login string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByLogin, login)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Login,
		&i.HashedPassword,
		&i.ImageName,
	)
	return i, err
}

const getUserFromBlackList = `-- name: GetUserFromBlackList :one
SELECT user_id, friend_id, created_at FROM black_list
WHERE user_id = $1 AND friend_id = $2 LIMIT 1
`

type GetUserFromBlackListParams struct {
	UserID   int64 `json:"user_id"`
	FriendID int64 `json:"friend_id"`
}

func (q *Queries) GetUserFromBlackList(ctx context.Context, arg GetUserFromBlackListParams) (BlackList, error) {
	row := q.db.QueryRowContext(ctx, getUserFromBlackList, arg.UserID, arg.FriendID)
	var i BlackList
	err := row.Scan(&i.UserID, &i.FriendID, &i.CreatedAt)
	return i, err
}

const getUsers = `-- name: GetUsers :many
SELECT id, name, login, hashed_password, image_name FROM users
`

func (q *Queries) GetUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Login,
			&i.HashedPassword,
			&i.ImageName,
		); err != nil {
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

const updateImageName = `-- name: UpdateImageName :one
UPDATE users
SET image_name = $2
WHERE id = $1
RETURNING id, name, login, hashed_password, image_name
`

type UpdateImageNameParams struct {
	ID        int64          `json:"id"`
	ImageName sql.NullString `json:"image_name"`
}

func (q *Queries) UpdateImageName(ctx context.Context, arg UpdateImageNameParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateImageName, arg.ID, arg.ImageName)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Login,
		&i.HashedPassword,
		&i.ImageName,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET name = $2, login = $3, hashed_password = $4
WHERE id = $1
RETURNING id, name, login, hashed_password, image_name
`

type UpdateUserParams struct {
	ID             int64  `json:"id"`
	Name           string `json:"name"`
	Login          string `json:"login"`
	HashedPassword string `json:"hashed_password"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.ID,
		arg.Name,
		arg.Login,
		arg.HashedPassword,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Login,
		&i.HashedPassword,
		&i.ImageName,
	)
	return i, err
}
