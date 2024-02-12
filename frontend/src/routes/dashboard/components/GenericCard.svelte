<script lang="ts">
	import type { Service } from '$lib/models/service';
	import { Card } from 'flowbite-svelte';
	import { useQuery } from '@sveltestack/svelte-query';

	export let data: Service;
	export let stats: any = null;
	export let actionButton: any = null;

	let online = useQuery(['ping', data.id], data.serviceType.executePing.bind(data.serviceType), {
		refetchInterval: 2000
	});
</script>

<Card class="relative shrink-0 flex-grow" on:click={() => data.goToService()}>
	<div class="flex gap-4">
		<div class="flex-none">
			<img src={data.icon} alt={data.name} class="h-16 w-16" />
		</div>
		<div class="flex-grow">
			<h3 class="text-lg font-semibold">{data.name}</h3>
			<p>{data.type}</p>
		</div>
		<svelte:component this={actionButton} {data} />
	</div>
	<svelte:component this={stats} {data} />
	{#if $online.isSuccess}
		{#if $online.data}
			<div class={`absolute right-3 top-0 flex-none rounded-b-lg bg-green-500 px-4`}>
				<p class="text-white">Online</p>
			</div>
		{:else}
			<div class={`absolute right-3 top-0 flex-none rounded-b-lg bg-red-500 px-4`}>
				<p class="text-white">Online</p>
			</div>
		{/if}
	{/if}
</Card>
