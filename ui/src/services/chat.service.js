import axios from 'axios';
import authHeader from './auth-header';

const API_URL = 'http://localhost:8080/api/v1/chats/';

class ChatService {
  list(userId) {
    return axios.get(API_URL + 'user/' + userId, { headers: authHeader() });
  }

  listChatMessages(chatId){
      return axios.get(API_URL  + chatId + "/messages/", { headers: authHeader() });
  }

  initWebSocketConn(userId, chatId){
    console.log("Service " + chatId)
    return new WebSocket("ws://localhost:8080/chatroom/" + chatId + "/user/" + userId);
  }

  create(payload) {
    return axios.post(API_URL + '/', payload, { headers: authHeader() });
  }

  addUserToChat(chatId, userId, payload) {
    return axios.post(API_URL + chatId + "/users/" + userId, payload, { headers: authHeader() });
  }
  removeUserFromChat(chatId, userId) {
    return axios.delete(API_URL + chatId + "/users/" + userId, { headers: authHeader() });
  }

  getChatMessages(chatId){
    return axios.get(API_URL + chatId + "/messages/", { headers: authHeader() });

  }

  _delete(chatId) {
    return axios.delete(API_URL + '/' + chatId, { headers: authHeader() });
  }
}

export default new ChatService();