package db

import (
	"context"
	"github.com/slavik22/chat/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomMessage(t *testing.T) Message {
	user := createRandomUser(t)
	chat := createRandomChat(t)

	require.NotEmpty(t, user)

	args := CreateMessageParams{
		ChatID:  chat.ID,
		UserID:  user.ID,
		Content: util.RandomString(10),
	}

	message, err := testStore.CreateMessage(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, message)

	require.Equal(t, args.ChatID, message.ChatID)
	require.Equal(t, args.UserID, message.UserID)
	require.Equal(t, args.Content, message.Content)

	return message
}

func TestCreateMessage(t *testing.T) {
	createRandomMessage(t)
}

func TestUpdateMessage(t *testing.T) {
	content := util.RandomString(6)
	msg := createRandomMessage(t)

	args := UpdateMessageParams{ID: msg.ID, Content: content}
	updatedMsg, err := testStore.UpdateMessage(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, updatedMsg)

	require.Equal(t, args.Content, updatedMsg.Content)
}

func TestDeleteMessage(t *testing.T) {
	msg := createRandomMessage(t)

	err := testStore.DeleteMessage(context.Background(), msg.ID)

	require.NoError(t, err)
}
