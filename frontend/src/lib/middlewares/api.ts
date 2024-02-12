import axios from 'axios';
import { authStore } from '$lib/stores';

const api = axios.create({
	baseURL: '/api',
	headers: {
		'Content-Type': 'application/json'
	},
	withCredentials: true
});

api.interceptors.request.use(
	async (config) => {
		const authToken = sessionStorage.getItem('token');
		if (authToken) {
			config.headers.Authorization = `Bearer ${authToken}`;
		}
		return config;
	},
	(error) => {
		return Promise.reject(error);
	}
);

api.interceptors.request.use(
	(response) => {
		return response;
	},
	async (error) => {
		const originalRequest = error.config;
		if (error.response.status === 401 && !originalRequest._retry) {
			originalRequest._retry = true;
			try {
				console.log('trying');
				const response = await axios.get('/api/auth/refresh');
				sessionStorage.setItem('token', response.data.token);
				authStore.login();
				return api.request(error.config);
			} catch (e) {
				sessionStorage.removeItem('token');
				console.log('am i logging out');
				authStore.logout();
			}
		}
		return Promise.reject(error);
	}
);

export default api;
