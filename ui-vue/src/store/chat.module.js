import ChatService from '../services/chat.service'
const user = JSON.parse(localStorage.getItem('user'));

const initialState = {
  chatList : {
    'data' : [],
  }
}

export const chat = {
  namespaced: true,
  state: initialState,
  actions: {
    async list({ commit }) {
      await ChatService.list().then(pg => commit('listSuccess', pg));
    },
    async delete({ commit }, chatId) {
      await ChatService.delete(chatId).then(pg => commit('httpSuccess', pg));
    },
    async create({ commit }, chatName) {
      await ChatService.create(chatName).then(res => commit('httpSuccess', res))
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