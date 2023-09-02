package api

import (
	"database/sql"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
	db "github.com/slavik22/chat/db/sqlc"
	"github.com/slavik22/chat/util"
	"net/http"
	"strconv"
)

type userLoginRequest struct {
	Login    string `json:"login" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

type loginUserResponse struct {
	AccessToken string `json:"access_token"`
}

func (server *Server) loginUser(ctx echo.Context) error {
	var req userLoginRequest
	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	user, err := server.store.GetUserByLogin(ctx.Request().Context(), req.Login)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return echo.NewHTTPError(http.StatusNotFound, err)
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

	}

	err = util.CheckPassword(req.Password, user.HashedPassword)

	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	accessToken, _, err := server.jwtMaker.CreateToken(
		user.ID,
		server.config.AccessTokenDuration,
	)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	rsp := loginUserResponse{
		AccessToken: accessToken,
	}

	return ctx.JSON(http.StatusOK, rsp)
}

func (server *Server) getUsers(ctx echo.Context) error {
	users, err := server.store.GetUsers(ctx.Request().Context())

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, users)
}

func (server *Server) getUser(ctx echo.Context) error {
	userId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	users, err := server.store.GetUserById(ctx.Request().Context(), int64(userId))

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, users)
}

type userRequest struct {
	Name     string `json:"name" binding:"required,alphanum"`
	Login    string `json:"login" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

type userResponse struct {
	Name  string `json:"name"`
	Login string `json:"login"`
}

func (server *Server) createUser(ctx echo.Context) error {
	var requestData userRequest

	if err := ctx.Bind(&requestData); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	hashedPassword, err := util.HashPassword(requestData.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	arg := db.CreateUserParams{
		Name:           requestData.Name,
		Login:          requestData.Login,
		HashedPassword: hashedPassword,
	}

	user, err := server.store.CreateUser(ctx.Request().Context(), arg)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				return echo.NewHTTPError(http.StatusForbidden, err)
			}
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusCreated, userResponse{Name: user.Name, Login: user.Login})
}
func (server *Server) updateUser(ctx echo.Context) error {
	var requestData userRequest

	if err := ctx.Bind(&requestData); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	hashedPassword, err := util.HashPassword(requestData.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	id, err := getUserId(ctx)

	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	arg := db.UpdateUserParams{
		ID:             int64(id),
		Name:           requestData.Name,
		Login:          requestData.Login,
		HashedPassword: hashedPassword,
	}

	user, err := server.store.UpdateUser(ctx.Request().Context(), arg)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				return echo.NewHTTPError(http.StatusForbidden, err)
			}
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, userResponse{Name: user.Name, Login: user.Login})

}
