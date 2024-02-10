<script lang="ts">
	import { slugs } from '@store/slugs';
	import { Button, Heading, Input, Label } from 'flowbite-svelte';

	let url = '';
	let slug = '';

	$: slug_data = {
		id: ($slugs.length+1).toString(),
		slug: slug,
		redirect: url
	} ;

	async function handleSubmit() {
		// Handle form submission here
		await slugs.add_slug(slug_data)
	}
</script>

<div>
	<Heading class="text-center text-2xl">Criar novo Slug</Heading>
	<form class="flex flex-col gap-2" on:submit|preventDefault={handleSubmit}>
		<Label for="url">URL</Label>
		<Input type="url" id="url" bind:value={url} />

		<Label for="slug">Slug</Label>
		<Input type="text" id="slug" bind:value={slug} />

		<Button type="submit">Adicionar</Button>
	</form>
</div>
