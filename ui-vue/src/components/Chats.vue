<template>
  <div class="container mt-4">
    <ul class="list-group">
      <li class="list-group-item" v-for="chat in chatList" :key="chat.id">
        <img src="http://localhost:8080/api/v1/download/2" alt="profile" class="img-circle">
        <router-link :to="'/chat/' + chat.id" class="chat-room-link">
          {{ chat.name1 === currentUser.user.name ? chat.name1: chat.name2 }}
        </router-link>
        <button class="btn btn-danger" @click="deleteChat(chat.id)">Delete</button>
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
    getImage(user1Id, user2Id){
      return "http://localhost:8080/api/v1/download/" + this.currentUser.id === user1Id ? user1Id : user2Id
    },
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