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
    </div>
  </template>
  <script>
import { mapGetters } from 'vuex';

  export default {
    data() {
      return {
        chatId: this.$route.params.chatId,
        newMessage: '', 
      };
    },
    computed: {
      ...mapGetters({
            messageList : 'message/messageList',
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
      this.conn.onmessage = this.receiveMessage;
    },
  };
  </script>
  