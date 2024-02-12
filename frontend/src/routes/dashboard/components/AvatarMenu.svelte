<script lang="ts">
	import { Dropdown, DropdownItem, Button } from 'flowbite-svelte';
	import { UserOutline } from 'flowbite-svelte-icons';
	import { useLogout } from '$lib/hooks/auth';
	import { useMutation } from '@sveltestack/svelte-query';
	import { darkMode } from '$lib/stores';

	const logout = useMutation(useLogout, {
		onSuccess: () => {
			sessionStorage.removeItem('token');
			window.location.reload();
		}
	});
</script>

<Button class="avatar" color={$darkMode ? 'dark' : 'alternative'} size="lg"><UserOutline /></Button>
<Dropdown triggeredBy=".avatar">
	<DropdownItem
		slot="footer"
		on:click={() => {
			$logout.mutate();
		}}>Sign out</DropdownItem
	>
</Dropdown>
