import ChatService from '../services/chat.service'
// const user = JSON.parse(localStorage.getItem('user'));

const initialState = {
  messageList : {
    'data' : [],
  },
}

export const message = {
  namespaced: true,
  state: initialState,
  actions: {
    list({ commit }, chatId) {
      return ChatService.listChatMessages(chatId)
      .then(
        pg => {
          commit('listSuccess', pg);
          return Promise.resolve(pg);
        },
        error => {
          return Promise.reject(error);
        }
      );
    },
  },
  mutations: {
    listSuccess(state, pg) {
      state.messageList = pg.data;
    },
  },
  getters : {
    messageList(state) {
      return state.messageList
    },
  }
};