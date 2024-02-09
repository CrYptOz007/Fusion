<script lang="ts">
	import { Navbar, NavBrand, Search } from 'flowbite-svelte';
	import AvatarMenu from './components/AvatarMenu.svelte';
	import AddService from './components/AddService.svelte';
	let search = '';
</script>

<Navbar class="py-6">
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
		<AvatarMenu />
	</div>
</Navbar>

<slot />
