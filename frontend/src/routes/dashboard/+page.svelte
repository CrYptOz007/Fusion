<script lang="ts">
	import { getServices } from '$lib/hooks/services';
	import { useQuery } from '@sveltestack/svelte-query';
	import { Card, Spinner } from 'flowbite-svelte';
	import { Service } from '../../lib/models/service';

	const servicesQuery = useQuery('services', getServices, {
		retry: true,
		refetchInterval: 5000,
		refetchOnMount: true,
		refetchOnReconnect: true,
		refetchOnWindowFocus: true
	});

	let data: Service[] = [];
	$: if ($servicesQuery.isSuccess) {
		data = $servicesQuery.data.data.map((service: any) => new Service(service));
	}

	$: console.log(data);
</script>

{#if $servicesQuery.isLoading}
	<main class="flex h-screen items-center justify-center">
		<Spinner size="10" color="primary" />
	</main>
{:else if $servicesQuery.isSuccess}
	<main class="grid gap-4 p-16 lg:grid-cols-3 xl:grid-cols-4">
		{#each data as service}
			<div class="card relative col-span-1 flex">
				<Card class="shrink-0 flex-grow" on:click={() => service.goToService()}>
					<div class="flex gap-4">
						<div class="flex-none">
							<img src={service.icon} alt={service.name} class="h-16 w-16" />
						</div>
						<div class="flex-grow">
							<h3 class="text-lg font-semibold">{service.name}</h3>
							<p>{service.type}</p>
						</div>
					</div>
				</Card>
				{#await service.serviceType.executePing() then online}
					{#if online}
						<div class={`absolute right-0 h-full w-4 flex-none rounded-r-lg bg-green-500`}></div>
					{:else}
						<div class={`absolute right-0 h-full w-4 flex-none rounded-r-lg bg-red-500`}></div>
					{/if}
				{:catch}
					<div class={`absolute right-0 h-full w-4 flex-none rounded-r-lg bg-red-500`}></div>
				{/await}
			</div>
		{/each}
	</main>
{/if}

<style>
	.card {
		transition: transform 0.3s;
	}

	.card:hover {
		transform: scale(1.05);
		cursor: pointer;
	}
</style>
