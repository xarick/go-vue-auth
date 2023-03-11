import _ from 'lodash';
window._ = _;

/**
 * We'll load the axios HTTP library which allows us to easily issue requests
 * to our Laravel back-end. This library automatically handles sending the
 * CSRF token as a header based on the value of the "XSRF" token cookie.
 */

import axios from 'axios';
import { ref, reactive, computed, watch, h, onMounted } from 'vue'
import { useRoute } from "vue-router";
import { RouterLink, createRouter, createWebHistory } from 'vue-router'


window.axios = axios;
window.axios.defaults.headers.common['X-Requested-With'] = 'XMLHttpRequest';
window.axios.defaults.headers.common['X-Locale'] = window.localStorage.getItem('settings') ? JSON.parse(window.localStorage.getItem('settings')).locale : 'uz'
window.ref = ref
window.reactive = reactive
window.computed = computed
window.watch = watch
window.h = h
window.onMounted = onMounted
window.useRoute = useRoute
window.RouterLink = RouterLink
window.createRouter = createRouter
window.createWebHistory = createWebHistory


/**
 * Echo exposes an expressive API for subscribing to channels and listening
 * for events that are broadcast by Laravel. Echo and event broadcasting
 * allows your team to easily build robust real-time web applications.
 */
if (!window.localStorage.getItem('settings')) {
    window.localStorage.setItem('settings', JSON.stringify({
        theme: 'light',
        locale: 'oz',
        sideBarCollapsed: false
    }))
}
