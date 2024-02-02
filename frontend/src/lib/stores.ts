import { writable } from 'svelte/store';

const createTokenStore = () => {
	const { subscribe, set, update } = writable('');

	return {
		subscribe,
		set: (k: string) => update((k) => k + 1),
		reset: () => set('')
	};
};

const createErrorStore = () => {
	const { subscribe, set } = writable('');

	return {
		subscribe,
		set: (k: string) => set(k),
		reset: () => set('')
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

export const tokenStore = createTokenStore();

export const errorStore = createErrorStore();
