<script lang="ts">
	import { count as step } from '$lib/stores';
	import { Card, Label, Input, Button, Progressbar, Helper, Spinner } from 'flowbite-svelte';
	import { useRegister, useLogin } from '$lib/hooks/auth';
	import { useMutation } from '@sveltestack/svelte-query';
	import axios from 'axios';

	let validations: number[] = [];
	let username = '';
	let password = '';
	let confirmPassword = '';
	let formError = '';

	$: validations = [
		password.length > 4 ? 1 : 0,
		password.length >= 8 ? 3 : 0,
		(password.match(/[A-Z]/)?.length || 0) * 2,
		(password.match(/[0-9]/)?.length || 0) * 2,
		password.match(/[&!@#$%^*]/)?.length ? 2 : 0
	];
	$: validated = validations[1] && validations[2] && validations[3];
	$: strength = validations.reduce((acc, curr) => acc + curr, 0);
	$: passwordMatch = confirmPassword == password;
	$: formData = new URLSearchParams({
		username: username,
		password: password
	});

	const register = useMutation(() => useRegister(formData), {
		onSuccess: () => {
			$login.mutate();
		},
		onError: (error) => {
			if (axios.isAxiosError(error)) {
				formError = error.response?.data.error[0];
			}
		}
	});

	const login = useMutation(() => useLogin(formData), {
		onSuccess: (data) => {
			sessionStorage.setItem('token', data.token);

			step.increment();
		}
	});

	function handleRegister(e: Event) {
		e.preventDefault();

		if (passwordMatch && validated) {
			$register.mutate();
		}
	}
</script>

<Card>
	<form class="flex flex-col space-y-6" on:submit={handleRegister}>
		<h3 class="text-xl font-black dark:text-white">Create your admin account</h3>
		<div class="flex w-full flex-col space-y-2">
			<Label for="username">Username</Label>
			<Input id="username" type="text" bind:value={username} placeholder="Username" required />
		</div>

		<div class="flex w-full flex-col space-y-2">
			<Label for="password">Password</Label>
			<Input
				color={`${password.length == 0 ? 'base' : validated ? 'green' : 'red'}`}
				bind:value={password}
				id="password"
				type="password"
				placeholder="Password"
				required
			/>
			<Progressbar
				animate
				size="h-2"
				color={`${strength <= 3 ? 'red' : strength <= 6 ? 'yellow' : 'green'}`}
				progress={strength * 10}
			/>
			<ul>
				<li class={`${validations[1] ? 'text-green-500' : 'text-red-500'}`}>
					Includes at least 8 characters
				</li>
				<li class={`${validations[2] ? 'text-green-500' : 'text-red-500'}`}>
					Includes at least 1 uppercase letter
				</li>
				<li class={`${validations[3] ? 'text-green-500' : 'text-red-500'}`}>
					Includes at least 1 number
				</li>
			</ul>
		</div>
		<div class="flex w-full flex-col space-y-2">
			<Label for="password">Confirm Password</Label>
			<Input
				color={`${confirmPassword.length == 0 ? 'base' : passwordMatch ? 'green' : 'red'}`}
				bind:value={confirmPassword}
				id="confirm-password"
				type="password"
				placeholder="Password"
				required
			/>
			{#if confirmPassword.length > 0 && !passwordMatch}<Helper class="mt-2" color="red"
					><span class="font-medium">The passwords must match!</span></Helper
				>{/if}
		</div>
		<Button color="primary" size="lg" type="submit"
			>{#if $register.isLoading}<Spinner
					class="me-3"
					size="4"
					color="white"
				/>{:else}Next{/if}</Button
		>
		{#if $register.isError}<Helper color="red"><span class="font-medium">{formError}</span></Helper
			>{/if}
	</form>
</Card>
