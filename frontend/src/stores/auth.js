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
      // await this.getToken();
      const data = await axios.get("/api/user");
      this.authUser = data.data;
    },

    async handleLogin(data) {
      // this.authErrors = [];
      // await this.getToken();

      try {
        await axios.post("/login", {
          email: data.email,
          password: data.password,
        });
        this.router.push("/");
      } catch (error) {
        if (error.response.status === 422) {
          this.authErrors = error.response.data.errors;
        }
      }
    },

    async handleRegister(data) {
      // this.authErrors = [];
      // await this.getToken();
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
          console.log(error.response.data.error);
        }
      }
    },

    async handleLogout() {
      await axios.post("/logout");
      this.authUser = null;
    },

    async handleForgotPassword(email) {
      // this.authErrors = [];
      // this.getToken();
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
      // this.authErrors = [];
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
