<script lang="ts">
	import { goto } from '$app/navigation';
	import SlugCreateForm from '@components/SlugCreateForm.svelte';
	import SlugsTable from '@components/SlugsTable.svelte';
	import { session } from '@store/auth';
	import { Heading } from 'flowbite-svelte';

	if (!$session.loggedIn) {
		goto('/login');
	}
</script>

<div class="flex flex-col gap-8">
	<Heading class="text-3xl">Dashboard</Heading>
	<hr />
	<SlugCreateForm />
	<hr />
	<SlugsTable />
	<hr class="pt-4" />
</div>
