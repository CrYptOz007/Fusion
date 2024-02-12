<script lang="ts">
	import { type Service } from '$lib/models/service';
	import { Card } from 'flowbite-svelte';
	import { useQuery } from '@sveltestack/svelte-query';
	import pihole from '$lib/hooks/pihole';

	export let data: Service;
	let summary = useQuery(['summary', data.id], () => pihole.getSummaryData(data.id), {
		refetchInterval: 2000
	});
</script>

<div class="grid grid-cols-2 justify-between gap-4 py-5">
	{#if $summary.isSuccess}
		<Card
			class="w-full border-sky-800 bg-sky-700 text-center text-white shadow-none dark:border-sky-900 dark:bg-sky-800 sm:p-2 sm:px-6"
		>
			<h3 class="text-lg font-medium">{$summary.data.data.ads_blocked_today} Ads</h3>
			<h3 class="text-lg font-semibold">Blocked today</h3>
		</Card>
		<Card
			class="w-full border-yellow-700 bg-yellow-600 text-center text-white shadow-none dark:border-yellow-800 dark:bg-yellow-700 sm:p-2 sm:px-6"
		>
			<h3 class="text-lg font-medium">{$summary.data.data.ads_percentage_today}% Ads</h3>
			<h3 class="text-lg font-semibold">Being blocked</h3>
		</Card>
		<Card
			class="w-full border-violet-800 bg-violet-700 text-center text-white shadow-none dark:border-violet-900 dark:bg-violet-800 sm:p-2 sm:px-6"
		>
			<h3 class="text-lg font-medium">{$summary.data.data.dns_queries_today}</h3>
			<h3 class="text-lg font-semibold">Queries today</h3>
		</Card>
		<Card
			class="w-full {$summary.data.data.status === 'enabled'
				? 'border-green-700 bg-green-600 dark:border-green-800 dark:bg-green-700'
				: 'border-red-800 bg-red-700 dark:border-red-900 dark:bg-red-800'} text-center text-white shadow-none sm:p-2 sm:px-6"
		>
			<h3 class="text-lg font-semibold">Status</h3>
			<h3 class="text-lg font-medium">
				{$summary.data.data.status.charAt(0).toUpperCase() + $summary.data.data.status.slice(1)}
			</h3>
		</Card>
	{/if}
</div>
