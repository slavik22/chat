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
  </template>
  
<script>
import axios from 'axios';
import authHeader from '@/services/auth-header';
    export default{
    data() {
        return {
            password: ""
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
    },
    methods: {
         async updateProfile() {
          await axios.put("http://localhost:8080/api/v1/users/", {name: this.currentUser.name, login: this.currentUser.login, password: this.password}, { headers: authHeader() });
          this.$store.dispatch("auth/logout")
          this.$router.push("/login");
        },
    },
    }
</script>

  <style scoped>
  .user-profile {
    padding: 20px;
  }
  
  h1 {
    font-size: 24px;
    margin-bottom: 20px;
  }
  
  .profile-form {
    max-width: 300px;
  }
  
  .profile-input {
    width: 100%;
    padding: 5px;
    margin-bottom: 10px;
    border: 1px solid #ccc;
    border-radius: 5px;
  }
  
  .save-button {
    padding: 10px 20px;
    background-color: #007bff;
    color: #fff;
    border: none;
    border-radius: 5px;
    cursor: pointer;
  }
  </style>
  