import { createRouter, createWebHistory } from "vue-router";

import Home from "../components/Home.vue";

const token = localStorage.getItem('token')
console.log(token);
const routes = [
  {
    path: "/",
    name: "Home",
    component: Home,
    meta: { title: "Home page" },
    beforeEnter(to, from, next) {
      if (token) {
        next()
      } else {
        next({ name: 'Login' })
      }
    }
  },
  {
    path: "/login",
    name: "Login",
    component: () => import("../components/Login.vue"),
    meta: { title: "Login page" },
  },
  {
    path: "/register",
    name: "Register",
    component: () => import("../components/Register.vue"),
    meta: { title: "Register page" },
  },
  {
    path: "/forgot-password",
    name: "ForgotPassword",
    component: () => import("../components/ForgotPassword.vue"),
  },
  {
    path: "/password-reset/:token",
    name: "ResetPassword",
    component: () => import("../components/ResetPassword.vue"),
  },
  {
    name: 'NotFound',
    path: '/:pathMatch(.*)*',
    component: () => import("../views/NotFound.vue"),
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach((to, from, next) => {
  document.title = to.meta.title
  next()
});

export default router;
