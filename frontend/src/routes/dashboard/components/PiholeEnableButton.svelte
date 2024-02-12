<script lang="ts">
	import { type Service } from '$lib/models/service';
	import { Button } from 'flowbite-svelte';
	import { useQueryClient, useQuery } from '@sveltestack/svelte-query';
	import { PlaySolid, PauseSolid } from 'flowbite-svelte-icons';
	import { useMutation } from '@sveltestack/svelte-query';
	import pihole from '$lib/hooks/pihole';
	import { darkMode } from '$lib/stores';

	export let data: Service;
	const queryClient = useQueryClient();
	let summary = useQuery(['summary', data.id], () => pihole.getSummaryData(data.id));

	const useEnable = useMutation(() => pihole.enable(data.id), {
		onMutate: async () => {
			queryClient.cancelQueries(['summary', data.id]);

			const previousValue = queryClient.getQueryData(['summary', data.id]);
			queryClient.setQueryData(['summary', data.id], {
				data: { ...(previousValue as any).data, status: 'enabled' }
			});

			return { previousValue };
		},
		onError: (err, newValue, context) => {
			queryClient.setQueryData(['summary', data.id], context);
		},
		onSettled: () => {
			setTimeout(() => {
				$summary.refetch();
			}, 1000);
		}
	});

	const useDisable = useMutation(() => pihole.disable(data.id), {
		onMutate: async () => {
			queryClient.cancelQueries(['summary', data.id]);

			const previousValue = queryClient.getQueryData(['summary', data.id]);
			queryClient.setQueryData(['summary', data.id], {
				data: { ...(previousValue as any).data, status: 'disabled' }
			});

			return { previousValue };
		},
		onError: (err, newValue, context) => {
			queryClient.setQueryData(['summary', data.id], context);
		},
		onSettled: () => {
			setTimeout(() => {
				$summary.refetch();
			}, 1000);
		}
	});

	const handleEnable = async (e: Event) => {
		e.stopPropagation();
		$useEnable.mutate();
	};

	const handleDisable = async (e: Event) => {
		e.stopPropagation();
		$useDisable.mutate();
	};
</script>

{#if $summary.isSuccess || $summary.isRefetching}
	{#if $summary.data.data.status === 'enabled'}
		<Button
			pill={true}
			color={$darkMode ? 'dark' : 'none'}
			class="m-auto h-12 !p-2"
			on:click={handleDisable}><PauseSolid class="h-8 w-8" /></Button
		>
	{:else}
		<Button
			pill={true}
			color={$darkMode ? 'dark' : 'none'}
			class="m-auto h-12 !p-2"
			on:click={handleEnable}><PlaySolid class="h-8 w-8" /></Button
		>
	{/if}
{/if}
