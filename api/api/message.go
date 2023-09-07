package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	db "github.com/slavik22/chat/db/sqlc"
	"golang.org/x/net/websocket"
	"net/http"
	"strconv"
	"sync"
)

type Client struct {
	UserId int
	Conn   *websocket.Conn
}

type Message struct {
	UserID  int    `json:"userId"`
	Message string `json:"message"`
}

type ChatRoom struct {
	id        int
	clients   map[*Client]bool
	broadcast chan Message
	mu        sync.Mutex
}

var (
	rooms map[int]*ChatRoom
)

func createChatRoom(id int) *ChatRoom {
	return &ChatRoom{
		id:        id,
		clients:   make(map[*Client]bool),
		broadcast: make(chan Message),
	}
}

func (server *Server) GetChatMessages(ctx echo.Context) error {
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

	if chat.User1ID != int64(userId) && chat.User2ID != int64(userId) {
		return echo.NewHTTPError(http.StatusUnauthorized, "User is not member of chat room")
	}

	messages, err := server.store.GetChatMessages(ctx.Request().Context(), int64(chatId))

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return echo.NewHTTPError(http.StatusNotFound, err)
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

	}

	return ctx.JSON(http.StatusOK, messages)
}

func (server *Server) addMessage(ctx echo.Context, req db.CreateMessageParams) (*db.Message, error) {
	msg, err := server.store.CreateMessage(ctx.Request().Context(), req)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, echo.NewHTTPError(http.StatusNotFound, err)
		} else {
			return nil, echo.NewHTTPError(http.StatusInternalServerError, err)
		}
	}

	return &msg, nil
}

func (c *ChatRoom) handleMessages(server *Server, ctx echo.Context) {
	for {
		msg := <-c.broadcast
		msgAdded, err := server.addMessage(ctx, db.CreateMessageParams{UserID: int64(msg.UserID), ChatID: int64(c.id), Content: msg.Message})

		if err != nil {
			_ = fmt.Errorf("Cannot create message %v\n", err)
			return
		}
		user, err := server.store.GetUserById(ctx.Request().Context(), int64(msg.UserID))

		if err != nil {
			_ = fmt.Errorf("Cannot get user %v\n", err)
			return
		}

		m := struct {
			Id        int64        `json:"id"`
			UserId    int64        `json:"userId"`
			Name      string       `json:"name"`
			CreatedAt sql.NullTime `json:"createdAt"`
			Content   string       `json:"content"`
		}{
			Id:        msgAdded.ID,
			UserId:    msgAdded.UserID,
			Name:      user.Name,
			CreatedAt: msgAdded.Createdat,
			Content:   msgAdded.Content,
		}

		for client := range c.clients {
			marshal, err := json.Marshal(m)
			s := string(marshal)

			if err != nil {
				_ = fmt.Errorf("cannot marshal json %v\n", err)
				return
			}

			err = websocket.Message.Send(client.Conn, s)

			if err != nil {
				fmt.Printf("Error sending message: %v\n", err)
				delete(c.clients, client)
				client.Conn.Close()
			}

		}
	}
}

func (c *ChatRoom) addClient(client *Client) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.clients[client] = true
}

func (c *ChatRoom) removeClient(client *Client) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.clients, client)
}

func (server *Server) webSocketConn(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("userId"))

	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	chatId, err := strconv.Atoi(c.Param("chatId"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	chat, err := server.store.GetChat(c.Request().Context(), int64(chatId))

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if chat.User1ID != int64(userId) && chat.User2ID != int64(userId) {
		return echo.NewHTTPError(http.StatusUnauthorized, "User is not member of chat room")
	}

	websocket.Handler(func(ws *websocket.Conn) {
		client := &Client{Conn: ws, UserId: userId}

		if rooms == nil {
			rooms = make(map[int]*ChatRoom)
		}
		room, exists := rooms[chatId]

		if !exists {
			room = createChatRoom(chatId)
			rooms[chatId] = room
			go room.handleMessages(server, c)
		}

		room.addClient(client)

		defer func() {
			room.removeClient(client)
			ws.Close()
		}()

		for k, _ := range rooms[chatId].clients {
			fmt.Printf("Client: userId= %v, valur = %v\n", k.UserId, rooms[chatId].clients)

		}

		for {
			var content string
			if err := websocket.Message.Receive(ws, &content); err != nil {
				fmt.Printf("Error receiving message: %v\n", err)
				break
			}

			msg := Message{
				UserID:  userId,
				Message: content,
			}

			rooms[chatId].broadcast <- msg
		}

	}).ServeHTTP(c.Response(), c.Request())

	if len(rooms[chatId].clients) == 0 {
		delete(rooms, chatId)
	}

	return nil
}
