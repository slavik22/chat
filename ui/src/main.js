import { createApp } from 'vue'
import { createRouter, createWebHistory } from 'vue-router'
import App from './App.vue'
import Register from './components/Register.vue'
import Login from './components/Login.vue'
import Chat from './components/Chat.vue'
import Chats from './components/Chats.vue'
import store from './store';

const router = createRouter({
    history: createWebHistory(),
    routes: [
        { path: '/register', component: Register },
        { path: '/login', component: Login },
        { path: '/', component: Chats },
        { path: '/chat/:chatId', component: Chat },
    ]
});

const app = createApp(App)
app.use(router)
app.use(store)

app.mount('#app')
