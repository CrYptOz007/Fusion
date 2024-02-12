<script lang="ts">
	import { type Service } from '$lib/models/service';
	import { Card } from 'flowbite-svelte';
	import { useQuery } from '@sveltestack/svelte-query';
	import proxmox from '$lib/hooks/proxmox';

	export let data: Service;
	let vms = useQuery('vms', () => proxmox.getVMs(data.id), { refetchInterval: 2000 });
	let containers = useQuery(['containers', data.id], () => proxmox.getContainers(data.id), {
		refetchInterval: 2000
	});
	let nodes = useQuery(['nodes', data.id], () => proxmox.getNodes(data.id), {
		refetchInterval: 2000
	});
	let runningVMs: any;
	let runningContainers: any;

	$: if ($vms.isSuccess) {
		runningVMs = $vms.data.filter((vm: { status: string }) => vm.status === 'running').length;
	}
	$: if ($containers.isSuccess) {
		runningContainers = $containers.data.filter(
			(container: { status: string }) => container.status === 'running'
		).length;
	}
</script>

<div class="grid grid-cols-2 justify-between gap-4 py-5">
	{#if $nodes.isSuccess}
		{#each $nodes.data.data.data as node}
			<Card
				class="w-full border-gray-100 bg-gray-100 text-center text-gray-700 shadow-none dark:border-gray-600 dark:bg-gray-600 sm:p-2 sm:px-6"
			>
				<h3 class="text-lg font-medium">{(node.cpu_usage * 100).toFixed(2)}%</h3>
				<h3 class="text-lg font-semibold">CPU Usage</h3>
			</Card>
			<Card
				class="w-full border-gray-100 bg-gray-100 text-center text-gray-700 shadow-none dark:border-gray-600 dark:bg-gray-600 sm:p-2 sm:px-6"
			>
				<h3 class="text-lg font-medium">{node.memory} GiB</h3>
				<h3 class="text-lg font-semibold">RAM Usage</h3>
			</Card>
		{/each}
	{/if}
	{#if $vms.isSuccess}
		<Card
			class="w-full border-gray-100 bg-gray-100 text-center text-gray-700 shadow-none dark:border-gray-600 dark:bg-gray-600 sm:p-2 sm:px-6"
		>
			<h3 class="text-lg font-medium">{runningVMs}/{$vms.data.length}</h3>
			<h3 class="text-lg font-semibold">VMS</h3>
		</Card>
	{/if}
	{#if $containers.isSuccess}
		<Card
			class="w-full border-gray-100 bg-gray-100 text-center text-gray-700 shadow-none dark:border-gray-600 dark:bg-gray-600 sm:p-2 sm:px-6"
		>
			<h3 class="text-lg font-medium">{runningContainers}/{$containers.data.length}</h3>
			<h3 class="text-lg font-semibold">Containers</h3>
		</Card>
	{/if}
</div>
