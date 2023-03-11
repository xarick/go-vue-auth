import '@/bootstrap';
import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'

import VueCookies from 'vue-cookies'
import naive from 'naive-ui'
import { NMessageProvider } from 'naive-ui';
import Navbar from '@/components/Navbar.vue'
import Sidebar from '@/components/Sidebar.vue'
import Footer from '@/components/Footer.vue'

import './assets/main.css'

const app = createApp(h(NMessageProvider, () => h(App)))

app.use(createPinia())
app.use(router)
app.use(VueCookies, { expire: "7d" });
app.use(naive);
app.component('Navbar', Navbar)
app.component('Sidebar', Sidebar)
app.component('Footer', Footer)

app.mount('#app')
