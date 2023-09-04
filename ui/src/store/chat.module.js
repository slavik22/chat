import ChatService from '../services/chat.service'
const user = JSON.parse(localStorage.getItem('user'));

const initialState = {
  chatList : {
    'data' : [],
  },
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
    create({ commit }, payload) {
      return ChatService.create(payload).then(
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
    httpSuccess(res) {
      console.log(user)
      console.log(res)
    },
  },
  getters : {
    chatList(state) {
      return state.chatList
    },
  }
};