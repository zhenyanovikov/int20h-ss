import axios from "axios";
import { BASE_URL } from "../constants/api";
import { STORE_NAME } from "../constants/store";

const apiClient = axios.create({
  baseURL: BASE_URL,
});

apiClient.interceptors.request.use(
  (config) => {
    const clientStore = JSON.parse(localStorage.getItem(STORE_NAME));
    const { token } = clientStore.state;

    if (token) {
      config.headers["Authorization"] = `Bearer ${token}`;
    }

    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

export default apiClient;
