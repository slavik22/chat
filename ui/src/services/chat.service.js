import axios from 'axios';
import authHeader from './auth-header';

const API_URL = 'http://localhost:8080/api/v1/chats/';

class ChatService {
  list() {
    return axios.get(API_URL, { headers: authHeader() });
  }

  listChatMessages(chatId){
      return axios.get(API_URL  + chatId + "/messages/", { headers: authHeader() });
  }

  initWebSocketConn(userId, chatId){
    console.log("Service " + chatId)
    return new WebSocket("ws://localhost:8080/chatroom/" + chatId + "/user/" + userId);
  }

  create(chatName) {
    return axios.post(API_URL, {name: chatName}, { headers: authHeader() });
  }

  delete(chatId) {
    return axios.delete(API_URL + chatId, { headers: authHeader() });
  }
}

export default new ChatService();