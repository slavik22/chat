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

func TestCreateChatAPI(t *testing.T) {
	user1, _ := randomUser(t)
	user2, _ := randomUser(t)
	chat := randomChat(user1.ID, user2.ID)

	testCases := []struct {
		name          string
		user2Id       int64
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.JWTMaker)
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name:    "OK",
			user2Id: user2.ID,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.JWTMaker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user1.ID, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.CreateChatRoomParams{
					User1ID: user1.ID,
					User2ID: user2.ID,
				}
				store.EXPECT().
					CreateChatRoom(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(chat, nil)

				args2 := db.GetUserFromBlackListParams{UserID: user2.ID, FriendID: user1.ID}

				store.EXPECT().
					GetUserFromBlackList(gomock.Any(), gomock.Eq(args2)).
					Times(1).
					Return(db.BlackList{}, sql.ErrNoRows)

			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusCreated, recorder.Code)
			},
		},
		{
			name:    "NoAuthorization",
			user2Id: user2.ID,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.JWTMaker) {
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateChatRoom(gomock.Any(), gomock.Any()).
					Times(0)

				store.EXPECT().
					GetUserFromBlackList(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name:    "InternalError",
			user2Id: user2.ID,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.JWTMaker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user1.ID, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateChatRoom(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Chat{}, sql.ErrConnDone)

				args2 := db.GetUserFromBlackListParams{UserID: user2.ID, FriendID: user1.ID}

				store.EXPECT().
					GetUserFromBlackList(gomock.Any(), gomock.Eq(args2)).
					Times(1).
					Return(db.BlackList{}, sql.ErrNoRows)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:    "InvalidUser2Id",
			user2Id: -1,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.JWTMaker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user1.ID, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				args := db.CreateChatRoomParams{User1ID: user1.ID, User2ID: -1}

				store.EXPECT().
					CreateChatRoom(gomock.Any(), gomock.Eq(args)).
					Times(1).
					Return(db.Chat{}, sql.ErrNoRows)

				store.EXPECT().
					GetUserFromBlackList(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.BlackList{}, sql.ErrNoRows)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},

		{
			name:    "IsInBlackList",
			user2Id: user2.ID,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.JWTMaker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user1.ID, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateChatRoom(gomock.Any(), gomock.Any()).
					Times(0)

				args2 := db.GetUserFromBlackListParams{UserID: user2.ID, FriendID: user1.ID}

				store.EXPECT().
					GetUserFromBlackList(gomock.Any(), gomock.Eq(args2)).
					Times(1).
					Return(db.BlackList{}, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
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

			url := "/api/v1/chats/users/" + strconv.FormatInt(tc.user2Id, 10)
			request := httptest.NewRequest(http.MethodPost, url, nil)

			tc.setupAuth(t, request, server.jwtMaker)
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}
func TestGetChatsAPI(t *testing.T) {
	user1, _ := randomUser(t)

	n := 5
	chats := make([]db.GetUserChatsRow, n)

	for i := 0; i < n; i++ {
		user2, _ := randomUser(t)
		chats[i] = randomUserChatRow(user1.Name, user2.Name)
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
					GetUserChats(gomock.Any(), gomock.Eq(user1.ID)).
					Times(1).
					Return(chats, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchChatRows(t, recorder.Body, chats)
			},
		},
		{
			name: "NoAuthorization",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.JWTMaker) {
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUserChats(gomock.Any(), gomock.Eq(user1.ID)).
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
					GetUserChats(gomock.Any(), gomock.Eq(user1.ID)).
					Times(1).
					Return([]db.GetUserChatsRow{}, sql.ErrConnDone)
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

			url := "/api/v1/chats/"
			request := httptest.NewRequest(http.MethodGet, url, nil)

			tc.setupAuth(t, request, server.jwtMaker)
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

func randomUserChatRow(user1Name string, user2Name string) db.GetUserChatsRow {
	return db.GetUserChatsRow{
		ID:    util.RandomInt(1, 1000),
		Name1: user1Name,
		Name2: user2Name,
	}
}
func randomChat(user1Id int64, user2Id int64) db.Chat {
	return db.Chat{
		ID:      util.RandomInt(1, 1000),
		User1ID: user1Id,
		User2ID: user2Id,
	}
}

func requireBodyMatchChatRows(t *testing.T, body *bytes.Buffer, chats []db.GetUserChatsRow) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotChats []db.GetUserChatsRow
	err = json.Unmarshal(data, &gotChats)
	require.NoError(t, err)
	require.Equal(t, chats, gotChats)
}
