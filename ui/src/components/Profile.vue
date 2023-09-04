<template>
    <div class="user-profile">
      <h1>User Profile</h1>
      <div class="profile-form">
        <div>
          <label for="name">Name:</label>
          <input v-model="user.name" id="name" class="profile-input" />
        </div>
        <div>
          <label for="login">Login:</label>
          <input v-model="user.email" id="email" class="profile-input" />
        </div>
        <button @click="updateProfile" class="save-button">Save Changes</button>
      </div>
    </div>
  </template>
  
<script>
    export default{
    data() {
        return {
            user: {
                name: currentUser.name,
                login: currentUser.login
            },
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
        updateProfile() {
            // Implement logic to update the user's profile
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
  