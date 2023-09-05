import ChatService from '../services/chat.service'
const user = JSON.parse(localStorage.getItem('user'));

const initialState = {
  chatList : {
    'data' : [],
  },
  chatUsers: {
    'data' : []
  }
}

export const chat = {
  namespaced: true,
  state: initialState,
  actions: {
    list({ commit }, userId) {
      return ChatService.list(userId).then(
        pg => {
          commit('listSuccess', pg);
          return Promise.resolve(pg);
        },
        error => {
          return Promise.reject(error);
        }
      );
    },
    listUsers({ commit }, chatId) {
      return ChatService.listChatUsers(chatId).then(
        pg => {
          commit('listChatUsersSuccess', pg);
          return Promise.resolve(pg);
        },
        error => {
          return Promise.reject(error);
        }
      );
    },

    deleteUser({ commit }, {userId, chatId}) {
      return ChatService.removeUserFromChat(chatId, userId).then(
        pg => {
          commit('httpSuccess', pg);
          return Promise.resolve(pg);
        },
        error => {
          return Promise.reject(error);
        }
      );
    },
    delete({ commit }, chatId) {
      return ChatService.delete(chatId).then(
        pg => {
          commit('httpSuccess', pg);
          return Promise.resolve(pg);
        },
        error => {
          return Promise.reject(error);
        }
      );
    },
    create({ commit }, chatName) {
      return ChatService.create(chatName).then(
         res => {
            commit('httpSuccess', res)
            return Promise.resolve(res);
        },
        error => {
          return Promise.reject(error);
        }
      );
    },
    addUser({ commit }, {chatId, login}) {
      console.log(login)
      return ChatService.addUserToChat(chatId ,login).then(
         res => {
            commit('httpSuccess', res)
            return Promise.resolve(res);
        },
        error => {
          return Promise.reject(error);
        }
      );
    },
  },
  mutations: {
    listSuccess(state, pg) {
      state.chatList = pg.data;
    },
    listChatUsersSuccess(state, pg) {
      state.chatUsers = pg.data;
    },
    httpSuccess(res) {
      console.log(user)
      console.log(res)
    },
  },
  getters : {
    chatList(state) {
      return state.chatList
    },
    userList(state) {
      return state.chatUsers
    },
  }
};