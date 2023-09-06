<template>
    <div class="container mt-3">
        <input type="text" v-model="searchQuery" placeholder="Search users" class="form-control"/>
        <ul class="list-group">
            <li v-for="user in filteredUsers" :key="user.id" class="list-group-item mt-2">
                {{ user.name }}
                <button class="btn btn-primary" @click="createChat(user.id)">Create chat</button>
                <button class="btn btn-success" @click="addToFriends(user.id)">Add to friends</button>
                <button class="btn btn-danger" @click="addToBlackList(user.id)">Add to blackList</button>
            </li>
        </ul>
    </div>
</template>
  
<script>
import axios from 'axios';
import authHeader from '../services/auth-header';
import { mapGetters } from 'vuex';

export default {
    data() {
        return {
            users: [],
            searchQuery: '',
        };
    },
    methods: {
        addToFriends(userId) {
            axios.post('http://localhost:8080/api/v1/users/friends/' + userId, null, { headers: authHeader() })
                .then(() => {
                    alert("Friend added succesfully")
                })
                .catch(() => {
                    alert("Already friends");
                });
        },
        addToBlackList(userId) {
            axios.post('http://localhost:8080/api/v1/users/blackList/' + userId, null, { headers: authHeader() })
                .then(() => {
                    alert("Black list added succesfully")
                })
                .catch(() => {
                    alert("Already in blacklist");
                });
        },
        createChat(userId){
            axios.post('http://localhost:8080/api/v1/chats/users/' + userId, null,{ headers: authHeader() })
                .then(() => {
                    alert("Chat added successfully")
                })
                .catch((err) => {
                    alert(err.response.data.message);
                });
        }
    },
    computed: {
        ...mapGetters({
            currentUser: 'auth/userCurr'
        }),
        filteredUsers() {
            return this.users.filter((user) =>
                user.id !== this.currentUser.user.id &&
                (user.name.toLowerCase().includes(this.searchQuery.toLowerCase()) ||
                    user.login.toLowerCase().includes(this.searchQuery.toLowerCase()))

            );
        },
    },
    mounted() {
        if (!this.currentUser) {
            this.$router.push("/login");
        }
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
