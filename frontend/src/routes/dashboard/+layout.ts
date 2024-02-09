import { authStore } from '$lib/stores';
import { redirect } from '@sveltejs/kit';

export async function load() {
	authStore.subscribe((value) => {
		if (!value) {
			throw redirect(302, '/login');
		}
	});
}
