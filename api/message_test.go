package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	mockdb "github.com/slavik22/chat/db/mock"
	db "github.com/slavik22/chat/db/sqlc"
	"github.com/slavik22/chat/token"
	"github.com/slavik22/chat/util"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"
)

//	func TestCreateMessageAPI(t *testing.T) {
//		user, _ := randomUser(t)
//		user2, _ := randomUser(t)
//
//		chat := randomChat(user.ID, user2.ID)
//
//		message := randomMessage(chat.ID, user.ID)
//
//		testCases := []struct {
//			name          string
//			user2Id       int64
//			setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.JWTMaker)
//			buildStubs    func(store *mockdb.MockStore)
//			checkResponse func(recoder *httptest.ResponseRecorder)
//		}{
//			{
//				name:    "OK",
//				user2Id: user2.ID,
//				setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.JWTMaker) {
//					addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, time.Minute)
//				},
//				buildStubs: func(store *mockdb.MockStore) {
//
//					arg := db.CreateMessageParams{
//						ChatID:  message.ChatID,
//						UserID:  message.UserID,
//						Content: message.Content,
//					}
//					store.EXPECT().
//						CreateMessage(gomock.Any(), gomock.Eq(arg)).
//						Times(1).
//						Return(message, nil)
//
//				},
//				checkResponse: func(recorder *httptest.ResponseRecorder) {
//					require.Equal(t, http.StatusCreated, recorder.Code)
//				},
//			},
//			{
//				name:    "NoAuthorization",
//				user2Id: user2.ID,
//				setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.JWTMaker) {
//				},
//				buildStubs: func(store *mockdb.MockStore) {
//					store.EXPECT().
//						CreateMessage(gomock.Any(), gomock.Any()).
//						Times(0)
//				},
//				checkResponse: func(recorder *httptest.ResponseRecorder) {
//					require.Equal(t, http.StatusUnauthorized, recorder.Code)
//				},
//			},
//			{
//				name:    "InternalError",
//				user2Id: user2.ID,
//				setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.JWTMaker) {
//					addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, time.Minute)
//				},
//				buildStubs: func(store *mockdb.MockStore) {
//					store.EXPECT().
//						CreateMessage(gomock.Any(), gomock.Any()).
//						Times(1).
//						Return(db.Message{}, sql.ErrConnDone)
//
//					args2 := db.GetUserFromBlackListParams{UserID: user2.ID, FriendID: user.ID}
//
//					store.EXPECT().
//						GetUserFromBlackList(gomock.Any(), gomock.Eq(args2)).
//						Times(1).
//						Return(db.BlackList{}, sql.ErrNoRows)
//				},
//				checkResponse: func(recorder *httptest.ResponseRecorder) {
//					require.Equal(t, http.StatusInternalServerError, recorder.Code)
//				},
//			},
//		}
//
//		for i := range testCases {
//			tc := testCases[i]
//
//			t.Run(tc.name, func(t *testing.T) {
//				ctrl := gomock.NewController(t)
//				defer ctrl.Finish()
//
//				store := mockdb.NewMockStore(ctrl)
//				tc.buildStubs(store)
//
//				server := newTestServer(t, store)
//				recorder := httptest.NewRecorder()
//
//				websocket.NewClient()
//
//				err = websocket.Message.Send(client.Conn, s)
//
//				url := "/chatroom/" + strconv.FormatInt(chat.ID, 10) + "/user/" + strconv.FormatInt(user.ID, 10) + "/"
//				request := httptest.NewRequest(http.MethodPost, url, nil)
//
//				tc.setupAuth(t, request, server.jwtMaker)
//				server.router.ServeHTTP(recorder, request)
//				tc.checkResponse(recorder)
//			})
//		}
//	}
func TestGetChatMessagesAPI(t *testing.T) {
	user1, _ := randomUser(t)
	user2, _ := randomUser(t)

	chat := randomChat(user1.ID, user2.ID)

	n := 5
	messages := make([]db.GetChatMessagesRow, n)

	for i := 0; i < n; i++ {
		messages[i] = randomChatMessageRow(user1.ID)
	}

	testCases := []struct {
		name          string
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.JWTMaker)
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.JWTMaker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user1.ID, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetChatMessages(gomock.Any(), gomock.Eq(chat.ID)).
					Times(1).
					Return(messages, nil)
				store.EXPECT().
					GetChat(gomock.Any(), gomock.Eq(chat.ID)).
					Times(1).
					Return(chat, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchMessageRows(t, recorder.Body, messages)
			},
		},
		{
			name: "NoAuthorization",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.JWTMaker) {
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetChatMessages(gomock.Any(), gomock.Eq(chat.ID)).
					Times(0)
				store.EXPECT().
					GetChat(gomock.Any(), gomock.Eq(chat.ID)).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "InternalError",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.JWTMaker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user1.ID, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetChatMessages(gomock.Any(), gomock.Eq(chat.ID)).
					Times(1).
					Return([]db.GetChatMessagesRow{}, sql.ErrConnDone)
				store.EXPECT().
					GetChat(gomock.Any(), gomock.Eq(chat.ID)).
					Times(1).
					Return(chat, nil)

			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			url := "/api/v1/chats/" + strconv.FormatInt(chat.ID, 10) + "/messages/"
			request := httptest.NewRequest(http.MethodGet, url, nil)

			tc.setupAuth(t, request, server.jwtMaker)
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

func randomChatMessageRow(userId int64) db.GetChatMessagesRow {
	return db.GetChatMessagesRow{
		ID:      util.RandomInt(1, 1000),
		UserID:  userId,
		Content: util.RandomString(20),
	}
}
func randomMessage(chatId int64, userId int64) db.Message {
	return db.Message{
		ID:      util.RandomInt(1, 1000),
		UserID:  userId,
		ChatID:  chatId,
		Content: util.RandomString(20),
	}
}

func requireBodyMatchMessageRows(t *testing.T, body *bytes.Buffer, Messages []db.GetChatMessagesRow) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotMessages []db.GetChatMessagesRow
	err = json.Unmarshal(data, &gotMessages)
	require.NoError(t, err)
	require.Equal(t, Messages, gotMessages)
}
