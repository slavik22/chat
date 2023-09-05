<template>
    <div class="chat-room">
      <div class="messages">
        <div v-for="message in messageList" :key="message.id" class="message">
          {{ message.name }}: <strong>{{ message.content }}</strong>
        </div>
      </div>
      <div class="input-box">
        <input v-model="newMessage" @keyup.enter="sendMessage" placeholder="Type your message..." class="message-input" />
        <button @click="sendMessage" class="send-button">Send</button>
      </div>
      <div class="users">
        <h2>Users in this Chat</h2>

        <div class="input-box">
          <input v-model="newUserLogin" @keyup.enter="addUser" placeholder="Type new user login..." class="message-input" />
          <button @click="addUser" class="send-button">Send</button>
        </div>
       
        <ol>
          <li v-for="user in userList" :key="user.id">
            {{ user.name }}
          <button @click="deleteUser(user.id)">Delete</button>
            
          </li>
        </ol>
      </div>
    </div>
  </template>
  <script>
import { mapGetters } from 'vuex';

  export default {
    data() {
      return {
        chatId: this.$route.params.chatId,
        newMessage: '', 
        newUserLogin: ''
      };
    },
    computed: {
      ...mapGetters({
            messageList : 'message/messageList',
            userList : 'chat/userList',
            currentUser : 'auth/userCurr'
        }),
        conn: function(){
          return new WebSocket("ws://localhost:8080/chatroom/" + this.chatId + "/user/" + this.currentUser.user.id);
        }
    },
    methods: {
      fetchMessageList() {
        this.$store.dispatch('message/list', this.chatId)
      },
      fetchUserList() {
        this.$store.dispatch('chat/listUsers', this.chatId)
      },
      addUser() {
        this.$store.dispatch('chat/addUser', {chatId: this.chatId, login: this.newUserLogin})
        this.newUserLogin = ""
      },
      deleteUser(userId){
        this.$store.dispatch('chat/deleteUser', {chatId: this.chatId, userId: userId})

      },
      sendMessage() {
        if (!this.conn) {
                alert("Connection error.");
                  return;
          }
        this.conn.send(this.newMessage);
        this.newMessage = "";
      },

      receiveMessage(event){
        let msg = JSON.parse(event.data);
        this.messageList.push(msg)
      }
    },
    mounted() {
      this.fetchMessageList()
      this.fetchUserList()
      this.conn.onmessage = this.receiveMessage;
    },
  };
  </script>
  <style scoped>
  .chat-room {
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    height: 100%;
    padding: 20px;
  }
  
  .messages {
    flex-grow: 1;
    overflow-y: auto;
    padding: 10px;
    border: 1px solid #ccc;
    border-radius: 5px;
  }
  
  .message {
    margin: 5px 0;
    padding: 5px;
    background-color: #f0f0f0;
    border-radius: 5px;
  }
  
  .input-box {
    display: flex;
    margin-top: 10px;
  }
  
  .message-input {
    flex-grow: 1;
    padding: 5px;
    border: 1px solid #ccc;
    border-radius: 5px;
  }
  
  .send-button {
    margin-left: 10px;
    padding: 5px 10px;
    background-color: #007bff;
    color: #fff;
    border: none;
    border-radius: 5px;
    cursor: pointer;
  }
  
  .users {
    margin-top: 20px;
  }
  
  ul {
    list-style: none;
    padding: 0;
  }
  
  li {
    margin: 5px 0;
  }
  </style>
  