import { createApp } from 'vue'
import { createRouter, createWebHistory } from 'vue-router'
import App from './App.vue'
import Register from './components/Register.vue'
import Login from './components/Login.vue'
import Chat from './components/Chat.vue'
import Chats from './components/Chats.vue'
import Users from './components/Users.vue'
import Profile from './components/Profile.vue'

import store from './store';

const router = createRouter({
    history: createWebHistory(),
    routes: [
        { path: '/register', component: Register },
        { path: '/login', component: Login },
        { path: '/users', component: Users },
        { path: '/profile', component: Profile },
        { path: '/chat/:chatId', component: Chat },
        { path: '/', component: Chats },
    ]
});

const app = createApp(App)
app.use(router)
app.use(store)

app.mount('#app')
