package api

import (
	"database/sql"
	"errors"
	"github.com/labstack/echo/v4"
	db "github.com/slavik22/chat/db/sqlc"
	"net/http"
	"strconv"
)

type createChatRequest struct {
	Name string `json:"name" binding:"required,alphanum"`
}

func (server *Server) getUserChatRooms(ctx echo.Context) error {
	userId, err := getUserId(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	rooms, err := server.store.GetUserChatRooms(ctx.Request().Context(), userId)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return echo.NewHTTPError(http.StatusNotFound, err)
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

	}

	return ctx.JSON(http.StatusOK, rooms)
}

func (server *Server) createChatRoom(ctx echo.Context) error {
	userId, err := getUserId(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	var req createChatRequest

	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	chat, err := server.store.CreateChatRoom(ctx.Request().Context(), req.Name)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return echo.NewHTTPError(http.StatusNotFound, err)
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

	}

	arg := db.AddUserToChatParams{UserID: userId, ChatRoomID: chat.ID}

	err = server.store.AddUserToChat(ctx.Request().Context(), arg)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return echo.NewHTTPError(http.StatusNotFound, err)
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
	}

	return ctx.NoContent(http.StatusCreated)
}
func (server *Server) deleteChatRoom(ctx echo.Context) error {
	userId, err := getUserId(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	chatId, err := strconv.Atoi(ctx.Param("chatId"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	isMember, err := server.isMemberOfChatRoom(ctx, userId, int64(chatId))

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if !isMember {
		return echo.NewHTTPError(http.StatusUnauthorized, "User is not member of chat room")
	}

	err = server.store.DeleteChatRoom(ctx.Request().Context(), int64(chatId))

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return echo.NewHTTPError(http.StatusNotFound, err)
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

	}

	return ctx.NoContent(http.StatusOK)
}

func (server *Server) addUserToChatRoom(ctx echo.Context) error {
	userId, err := getUserId(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	chatId, err := strconv.Atoi(ctx.Param("chatId"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	userToAddId, err := strconv.Atoi(ctx.Param("userId"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	isMember, err := server.isMemberOfChatRoom(ctx, userId, int64(chatId))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	if !isMember {
		return echo.NewHTTPError(http.StatusUnauthorized, "User is not member of chat room")
	}

	arg := db.AddUserToChatParams{
		ChatRoomID: int64(chatId),
		UserID:     int64(userToAddId),
	}

	err = server.store.AddUserToChat(ctx.Request().Context(), arg)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return echo.NewHTTPError(http.StatusNotFound, err)
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

	}

	return ctx.NoContent(http.StatusCreated)
}
func (server *Server) removeUserFromChatRoom(ctx echo.Context) error {
	userId, err := getUserId(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	chatId, err := strconv.Atoi(ctx.Param("chatId"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	userToAddId, err := strconv.Atoi(ctx.Param("userId"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	isMember, err := server.isMemberOfChatRoom(ctx, userId, int64(chatId))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	if !isMember {
		return echo.NewHTTPError(http.StatusUnauthorized, "User is not member of chat room")
	}

	arg := db.DeleteUserFromChatParams{
		ChatRoomID: int64(chatId),
		UserID:     int64(userToAddId),
	}

	err = server.store.DeleteUserFromChat(ctx.Request().Context(), arg)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return echo.NewHTTPError(http.StatusNotFound, err)
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

	}

	return ctx.NoContent(http.StatusCreated)
}

func (server *Server) isMemberOfChatRoom(ctx echo.Context, userId int64, chatId int64) (bool, error) {
	rooms, err := server.store.GetUserChatRooms(ctx.Request().Context(), userId)
	if err != nil {
		return false, err
	}

	for _, room := range rooms {
		if room.ChatRoomID == int64(chatId) {
			return true, nil
		}
	}

	return false, nil
}
