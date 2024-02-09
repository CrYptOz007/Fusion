<script lang="ts">
	import { goto } from '$app/navigation';
	import { authStore } from '$lib/stores';
	import { Card, Label, Input, Button, Helper, Spinner } from 'flowbite-svelte';
	import { useLogin } from '$lib/hooks/auth';
	import { useMutation } from '@sveltestack/svelte-query';
	import axios from 'axios';

	let username = '';
	let password = '';
	let formError = '';
	$: formData = new URLSearchParams({
		username: username,
		password: password
	});
  $: if($authStore) {
    goto('/dashboard');
  }

	const login = useMutation(() => useLogin(formData), {
		onSuccess: (data) => {
			sessionStorage.setItem('token', data.token);
			authStore.login();
			goto('/dashboard');
		},
		onError: (error) => {
			if (axios.isAxiosError(error)) {
				formError = error.response?.data.error[0];
			}
		}
	});

	function handleLogin(e: Event) {
		e.preventDefault();
		$login.mutate();
	}
  
</script>

{#await authStore.isLoggedIn()}
	<Spinner size="10" color="primary" />
{:catch}
	<Card>
		<form class="flex flex-col space-y-6 px-5 py-5" on:submit={handleLogin}>
			<h3 class="justify-self:center px-10 text-xl font-black dark:text-white">Login to Fusion</h3>
			<div class="flex w-full flex-col space-y-2">
				<Label for="username">Username</Label>
				<Input id="username" type="text" bind:value={username} placeholder="Username" required />
			</div>

			<div class="flex w-full flex-col space-y-2">
				<Label for="password">Password</Label>
				<Input
					bind:value={password}
					id="password"
					type="password"
					placeholder="Password"
					required
				/>
			</div>
			<Button color="primary" size="lg" type="submit"
				>{#if $login.isLoading}<Spinner
						class="me-3"
						size="4"
						color="white"
					/>{:else}Login{/if}</Button
			>
			{#if $login.isError}<Helper color="red"><span class="font-medium">{formError}</span></Helper
				>{/if}
		</form>
	</Card>
{/await}
