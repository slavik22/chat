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

  listChatUsers(chatId){
    return axios.get(API_URL  + chatId + "/users/", { headers: authHeader() });
}

  initWebSocketConn(userId, chatId){
    console.log("Service " + chatId)
    return new WebSocket("ws://localhost:8080/chatroom/" + chatId + "/user/" + userId);
  }

  create(chatName) {
    return axios.post(API_URL, {name: chatName}, { headers: authHeader() });
  }

  addUserToChat(chatId, login) {
    return axios.post(API_URL + chatId + "/users/",{login: login}, { headers: authHeader() });
  }


  removeUserFromChat(chatId, userId) {
    return axios.delete(API_URL + chatId + "/users/" + userId, { headers: authHeader() });
  }

  getChatMessages(chatId){
    return axios.get(API_URL + chatId + "/messages/", { headers: authHeader() });

  }

  delete(chatId) {
    return axios.delete(API_URL + chatId, { headers: authHeader() });
  }
}

export default new ChatService();