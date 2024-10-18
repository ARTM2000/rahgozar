import axios from 'axios'

export const httpClient = axios.create({
    baseURL: process.env.SERVER_BASE_URL,
    timeout: 10000,
});
