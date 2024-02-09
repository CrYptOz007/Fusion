<script lang="ts">
	import { Button, Modal, Label, Input, Select, Tooltip } from 'flowbite-svelte';
	import { PlusOutline } from 'flowbite-svelte-icons';
	import { useMutation } from '@sveltestack/svelte-query';
	import { addService } from '$lib/hooks/services';
	let formModel = false;
	let type: Service;
	$: formData = new FormData();

	type Service = {
		type: string;
		icon: string;
		port: number;
		apiKey: boolean;
		username: boolean;
		password: boolean;
	};

	const services: Service[] = [
		{
			type: 'Proxmox',
			icon: '/icons/proxmox.png',
			port: 8006,
			apiKey: true,
			username: true,
			password: false
		},
		{
			type: 'IPMI',
			icon: '',
			port: 623,
			apiKey: false,
			username: true,
			password: true
		},
		{
			type: 'Pihole',
			icon: '/icons/pihole.svg',
			port: 80,
			apiKey: true,
			username: false,
			password: false
		},
		{
			type: 'Unifi',
			icon: '/icons/unifi.png',
			port: 8443,
			apiKey: true,
			username: false,
			password: false
		},
		{
			type: 'Truenas',
			icon: '/icons/truenas.svg',
			port: 443,
			apiKey: true,
			username: false,
			password: false
		},
		{
			type: 'Home Assistant',
			icon: '/icons/home-assistant.png',
			port: 8123,
			apiKey: true,
			username: false,
			password: false
		},
		{
			type: 'Generic',
			icon: '',
			port: 80,
			apiKey: false,
			username: false,
			password: false
		}
	];

	const useAddService = useMutation(() => addService(formData), {
		onSuccess: () => {
			formModel = false;
		}
	});

	const handleSubmit = (e: Event) => {
		e.preventDefault();
		const formElement = e.target as HTMLFormElement;
		formData = new FormData(formElement);
		formData.set('type', type.type);
		$useAddService.mutate();
	};
</script>

<Button id="add-service" on:click={() => (formModel = true)}><PlusOutline /></Button>
<Tooltip triggeredBy="[id^='add-service']" placement="bottom">Add service</Tooltip>

<Modal bind:open={formModel} size="md" class="w-full" title="Add Service">
	<form class="flex flex-col space-y-6" on:submit={handleSubmit}>
		<Label for="name">Name</Label>
		<Input id="name" name="name" type="text" placeholder="Name" required />

		<Label for="type">Type</Label>
		<Select id="type" name="type" bind:value={type} required>
			{#each services as service}
				<option value={service}
					><img src={service.icon} alt={service.type} class="h-2 w-2" />{service.type}</option
				>
			{/each}
		</Select>

		<Label for="hostname">Hostname</Label>
		<Input id="hostname" name="hostname" type="text" placeholder="Hostname" required />

		<Label for="port">Port</Label>
		<Input id="port" name="port" type="number" placeholder="Port" value={type?.port} required />

		<Label for="icon">Icon</Label>
		<Input id="icon" name="icon" type="text" placeholder="Icon" value={type?.icon} required />

		{#if type?.apiKey}
			<Label for="apiKey">API Key</Label>
			<Input id="apiKey" name="api_key" type="password" placeholder="API Key" />
		{/if}

		{#if type?.username}
			<Label for="username">Username</Label>
			<Input id="username" name="username" type="text" placeholder="Username" />
		{/if}

		{#if type?.password}
			<Label for="password">Password</Label>
			<Input id="password" name="password" type="password" placeholder="Password" />
		{/if}

		<div class="flex flex-row justify-end space-x-2">
			<Button color="primary" type="submit">Add</Button>
		</div>
	</form>
</Modal>
