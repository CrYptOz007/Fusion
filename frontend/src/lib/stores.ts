import { writable } from 'svelte/store';
import api from '$lib/middlewares/api';
import axios from 'axios';
import { useRefresh } from './hooks/auth';

const createErrorStore = () => {
	const { subscribe, set } = writable('');

	return {
		subscribe,
		set: (k: string) => set(k),
		reset: () => set('')
	};
};

const createAuthStore = () => {
	const { subscribe, set } = writable(false);

	return {
		subscribe,
		login: () => set(true),
		logout: () => set(false),
		isLoggedIn: async () => {
			const response = await useRefresh();
			if (response.status === 200 && response.data.token) {
				sessionStorage.setItem('authToken', response.data.token);
				set(true);
			} else {
				set(false);
			}
		}
	};
};

const createCount = () => {
	const { subscribe, set, update } = writable(1);

	return {
		subscribe,
		increment: () => update((n) => n + 1),
		decrement: () => update((n) => n - 1),
		reset: () => set(1)
	};
};

export const count = createCount();

export const authStore = createAuthStore();

export const errorStore = createErrorStore();
