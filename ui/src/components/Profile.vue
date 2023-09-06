<template>
    <div class="user-profile">
      <h1>User Profile</h1>
      <div class="profile-form">
        <div v-if="currentUser">
          <label for="name">Name:</label>
          <input v-model="currentUser.name" id="name"  class="profile-input" />
        </div>
        <div v-if="currentUser">
          <label for="login">Login:</label>
          <input v-model="currentUser.login" id="email" class="profile-input" />
        </div>
        <div v-if="currentUser">
          <label for="login">New password:</label>
          <input v-model="password" id="password" class="profile-input"/>
        </div>
        <button v-if="currentUser" @click="updateProfile" class="save-button">Save Changes</button>
      </div>
    </div>
    <div class="friends">
      <h2>Friends</h2>

      <ul v-if="friends">
        <li v-for="user in friends" :key="user.id">
          {{ user.name }}
           <button @click="deleteFriend(user.id)">Delete</button>
        </li>
      </ul>
    </div>
    <div class="blackList">
      <h2>BlackList</h2>

      <ul v-if="blackList">
        <li v-for="user in blackList" :key="user.id">
          {{ user.name }}
          <button @click="deleteBlackList(user.id)">Delete</button>
        </li>
      </ul>
    </div>
  </template>
  
<script>
import axios from 'axios';
import authHeader from '@/services/auth-header';
    export default{
    data() {
        return {
            password: "",
            friends: null,
            blackList: null,
        };
    },

    computed: {
        currentUser() {
            return this.$store.state.auth.user;
        },
    },
    mounted() {
        if (!this.currentUser) {
            this.$router.push("/login");
        }

        axios.get('http://localhost:8080/api/v1/users/friends/', { headers: authHeader() })
          .then((response) => {
          this.friends = response.data;
        })
          .catch((error) => {
          console.error('Error fetching users:', error);
        });

        axios.get('http://localhost:8080/api/v1/users/blackList/', { headers: authHeader() })
          .then((response) => {
          this.blackList = response.data;
        })
          .catch((error) => {
          console.error('Error fetching users:', error);
        });

    },
    methods: {
         async updateProfile() {
          await axios.put("http://localhost:8080/api/v1/users/", {name: this.currentUser.name, login: this.currentUser.login, password: this.password}, { headers: authHeader() });
          this.$store.dispatch("auth/logout")
          this.$router.push("/login");
        },
        async deleteFriend(userId){
          await axios.delete("http://localhost:8080/api/v1/users/friends/" + userId,{ headers: authHeader() });
          location.reload()

        },
        async deleteBlackList(userId){
          await axios.delete("http://localhost:8080/api/v1/users/blackList/" + userId, { headers: authHeader() });
          location.reload()
        }
    },
    }
</script>
  