import axios from "axios";

axios.defaults.withCredentials = true;
axios.defaults.baseURL = "http://127.0.0.1:8080/api";
axios.defaults.headers.common['Content-Type'] = 'application/json';
