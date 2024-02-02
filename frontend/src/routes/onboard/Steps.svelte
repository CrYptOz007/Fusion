<script>
	import { Button, Progressbar } from 'flowbite-svelte';
	import { count as step } from '$lib/stores';
</script>

<div class="flex flex-row items-center space-x-2 text-center">
	<Button pill={true} color={`${$step > 0 ? 'primary' : 'dark'}`} class="px-4">1</Button>
	<Progressbar animate size="h-1.5" color="gray" progress={`${$step > 1 ? 100 : 0}`} />
	<Button pill={true} color={`${$step > 1 ? 'primary' : 'dark'}`} class="px-4">2</Button>
	<Progressbar animate size="h-1.5" color="gray" progress={`${$step > 2 ? 100 : 0}`} />
	<Button pill={true} color={`${$step > 2 ? 'primary' : 'dark'}`} class="px-4">3</Button>
</div>
