<script lang="ts">
	import { Navbar, NavBrand, Search, Button } from 'flowbite-svelte';
	import AvatarMenu from './components/AvatarMenu.svelte';
	import AddService from './components/AddService.svelte';
	import { MoonSolid } from 'flowbite-svelte-icons';
	import { darkMode } from '$lib/stores';
	let search = '';

	const toggleDarkMode = () => {
		darkMode.toggle();
		document.body.className = $darkMode ? 'dark bg-gray-700' : 'bg-gray-100';
	};
</script>

<Navbar class="border-b-1 py-4 dark:border-gray-950 dark:bg-gray-800">
	<div class="flex flex-auto gap-4">
		<NavBrand href="/">
			<span class="self-center whitespace-nowrap text-2xl font-semibold dark:text-white"
				>Fusion</span
			>
		</NavBrand>
		<div class="w-full shrink" />
		<Search bind:search size="lg" />
		<div class="w-full shrink" />
		<AddService />
		<Button color={$darkMode ? 'dark' : 'alternative'} size="lg" on:click={toggleDarkMode}
			><MoonSolid class="h-4 w-4" /></Button
		>
		<AvatarMenu />
	</div>
</Navbar>

<slot />
