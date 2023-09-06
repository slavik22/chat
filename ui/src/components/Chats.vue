<template>
  <div class="chat-rooms">
    <h1>Chat</h1>
    <ul>
      <li v-for="chat in chatList" :key="chat.id">
        <router-link :to="'/chat/' + chat.id" class="chat-room-link">
          {{ chat.name1 === currentUser.user.name ? chat.name1: chat.name2 }}
        </router-link>
        <button @click="deleteChat(chat.id)">Delete</button>
      </li>
    </ul>
  </div>
</template>

<script>
import { mapActions, mapGetters } from 'vuex';

export default {
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