package db

import (
	"context"
	"github.com/slavik22/chat/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomUser(t *testing.T) User {
	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)

	arg := CreateUserParams{
		Name:           util.RandomString(6),
		Login:          util.RandomOwner(),
		HashedPassword: hashedPassword,
	}

	user, err := testStore.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Name, user.Name)
	require.Equal(t, arg.Login, user.Login)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testStore.GetUser(context.Background(), user1.Login)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Name, user2.Name)
	require.Equal(t, user1.Login, user2.Login)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
}

func TestUpdateUser(t *testing.T) {
	name := util.RandomString(6)
	login := util.RandomString(6)
	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)

	user := createRandomUser(t)
	userUpdated, err := testStore.UpdateUser(context.Background(),
		UpdateUserParams{ID: user.ID, Name: name, Login: login, HashedPassword: hashedPassword})

	require.NoError(t, err)
	require.NotEmpty(t, userUpdated)

	require.Equal(t, userUpdated.Name, name)
	require.Equal(t, userUpdated.Login, login)
	require.Equal(t, userUpdated.HashedPassword, hashedPassword)
}
