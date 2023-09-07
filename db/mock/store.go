// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/slavik22/chat/db/sqlc (interfaces: Store)

// Package mock_sqlc is a generated GoMock package.
package mock_sqlc

import (
	context "context"
	reflect "reflect"

	db "github.com/slavik22/chat/db/sqlc"
	gomock "go.uber.org/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// AddBlackList mocks base method.
func (m *MockStore) AddBlackList(arg0 context.Context, arg1 db.AddBlackListParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddBlackList", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddBlackList indicates an expected call of AddBlackList.
func (mr *MockStoreMockRecorder) AddBlackList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBlackList", reflect.TypeOf((*MockStore)(nil).AddBlackList), arg0, arg1)
}

// AddFriend mocks base method.
func (m *MockStore) AddFriend(arg0 context.Context, arg1 db.AddFriendParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddFriend", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddFriend indicates an expected call of AddFriend.
func (mr *MockStoreMockRecorder) AddFriend(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFriend", reflect.TypeOf((*MockStore)(nil).AddFriend), arg0, arg1)
}

// CreateChatRoom mocks base method.
func (m *MockStore) CreateChatRoom(arg0 context.Context, arg1 db.CreateChatRoomParams) (db.Chat, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateChatRoom", arg0, arg1)
	ret0, _ := ret[0].(db.Chat)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateChatRoom indicates an expected call of CreateChatRoom.
func (mr *MockStoreMockRecorder) CreateChatRoom(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateChatRoom", reflect.TypeOf((*MockStore)(nil).CreateChatRoom), arg0, arg1)
}

// CreateMessage mocks base method.
func (m *MockStore) CreateMessage(arg0 context.Context, arg1 db.CreateMessageParams) (db.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMessage", arg0, arg1)
	ret0, _ := ret[0].(db.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMessage indicates an expected call of CreateMessage.
func (mr *MockStoreMockRecorder) CreateMessage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMessage", reflect.TypeOf((*MockStore)(nil).CreateMessage), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockStore) CreateUser(arg0 context.Context, arg1 db.CreateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStoreMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), arg0, arg1)
}

// DeleteBlackList mocks base method.
func (m *MockStore) DeleteBlackList(arg0 context.Context, arg1 db.DeleteBlackListParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBlackList", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBlackList indicates an expected call of DeleteBlackList.
func (mr *MockStoreMockRecorder) DeleteBlackList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBlackList", reflect.TypeOf((*MockStore)(nil).DeleteBlackList), arg0, arg1)
}

// DeleteChatRoom mocks base method.
func (m *MockStore) DeleteChatRoom(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteChatRoom", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteChatRoom indicates an expected call of DeleteChatRoom.
func (mr *MockStoreMockRecorder) DeleteChatRoom(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteChatRoom", reflect.TypeOf((*MockStore)(nil).DeleteChatRoom), arg0, arg1)
}

// DeleteFriend mocks base method.
func (m *MockStore) DeleteFriend(arg0 context.Context, arg1 db.DeleteFriendParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFriend", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFriend indicates an expected call of DeleteFriend.
func (mr *MockStoreMockRecorder) DeleteFriend(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFriend", reflect.TypeOf((*MockStore)(nil).DeleteFriend), arg0, arg1)
}

// DeleteMessage mocks base method.
func (m *MockStore) DeleteMessage(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMessage", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMessage indicates an expected call of DeleteMessage.
func (mr *MockStoreMockRecorder) DeleteMessage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMessage", reflect.TypeOf((*MockStore)(nil).DeleteMessage), arg0, arg1)
}

// GetBlackList mocks base method.
func (m *MockStore) GetBlackList(arg0 context.Context, arg1 int64) ([]db.GetBlackListRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlackList", arg0, arg1)
	ret0, _ := ret[0].([]db.GetBlackListRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlackList indicates an expected call of GetBlackList.
func (mr *MockStoreMockRecorder) GetBlackList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlackList", reflect.TypeOf((*MockStore)(nil).GetBlackList), arg0, arg1)
}

// GetChat mocks base method.
func (m *MockStore) GetChat(arg0 context.Context, arg1 int64) (db.Chat, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChat", arg0, arg1)
	ret0, _ := ret[0].(db.Chat)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChat indicates an expected call of GetChat.
func (mr *MockStoreMockRecorder) GetChat(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChat", reflect.TypeOf((*MockStore)(nil).GetChat), arg0, arg1)
}

// GetChatMessages mocks base method.
func (m *MockStore) GetChatMessages(arg0 context.Context, arg1 int64) ([]db.GetChatMessagesRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChatMessages", arg0, arg1)
	ret0, _ := ret[0].([]db.GetChatMessagesRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChatMessages indicates an expected call of GetChatMessages.
func (mr *MockStoreMockRecorder) GetChatMessages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChatMessages", reflect.TypeOf((*MockStore)(nil).GetChatMessages), arg0, arg1)
}

// GetFriends mocks base method.
func (m *MockStore) GetFriends(arg0 context.Context, arg1 int64) ([]db.GetFriendsRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFriends", arg0, arg1)
	ret0, _ := ret[0].([]db.GetFriendsRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFriends indicates an expected call of GetFriends.
func (mr *MockStoreMockRecorder) GetFriends(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFriends", reflect.TypeOf((*MockStore)(nil).GetFriends), arg0, arg1)
}

// GetUserById mocks base method.
func (m *MockStore) GetUserById(arg0 context.Context, arg1 int64) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserById", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserById indicates an expected call of GetUserById.
func (mr *MockStoreMockRecorder) GetUserById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockStore)(nil).GetUserById), arg0, arg1)
}

// GetUserByLogin mocks base method.
func (m *MockStore) GetUserByLogin(arg0 context.Context, arg1 string) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByLogin", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByLogin indicates an expected call of GetUserByLogin.
func (mr *MockStoreMockRecorder) GetUserByLogin(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByLogin", reflect.TypeOf((*MockStore)(nil).GetUserByLogin), arg0, arg1)
}

// GetUserChats mocks base method.
func (m *MockStore) GetUserChats(arg0 context.Context, arg1 int64) ([]db.GetUserChatsRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserChats", arg0, arg1)
	ret0, _ := ret[0].([]db.GetUserChatsRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserChats indicates an expected call of GetUserChats.
func (mr *MockStoreMockRecorder) GetUserChats(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserChats", reflect.TypeOf((*MockStore)(nil).GetUserChats), arg0, arg1)
}

// GetUserFromBlackList mocks base method.
func (m *MockStore) GetUserFromBlackList(arg0 context.Context, arg1 db.GetUserFromBlackListParams) (db.BlackList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserFromBlackList", arg0, arg1)
	ret0, _ := ret[0].(db.BlackList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserFromBlackList indicates an expected call of GetUserFromBlackList.
func (mr *MockStoreMockRecorder) GetUserFromBlackList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserFromBlackList", reflect.TypeOf((*MockStore)(nil).GetUserFromBlackList), arg0, arg1)
}

// GetUsers mocks base method.
func (m *MockStore) GetUsers(arg0 context.Context) ([]db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers", arg0)
	ret0, _ := ret[0].([]db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsers indicates an expected call of GetUsers.
func (mr *MockStoreMockRecorder) GetUsers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockStore)(nil).GetUsers), arg0)
}

// UpdateMessage mocks base method.
func (m *MockStore) UpdateMessage(arg0 context.Context, arg1 db.UpdateMessageParams) (db.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMessage", arg0, arg1)
	ret0, _ := ret[0].(db.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateMessage indicates an expected call of UpdateMessage.
func (mr *MockStoreMockRecorder) UpdateMessage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMessage", reflect.TypeOf((*MockStore)(nil).UpdateMessage), arg0, arg1)
}

// UpdateUser mocks base method.
func (m *MockStore) UpdateUser(arg0 context.Context, arg1 db.UpdateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockStoreMockRecorder) UpdateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockStore)(nil).UpdateUser), arg0, arg1)
}