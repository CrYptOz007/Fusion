import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';

export default defineConfig({
	mode: process.env.ENVIRONMENT || 'development',
	server: {
		port: 4000
	},
	plugins: [sveltekit()]
});
