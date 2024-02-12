<script lang="ts">
	import { getServices } from '$lib/hooks/services';
	import { useQuery } from '@sveltestack/svelte-query';
	import { Spinner } from 'flowbite-svelte';
	import { GenericService, Service } from '../../lib/models/service';
	import GenericCard from './components/GenericCard.svelte';
	import Proxmox from './components/Proxmox.svelte';
	import Pihole from './components/Pihole.svelte';
	import PiholeEnableButton from './components/PiholeEnableButton.svelte';

	const servicesQuery = useQuery('services', getServices, {
		refetchInterval: 5000
	});

	let data: Service[] = [];
	$: if ($servicesQuery.isSuccess) {
		data = $servicesQuery.data.data.map((service: any) => new Service(service));
	}
</script>

{#if $servicesQuery.isLoading}
	<main class="flex h-screen items-center justify-center">
		<Spinner size="10" color="primary" />
	</main>
{:else if $servicesQuery.isSuccess}
	<main class="grid grid-flow-dense grid-cols-1 gap-6 p-4 sm:grid-cols-2 md:p-16 xl:grid-cols-3">
		{#each data as service}
			<div
				class="card relative col-span-1 w-full {service.serviceType instanceof GenericService
					? 'row-span-1'
					: 'row-span-2'} flex"
			>
				{#if service.type === 'proxmox'}
					<GenericCard stats={Proxmox} data={service} />
				{:else if service.type === 'pihole'}
					<GenericCard stats={Pihole} actionButton={PiholeEnableButton} data={service} />
				{:else}
					<GenericCard data={service} />
				{/if}
			</div>
		{/each}
	</main>
{/if}

<style>
	.card {
		transition: transform 0.3s;
	}

	.card:hover {
		transform: scale(1.02);
		cursor: pointer;
	}
</style>
