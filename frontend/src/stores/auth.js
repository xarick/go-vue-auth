import { defineStore } from "pinia";
import axios from "axios";

export const useAuthStore = defineStore("auth", {
  state: () => ({
    authUser: null,
    authError: null,
    authStatus: null,
  }),
  getters: {
    user: (state) => state.authUser,
    error: (state) => state.authError,
    status: (state) => state.authStatus,
  },
  actions: {
    // async getToken() {
    //   await axios.get("/sanctum/csrf-cookie");
    // },

    async getUser() {
      const token = localStorage.getItem('token')
      if (token) {
        const data = await axios.get("/user/get-user", {
          headers: {
            'Authorization': `Bearer ${token}`
          }
        });
        this.authUser = data.data;
      }
    },

    async handleLogin(data) {
      try {
        const response = await axios.post("/auth/login", {
          email: data.email,
          password: data.password,
        });
        localStorage.setItem('token', response.data.token)
        setTimeout(() => {
          this.router.push("/");
        }, 500);
      } catch (error) {
        if (error.response.status != 200) {
          this.authError = error.response.data.error;
        }
      }
    },

    async handleRegister(data) {
      try {
        await axios.post("/auth/register", {
          name: data.name,
          email: data.email,
          password: data.password,
          passwordConfirm: data.passwordConfirm,
        });
        this.router.push("/");
      } catch (error) {
        if (error.response.status != 200) {
          this.authError = error.response.data.error;
        }
      }
    },

    async handleLogout() {
      await axios.post("/logout");
      // localStorage.removeItem('token')
      this.authUser = null;
    },

    async handleForgotPassword(email) {
      try {
        const response = await axios.post("/forgot-password", {
          email: email,
        });
        this.authStatus = response.data.status;
      } catch (error) {
        if (error.response.status === 422) {
          // this.authErrors = error.response.data.errors;
        }
      }
    },

    async handleResetPassword(resetData) {
      try {
        const response = await axios.post("/reset-password", resetData);
        this.authStatus = response.data.status;
      } catch (error) {
        if (error.response.status === 422) {
          // this.authErrors = error.response.data.errors;
        }
      }
    },

  },
});
