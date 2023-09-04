<template>
<div class="header">
  <a href="/" class="logo">Chat</a>
  <div class="header-right">
    <a class="active" href="#home">Home</a>
    <a v-if="loggedIn" to="/login" @click.prevent="logout">Log Out</a>

    <router-link v-if="!loggedIn" to="/register">Sign Up</router-link>
    <router-link v-if="!loggedIn" to="/login">Sign In</router-link><br>
  </div>
</div>

<div style="padding-left:20px">
  <router-view></router-view>
</div>
</template>

<script>

export default {
  name: 'App',
  components: {},
  computed: {
    loggedIn() {
      return this.$store.state.auth.status.loggedIn;
    }
  },
  methods: {
    logout(){
      this.$store.dispatch('auth/logout');
      this.$router.push('/login');
    }
    }
  }
</script>
<style>
* {box-sizing: border-box;}

body { 
  margin: 0;
  font-family: Arial, Helvetica, sans-serif;
}

.header {
  overflow: hidden;
  background-color: #f1f1f1;
  padding: 20px 10px;
}

.header a {
  float: left;
  color: black;
  text-align: center;
  padding: 12px;
  text-decoration: none;
  font-size: 18px; 
  line-height: 25px;
  border-radius: 4px;
}

.header a.logo {
  font-size: 25px;
  font-weight: bold;
}

.header a:hover {
  background-color: #ddd;
  color: black;
}

.header a.active {
  background-color: dodgerblue;
  color: white;
}

.header-right {
  float: right;
}

@media screen and (max-width: 500px) {
  .header a {
    float: none;
    display: block;
    text-align: left;
  }
  
  .header-right {
    float: none;
  }
}
</style>
