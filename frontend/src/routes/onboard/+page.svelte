<script lang="ts">
	import { OnMount } from 'fractils';
	import { count as step } from '$lib/stores';
	import { Button, Card } from 'flowbite-svelte';
	import { slide, fade, fly } from 'svelte/transition';
	import Register from './Register.svelte';
	import Steps from './Steps.svelte';
</script>

<div class="flex flex-col space-y-10">
	{#if $step > 1}<Steps />{/if}
	{#if $step == 1}
		<OnMount>
			<div class="flex flex-col items-center space-y-6 text-center">
				<h1 class="text-4xl font-black dark:text-white" in:fly={{ y: 100, duration: 1000 }}>
					Welcome to Fusion
				</h1>
				<h3
					class="text-xl font-medium dark:text-white"
					in:fly={{ y: 100, duration: 1000, delay: 300 }}
				>
					An all in one dashboard for all your self-hosted applications.
				</h3>
				<div in:fade={{ duration: 1000 }}>
					<Button color="primary" size="lg" on:click={step.increment}>Get Started</Button>
				</div>
			</div>
		</OnMount>
	{:else if $step == 2}
		<div in:slide={{ duration: 500, delay: 200 }}>
			<Register />
		</div>
	{:else}
		<div in:fly={{ x: -200, duration: 500, delay: 100 }}>
			<Card class="flex flex-col items-center space-y-6 text-center">
				<h1 class="text-4xl font-black dark:text-white">Congratulations!</h1>
				<h3 class="text-l font-medium dark:text-white">
					You're ready to go. Let's go to the dashboard and start adding some applications
				</h3>
				<div in:fade={{ duration: 1000 }}>
					<Button color="primary" size="lg" on:click={step.increment}>Go to Dashboard</Button>
				</div>
			</Card>
		</div>
	{/if}
</div>
