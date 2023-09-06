package db

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomChat(t *testing.T) Chat {
	user1 := createRandomUser(t)
	user2 := createRandomUser(t)

	require.NotEmpty(t, user1)
	require.NotEmpty(t, user2)

	args := CreateChatRoomParams{
		User1ID: user1.ID,
		User2ID: user2.ID,
	}

	chat, err := testStore.CreateChatRoom(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, chat)

	require.Equal(t, chat.User1ID, args.User1ID)
	require.Equal(t, chat.User2ID, args.User2ID)

	return chat
}

func TestCreateChat(t *testing.T) {
	createRandomUser(t)
}

func TestGetChat(t *testing.T) {
	chat := createRandomChat(t)

	createdChat, err := testStore.GetChat(context.Background(), chat.ID)
	require.NoError(t, err)
	require.NotEmpty(t, createdChat)

	require.Equal(t, chat.User1ID, createdChat.User1ID)
	require.Equal(t, chat.User2ID, createdChat.User2ID)
}

func TestDeleteChat(t *testing.T) {
	chat := createRandomChat(t)

	err := testStore.DeleteChatRoom(context.Background(), chat.ID)

	require.NoError(t, err)

	_, err = testStore.GetChat(context.Background(), chat.ID)

	require.Error(t, sql.ErrNoRows)
}
