<template>
  <div class="chat-rooms">
    <h1>Chat Rooms</h1>
    <ul>
      <li v-for="chatRoom in chatList" :key="chatRoom.chat_room_id">
        <router-link :to="'/chat/' + chatRoom.chat_room_id" class="chat-room-link">{{ chatRoom.name }}</router-link>
        <button @click="deleteChat(chatRoom.chat_room_id)">Delete</button>
      </li>
    </ul>
    <div class="input-box">
          <input v-model="chatName" @keyup.enter="addChat" placeholder="Type chat name..." class="message-input" />
          <button @click="addChat" class="create-button">Create New Chat Room</button>
    </div>
  </div>
</template>

<script>
import { mapActions, mapGetters } from 'vuex';

export default {
  data() {
      return {
        chatName: "",
      };
    },
  computed: {
    ...mapGetters({
            chatList : 'chat/chatList',
            currentUser : 'auth/userCurr'
        })
  },
  methods: {
    deleteChat(chatId){
      this.chatDelete(chatId)
      location.reload() 
    },
    addChat() {
      this.chatCreate(this.chatName)
      location.reload() 
    },

    ...mapActions({
      fetchChatList: 'chat/list',
      chatDelete: 'chat/delete',
      chatCreate: 'chat/create'
    })
  },
  mounted() {
    if (!this.currentUser) {
            this.$router.push("/login");
    }
    this.fetchChatList()
  },
};
</script>

<style scoped>
.chat-rooms {
  padding: 20px;
}

h1 {
  font-size: 24px;
  margin-bottom: 20px;
}

ul {
  list-style: none;
  padding: 0;
}

li {
  margin: 10px 0;
}

.chat-room-link {
  text-decoration: none;
  color: #007bff;
  font-weight: bold;
}

.create-button {
  margin-top: 20px;
  padding: 10px 20px;
  background-color: #007bff;
  color: #fff;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}
</style>
