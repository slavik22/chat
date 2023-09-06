package api

import (
	"database/sql"
	"errors"
	"github.com/labstack/echo/v4"
	db "github.com/slavik22/chat/db/sqlc"
	"net/http"
	"strconv"
)

func (server *Server) getChats(ctx echo.Context) error {
	userId, err := getUserId(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	chats, err := server.store.GetUserChats(ctx.Request().Context(), userId)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return echo.NewHTTPError(http.StatusNotFound, err)
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

	}

	return ctx.JSON(http.StatusOK, chats)
}
func (server *Server) createChat(ctx echo.Context) error {
	user1Id, err := getUserId(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	user2Id, err := strconv.Atoi(ctx.Param("userId"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	arg := db.GetUserFromBlackListParams{UserID: int64(user2Id), FriendID: user1Id}

	_, err = server.store.GetUserFromBlackList(ctx.Request().Context(), arg)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			req := db.CreateChatRoomParams{
				User1ID: user1Id,
				User2ID: int64(user2Id),
			}

			err = server.store.CreateChatRoom(ctx.Request().Context(), req)

			if err != nil {
				if errors.Is(err, sql.ErrNoRows) {
					return echo.NewHTTPError(http.StatusNotFound, err)
				} else {
					return echo.NewHTTPError(http.StatusInternalServerError, err)
				}

			}

			return ctx.NoContent(http.StatusCreated)
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
	}
	return echo.NewHTTPError(http.StatusBadRequest, "User is in black list")

}
func (server *Server) deleteChat(ctx echo.Context) error {
	userId, err := getUserId(ctx)

	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	chatId, err := strconv.Atoi(ctx.Param("chatId"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	chat, err := server.store.GetChat(ctx.Request().Context(), int64(chatId))

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if chat.User1ID != userId && chat.User2ID != userId {
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
