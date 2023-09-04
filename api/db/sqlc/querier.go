// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package db

import (
	"context"
)

type Querier interface {
	AddBlackList(ctx context.Context, arg AddBlackListParams) error
	AddFriend(ctx context.Context, arg AddFriendParams) error
	AddUserToChat(ctx context.Context, arg AddUserToChatParams) error
	CreateChatRoom(ctx context.Context, name string) (ChatRoom, error)
	CreateMessage(ctx context.Context, arg CreateMessageParams) (Message, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteBlackList(ctx context.Context, arg DeleteBlackListParams) error
	DeleteChatRoom(ctx context.Context, id int64) error
	DeleteFriend(ctx context.Context, arg DeleteFriendParams) error
	DeleteMessage(ctx context.Context, id int64) error
	DeleteUserFromChat(ctx context.Context, arg DeleteUserFromChatParams) error
	GetBlackList(ctx context.Context, userID int64) ([]GetBlackListRow, error)
	GetChatMessages(ctx context.Context, chatRoomID int64) ([]GetChatMessagesRow, error)
	GetFriends(ctx context.Context, userID int64) ([]GetFriendsRow, error)
	GetUserById(ctx context.Context, id int64) (User, error)
	GetUserByLogin(ctx context.Context, login string) (User, error)
	GetUserChatRooms(ctx context.Context, userID int64) ([]GetUserChatRoomsRow, error)
	GetUsers(ctx context.Context) ([]User, error)
	UpdateMessage(ctx context.Context, arg UpdateMessageParams) (Message, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
}

var _ Querier = (*Queries)(nil)
