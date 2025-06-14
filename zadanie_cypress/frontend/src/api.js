import axios from "axios";

export const api = axios.create({
  baseURL: process.env.REACT_APP_API_URL || "http://localhost:8080",
  withCredentials: true,
  headers: {
    "Context-Type": "application/json",
  },
});
