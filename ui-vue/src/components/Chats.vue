<template>
  <div class="container mt-4">
    <ul class="list-group">
      <li class="list-group-item d-flex justify-content-between align-items-center  " v-for="chat in chatList" :key="chat.id">
        <img :src="getImage(chat.user1id, chat.user2id)" alt="profile" class="img" style="border-radius: 100%;" width="100">
        <router-link :to="'/chat/' + chat.id" class="chat-room-link">
          <span class="fs-2">
          {{ chat.name1 === currentUser.user.name ? chat.name2 : chat.name1 }}
          </span>
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

      const id = this.currentUser.user.id == user1Id ? user2Id : user1Id 
      let l = "http://localhost:8080/api/v1/download/" + id
      return l
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