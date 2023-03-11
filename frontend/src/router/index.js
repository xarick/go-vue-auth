const is_user_logged = () => !!$cookies.get('token_user')
new URL(window.location.href).pathname

const router = createRouter({
    history: createWebHistory(),
    routes: [
        {
            path: '/',
            beforeEnter(to, from, next) {
                next({ name: 'login' })
            },
        },
        {
            path: '/login',
            query: { code: 'code' },
            name: 'login',
            component: () => import('../auth/Login.vue'),
            meta: { title: "Laravel | Kirish" },
            beforeEnter(to, from, next) {
                if (is_user_logged()) {
                    next({ name: 'home' })
                } else {
                    next()
                }
            },
        },
        {
            path: '/cabinet',
            name: 'cabinet',
            component: () => import('../auth/Auth.vue'),
            children: [
                {
                    path: '/cabinet/home',
                    name: 'cabinet.home',
                    component: () => import('../views/Home.vue'),
                    meta: { title: "Laravel | Bosh sahifa" },
                },
                {
                    path: '/cabinet/menu',
                    name: 'cabinet.menu',
                    component: () => import('../views/Menu.vue'),
                    meta: { title: "Laravel | Menyu" },
                },
            ],
            beforeEnter(to, from, next) {
                next()
                // if (is_user_logged()) {
                //     next()
                // } else {
                //     next({ name: 'login' })
                // }
            },
        }
    ],
})

router.beforeEach((to, from, next) => {
    document.title = to.meta.title
    next()
});

export default router
