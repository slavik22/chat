<template>
    <div>
      <input type="text" v-model="searchQuery" placeholder="Search users" />
      <ul>
        <li v-for="user in filteredUsers" :key="user.id">{{ user.name }}</li>
      </ul>
    </div>
  </template>
  
  <script>
  import axios from 'axios';
  import authHeader from '../services/auth-header';
  
  export default {
    data() {
      return {
        users: [],
        searchQuery: '',
      };
    },
    computed: {
      filteredUsers() {
        return this.users.filter((user) =>
          user.name.toLowerCase().includes(this.searchQuery.toLowerCase()) || 
          user.login.toLowerCase().includes(this.searchQuery.toLowerCase()) 
        );
      },
    },
    mounted() {
      axios.get('http://localhost:8080/api/v1/users/', { headers: authHeader() })
        .then((response) => {
          this.users = response.data;
        })
        .catch((error) => {
          console.error('Error fetching users:', error);
        });
    },
  };
  </script>
  
<style>
input {
  width: 100%;
  padding: 10px;
  margin-bottom: 20px;
  font-size: 16px;
  border: 1px solid #ccc;
  border-radius: 4px;
}

ul {
  list-style-type: none;
  padding: 0;
}

li {
  margin: 10px 0;
  padding: 10px;
  border: 1px solid #ccc;
  border-radius: 4px;
  background-color: #f9f9f9;
  transition: background-color 0.3s;
}

li:hover {
  background-color: #e0e0e0;
}
</style>

  