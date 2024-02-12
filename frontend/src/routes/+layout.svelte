<script lang="ts">
	import '../app.pcss';
	import { QueryClient, QueryClientProvider } from '@sveltestack/svelte-query';

	const queryClient = new QueryClient({
		defaultOptions: {
			queries: {
				retry: true,
				refetchOnWindowFocus: false,
				refetchOnReconnect: true
			}
		}
	});
</script>

<QueryClientProvider client={queryClient}>
	<slot />
</QueryClientProvider>
