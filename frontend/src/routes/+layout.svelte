<script lang="ts">
	import '../app.pcss';
	import { Navbar, NavBrand } from 'flowbite-svelte';
	import { QueryClient, QueryClientProvider } from '@sveltestack/svelte-query';

	const queryClient = new QueryClient();
</script>

<Navbar>
	<NavBrand href="/">
		<span class="self-center whitespace-nowrap text-2xl font-semibold dark:text-white">Fusion</span>
	</NavBrand>
</Navbar>

<QueryClientProvider client={queryClient}>
	<slot />
</QueryClientProvider>
